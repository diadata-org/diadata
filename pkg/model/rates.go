package models

import (
	log "github.com/sirupsen/logrus"
)

// GetRates returns a (unique) list of all rates that have been written into the database
func (db *DB) GetRates() []string {
	log.Info("writing rate into set of rates")
	allRates := db.redisClient.SMembers(keyAllRates).Val()
	return allRates
}
