package models

import "time"

type GithubCommit struct {
	User            string
	Repository      string
	Hash            string
	Timestamp       time.Time
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

// GetLatestCommit returns the latest commit from influx.
// Returns empty struct and nil if no commits are in the database.
func (db *DB) GetLatestCommit(user, repository string) (GithubCommit, error) {
	// TO DO
	return GithubCommit{}, nil
}
