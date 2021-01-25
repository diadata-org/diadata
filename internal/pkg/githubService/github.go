package githubservice

import (
	"encoding/json"
	"fmt"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

const (
	githubURL                     = "https://api.github.com/graphql"
	githubGraphQLQueryTimeFormat  = "2006-01-02T15:04:05+00:00"
	githubGraphQLReturnTimeFormat = "2006-01-02T15:04:05Z"
)

type GithubRepo struct {
	Data Data `json:"data"`
}

type Data struct {
	Repository Repository `json:"repository"`
}

type Repository struct {
	ID     string `json:"id"`
	Object Object `json:"object"`
}

type Object struct {
	CommitURL string  `json:"commitUrl"`
	ID        string  `json:"id"`
	History   History `json:"history"`
}

type History struct {
	Nodes      []Node   `json:"nodes"`
	PageInfo   PageInfo `json:"pageInfo"`
	TotalCount int      `json:"totalCount"`
}

type Node struct {
	Hash            string `json:"oid"`
	Timestamp       string `json:"committedDate"`
	Author          Author `json:"author"`
	NumAdditions    int    `json:"additions"`
	NumDeletions    int    `json:"deletions"`
	NumChangedFiles int    `json:"changedFiles"`
	Message         string `json:"message"`
}

type PageInfo struct {
	EndCursor string `json:"endCursor"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FetchCommitsByDate returns all commits in repository @nameRepository of user @nameUser in the time range
// given by @timeInit and @timeFinal including both borders because the graphQL query does so.
func FetchCommitsByDate(nameUser, nameRepository, apiKey string, timeInit time.Time, timeFinal time.Time) (githubCommits []models.GithubCommit, err error) {
	// And interactive graphQL Explorer:
	// https://developer.github.com/v4/explorer/
	jsonData := map[string]string{
		"query": fmt.Sprintf(`{
			repository(owner:"%s", name:"%s") {
				id
				object(expression:"master") {
				  	commitUrl
				 	id
				  	... on Commit{
						history(since:"%s", until:"%s") {
					  		nodes{
								oid
								committedDate
                				author {
									  name
									  email
                				}
                				additions
                				deletions
								changedFiles
                				message
							}
							pageInfo {
                				endCursor
              				}
              				totalCount
						}
				  	}
				}
			}
		}
		`, nameUser, nameRepository, timeInit.Format(githubGraphQLQueryTimeFormat), timeFinal.Format(githubGraphQLQueryTimeFormat)),
	}

	jsonValue, _ := json.Marshal(jsonData)
	var bearer = "Bearer " + apiKey
	body, err := utils.GraphQLGet(githubURL, jsonValue, bearer)
	if err != nil {
		return
	}
	var repository GithubRepo
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return
	}
	githubCommits = githubRepoToCommit(nameUser, nameRepository, repository)

	// Check, whether all commits in the given time range were returned by a single query.
	endCursor := repository.Data.Repository.Object.History.PageInfo.EndCursor
	// log.Info("iterate through commits...")
	for endCursor != "" {
		// log.Info("current end cursor: ", endCursor)
		repository, endCursor, err = fetchCommitsAfterCursorInRange(nameUser, nameRepository, endCursor, timeInit, timeFinal, apiKey)
		aux := githubRepoToCommit(nameUser, nameRepository, repository)
		githubCommits = append(githubCommits, aux...)
	}
	return

}

func fetchCommitsAfterCursorInRange(nameUser, nameRepository, cursor string, timeInit, timeFinal time.Time, apiKey string) (GithubRepo, string, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`{
			repository(owner:"%s", name:"%s") {
				id
				object(expression:"master") {
					commitUrl
				  	id
				  	... on Commit{
            			history(since:"%s", until:"%s", after:"%s"){
                			nodes{
                  				oid
                  				committedDate
                  				author {
                    				name
                    				email
                  				}
                  			additions
                  			deletions
                  			changedFiles
                  			message
                			}
              			pageInfo {
                			endCursor
              			}
              			totalCount
            		}
				}
			}
		}
	}
	`, nameUser, nameRepository, timeInit.Format(githubGraphQLQueryTimeFormat), timeFinal.Format(githubGraphQLQueryTimeFormat), cursor),
	}

	jsonValue, _ := json.Marshal(jsonData)
	var bearer = "Bearer " + apiKey
	body, err := utils.GraphQLGet(githubURL, jsonValue, bearer)
	if err != nil {
		return GithubRepo{}, "", err
	}

	var repository GithubRepo
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return GithubRepo{}, "", err
	}
	endCursor := repository.Data.Repository.Object.History.PageInfo.EndCursor
	return repository, endCursor, nil
}

// FetchAllCommits returns all commits in repository @nameRepository of user @nameUser.
// @numPerPage is the number of commits fetched per request. For github this is limited to 100 atm.
func FetchAllCommits(nameUser, nameRepository string, numPerPage int, apiKey string) (githubCommits []models.GithubCommit, err error) {
	// Inital fetch to initiate pageing through all commits
	jsonData := map[string]string{
		"query": fmt.Sprintf(`{
			repository(owner:"%s", name:"%s") {
				id
				object(expression:"master") {
					commitUrl
				  	id
				  	... on Commit{
            			history(first:1){
                			nodes{
                  				oid
                  				committedDate
                  				author {
                    				name
                    				email
                  				}
                  				additions
                  				deletions
                  				changedFiles
                  				message
                			}
              				pageInfo {
                				endCursor
              				}
              				totalCount
            			}
					}
				}
			}
		}
		`, nameUser, nameRepository),
	}

	jsonValue, _ := json.Marshal(jsonData)
	var bearer = "Bearer " + apiKey
	body, err := utils.GraphQLGet(githubURL, jsonValue, bearer)
	if err != nil {
		return
	}
	var repository GithubRepo
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return
	}
	githubCommits = githubRepoToCommit(nameUser, nameRepository, repository)

	// Iterate over commits until cursor==null
	endCursor := repository.Data.Repository.Object.History.PageInfo.EndCursor
	for endCursor != "" {
		// log.Info("current end cursor: ", endCursor)
		repository, endCursor, err = fetchCommitsAfterCursor(nameUser, nameRepository, endCursor, numPerPage, apiKey)
		aux := githubRepoToCommit(nameUser, nameRepository, repository)
		githubCommits = append(githubCommits, aux...)
	}

	return
}

// FetchCommitsAfterCursor fetches the @numCommits next commits after @cursor and returns them, and the updated cursor.
func fetchCommitsAfterCursor(nameUser, nameRepository, cursor string, numCommits int, apiKey string) (GithubRepo, string, error) {
	jsonData := map[string]string{
		"query": fmt.Sprintf(`{
			repository(owner:"%s", name:"%s") {
				id
				object(expression:"master") {
					commitUrl
				  	id
				  	... on Commit{
            			history(first:%v, after:"%s"){
                			nodes{
                  				oid
                  				committedDate
                  				author {
                    				name
                    				email
                  				}
                  			additions
                  			deletions
                  			changedFiles
                  			message
                			}
              			pageInfo {
                			endCursor
              			}
              			totalCount
            		}
				}
			}
		}
	}
	`, nameUser, nameRepository, numCommits, cursor),
	}

	jsonValue, _ := json.Marshal(jsonData)
	var bearer = "Bearer " + apiKey
	body, err := utils.GraphQLGet(githubURL, jsonValue, bearer)
	if err != nil {
		return GithubRepo{}, "", err
	}

	var repository GithubRepo
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return GithubRepo{}, "", err
	}
	endCursor := repository.Data.Repository.Object.History.PageInfo.EndCursor
	return repository, endCursor, nil
}

func githubRepoToCommit(nameUser, nameRepository string, repo GithubRepo) (githubCommits []models.GithubCommit) {
	nodes := repo.Data.Repository.Object.History.Nodes

	var commit models.GithubCommit
	for _, node := range nodes {
		timestamp, err := time.Parse(githubGraphQLReturnTimeFormat, node.Timestamp)
		if err != nil {
			log.Error("error parsing time: ", err)
		}
		commit.User = nameUser
		commit.Repository = nameRepository
		commit.Hash = node.Hash
		commit.Timestamp = timestamp
		commit.Author.Name = node.Author.Name
		commit.Author.Email = node.Author.Email
		commit.NumAdditions = node.NumAdditions
		commit.NumDeletions = node.NumDeletions
		commit.NumChangedFiles = node.NumChangedFiles
		commit.Message = node.Message

		githubCommits = append(githubCommits, commit)
	}
	return
}
