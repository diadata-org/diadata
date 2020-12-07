package utils

import (
	log "github.com/sirupsen/logrus"
)

// UniqueStrings returns a subslice of @s such that each entry only appears once
func UniqueStrings(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}
	set := make(map[string]struct{})
	sUnique := []string{}
	for _, val := range s {
		if _, ok := set[val]; !ok {
			sUnique = append(sUnique, val)
		}
	}
	return sUnique
}

// Contains takes a slice of strings and a string and checks if it is contained in the slice.
func Contains(s *[]string, str string) bool {
	for _, a := range *s {
		if a == str {
			return true
		}
	}
	return false
}

// SliceDifference returns the elements in @slice1 that aren't in @slice2.
func SliceDifference(slice1, slice2 []string) []string {
	var diff []string
	for _, str := range slice1 {
		if !Contains(&slice2, str) {
			diff = append(diff, str)
		}
	}
	return diff
}

// MaxString return the maximum of a slice of strings along with its index
func MaxString(sl []string) (string, int64) {
	if len(sl) < 1 {
		log.Error("Cannot find maximum in empty slice")
		return "", 0
	}
	index := int64(0)
	max := sl[0]
	for k, entry := range sl {
		if entry > max {
			max = entry
			index = int64(k)
		}
	}
	return max, index
}

// MinString return the maximum of a slice of strings along with its index
func MinString(sl []string) (string, int64) {
	if len(sl) < 1 {
		log.Error("Cannot find minimum in empty slice")
		return "", 0
	}
	index := int64(0)
	min := sl[0]
	for k, entry := range sl {
		if entry < min {
			min = entry
			index = int64(k)
		}
	}
	return min, index
}
