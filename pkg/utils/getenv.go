package utils

import (
	"os"
	"strconv"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func GetenvUint(key string, fallback uint) uint {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	uint64Val, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Printf("Error parsing %s, using default value %d: %v", key, fallback, err)
		return fallback
	}

	return uint(uint64Val)
}

func IsEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
}
