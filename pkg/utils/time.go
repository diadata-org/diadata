package utils

import "time"

func GetTimeDurationFromIntAsMilliseconds(input int) time.Duration {
	result := time.Duration(input) * time.Millisecond
	return result
}
