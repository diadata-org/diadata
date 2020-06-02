package models

import (
	"time"
)

// GetConfigTogglePairDiscovery switches to true between midnight and midnight + duration
func (db *DB) GetConfigTogglePairDiscovery(d time.Duration) bool {
	t := time.Now()
	secondsAfterMidnight := t.Hour()*3600 + t.Minute()*60 + t.Second()
	if float64(secondsAfterMidnight) < d.Seconds()+10 {
		return true
	}
	return false
}
