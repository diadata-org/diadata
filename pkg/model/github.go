package models

import (
	"encoding/json"
	"fmt"
	"time"

	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	log "github.com/sirupsen/logrus"
)

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
	log.Info("set commit: ", commit)
	fields := map[string]interface{}{
		"numAdditions":    commit.NumAdditions,
		"numDeletions":    commit.NumDeletions,
		"numChangedFiles": commit.NumChangedFiles,
		"message":         commit.Message,
	}
	tags := map[string]string{
		"user":       commit.User,
		"repository": commit.Repository,
		"hash":       commit.Hash,
		"authorname": commit.Author.Name,
		"authormail": commit.Author.Email,
	}
	pt, err := clientInfluxdb.NewPoint(influxDbGithubCommitTable, tags, fields, commit.Timestamp)
	if err != nil {
		log.Errorln("SetCommit:", err)
	} else {
		db.addPoint(pt)
	}

	err = db.WriteBatchInflux()
	if err != nil {
		log.Errorln("SetCommit: ", err)
	}

	return err
}

// GetCommitByDate returns the latest commit from @repository of github user @user before @date.
func (db *DB) GetCommitByDate(user, repository string, date time.Time) (GithubCommit, error) {
	var commit GithubCommit
	q := fmt.Sprintf("select authorname,authormail,hash,message,numAdditions,numDeletions,numChangedFiles from %s where \"user\"='%s' and \"repository\"='%s' and time<%d order by desc limit 1", influxDbGithubCommitTable, user, repository, date.UnixNano())
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return commit, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		commit.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return commit, err
		}
		commit.Author.Name = res[0].Series[0].Values[0][1].(string)
		if err != nil {
			return commit, err
		}
		commit.Author.Email = res[0].Series[0].Values[0][2].(string)
		commit.Hash = res[0].Series[0].Values[0][3].(string)
		commit.Message = res[0].Series[0].Values[0][4].(string)
		additions, err := res[0].Series[0].Values[0][5].(json.Number).Int64()
		if err != nil {
			return commit, err
		}
		commit.NumAdditions = int(additions)
		deletions, err := res[0].Series[0].Values[0][6].(json.Number).Int64()
		if err != nil {
			return commit, err
		}
		commit.NumDeletions = int(deletions)
		changedFiles, err := res[0].Series[0].Values[0][7].(json.Number).Int64()
		if err != nil {
			return commit, err
		}
		commit.NumChangedFiles = int(changedFiles)
		if err != nil {
			return commit, err
		}
		commit.User = user
		commit.Repository = repository
	}
	return commit, nil
}

// GetCommitByHash returns the commit from @repository of github user @user with hash @hash.
func (db *DB) GetCommitByHash(user, repository, hash string) (GithubCommit, error) {
	var commit GithubCommit
	q := fmt.Sprintf("select authorname,authormail,hash,message,numAdditions,numDeletions,numChangedFiles from %s where \"user\"='%s' and \"repository\"='%s' and \"hash\"='%s'", influxDbGithubCommitTable, user, repository, hash)
	res, err := queryInfluxDB(db.influxClient, q)
	if err != nil {
		return commit, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		commit.Timestamp, err = time.Parse(time.RFC3339, res[0].Series[0].Values[0][0].(string))
		if err != nil {
			return commit, err
		}
		commit.Author.Name = res[0].Series[0].Values[0][1].(string)
		if err != nil {
			return commit, err
		}
		commit.Author.Email = res[0].Series[0].Values[0][2].(string)
		commit.Hash = res[0].Series[0].Values[0][3].(string)
		commit.Message = res[0].Series[0].Values[0][4].(string)
		additions, err := res[0].Series[0].Values[0][5].(json.Number).Int64()
		if err != nil {
			return commit, err
		}
		commit.NumAdditions = int(additions)
		deletions, err := res[0].Series[0].Values[0][6].(json.Number).Int64()
		if err != nil {
			return commit, err
		}
		commit.NumDeletions = int(deletions)
		changedFiles, err := res[0].Series[0].Values[0][7].(json.Number).Int64()
		if err != nil {
			return commit, err
		}
		commit.NumChangedFiles = int(changedFiles)
		if err != nil {
			return commit, err
		}
		commit.User = user
		commit.Repository = repository
	}
	return commit, nil
}

// GetLatestCommit returns the latest commit from influx.
// Returns empty struct and nil if no commits are in the database.
func (db *DB) GetLatestCommit(user, repository string) (GithubCommit, error) {
	commit, err := db.GetCommitByDate(user, repository, time.Now())
	if err != nil {
		return GithubCommit{}, err
	}
	return commit, nil
}
