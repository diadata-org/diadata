package models

import "time"

type GithubCommit struct {
	User            string
	Repository      string
	Hash            string
	Timestamp       string
	Author          Author
	NumAdditions    int
	NumDeletions    int
	NumChangedFiles int
	Message         string
}

type Author struct {
	Name  string
	Email string
}

// SetCommit stores a github commit in influx
func (db *DB) SetCommit(commit *GithubCommit) error {
	// TO DO
	return nil
}

// GetCommitByDate returns the latest commit from @repository of github user @user before @date.
func (db *DB) GetCommitByDate(user, repository string, date time.Time) (GithubCommit, error) {
	// TO DO
	return GithubCommit{}, nil
}

// GetCommitByHash returns the commit from @repository of github user @user with hash @hash.
func (db *DB) GetCommitByHash(user, repository, hash string) (GithubCommit, error) {
	// TO DO
	return GithubCommit{}, nil
}
