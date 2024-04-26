package utils

import (
	"runtime"
)

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	log.Infof("Alloc = %v MiB", bToMb(m.Alloc))
	log.Infof("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	log.Infof("\tSys = %v MiB", bToMb(m.Sys))
	log.Infof("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
