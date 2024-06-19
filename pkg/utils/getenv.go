package utils

import (
	"os"
	"strconv"
)

func Getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func IsEnvExist(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		return true
	}
	return false
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

func GetenvBool(key string, fallback bool) bool {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Error parsing %s, using default value %d: %v", key, fallback, err)
		return fallback
	}
	return boolValue
}

func GetenvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	intValue, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Printf("Error parsing %s, using default value %d: %v", key, fallback, err)
		return fallback
	}
	return int(intValue)
}
