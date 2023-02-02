package constants

import "time"

const (
	CachingTime1Sec   = 1 * time.Second
	CachingTime20Secs = 20 * time.Second
	CachingTimeShort  = time.Minute * 2
	CachingTimeMedium = time.Minute * 10
	CachingTimeLong   = time.Minute * 100
)
