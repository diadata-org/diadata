package githubservice

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
)

const (
	githubURL               = "https://api.github.com/graphql"
	githubGraphQLTimeFormat = "2006-01-02T15:04:05+00:00"
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

// GetCommitInfo returns all commits in repository @nameRepository of user @nameUser in the time range
// given by @timeInit and @timeFinal including both borders.
func GetCommitInfo(nameUser, nameRepository, apiKey string, timeInit time.Time, timeFinal time.Time) (GithubRepo, error) {

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
                				}
                				additions
                				deletions
								changedFiles
                				message
					  		}
						}
				  	}
				}
			}
		}
		`, nameUser, nameRepository, timeInit.Format(githubGraphQLTimeFormat), timeFinal.Format(githubGraphQLTimeFormat)),
	}

	jsonValue, _ := json.Marshal(jsonData)
	var bearer = "Bearer " + apiKey
	body, err := utils.GraphQLGet(githubURL, jsonValue, bearer)
	if err != nil {
		return GithubRepo{}, err
	}

	var repository GithubRepo
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return GithubRepo{}, err
	}

	return repository, nil

}

// GetAllCommits returns all commits in repository @nameRepository of user @nameUser.
func GetAllCommits(nameUser, nameRepository, apiKey string) (GithubRepo, error) {
	// TO DO: fetch commits using history(first:10,after:endCursor) until endCursor==nil
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
		return GithubRepo{}, err
	}

	var repository GithubRepo
	err = json.Unmarshal(body, &repository)
	if err != nil {
		return GithubRepo{}, err
	}
	fmt.Println("total count: ", repository.Data.Repository.Object.History.TotalCount)
	fmt.Println("endCursor: ", repository.Data.Repository.Object.History.PageInfo.EndCursor)
	return repository, nil
}
