package constants

import "time"

const (
	cachingTime1Sec   = 1 * time.Second
	cachingTime20Secs = 20 * time.Second
	cachingTimeShort  = time.Minute * 2
	// cachingTimeMedium = time.Minute * 10
	cachingTimeLong = time.Minute * 100
)
