package utils

import (
	"runtime"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
// Taken from:
// https://gist.github.com/j33ty/79e8b736141be19687f565ea4c6f4226
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Infof("Alloc = %v MiB", bToMb(m.Alloc))
	log.Infof("TotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	log.Infof("Sys = %v MiB", bToMb(m.Sys))
	log.Infof("NumGC = %v", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
