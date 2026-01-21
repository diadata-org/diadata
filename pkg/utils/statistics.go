package utils

import (
	"errors"
	"math"
	"sort"
)

// DiscardOutliers discards every data point from @prices and @volumes that deviates from
// the price median by more than @basispoints basis points.
func DiscardOutliers(prices []float64, volumes []float64, basispoints float64) (newPrices []float64, newVolumes []float64, discarded []int, err error) {
	if len(prices) != len(volumes) {
		err = errors.New("number of prices does not equal number of volumes ")
		return
	}
	median := ComputeMedian(prices)
	threshold := basispoints * float64(0.0001) * median
	for i := 0; i < len(prices); i++ {
		if math.Abs(prices[i]-median) < threshold {
			newPrices = append(newPrices, prices[i])
			newVolumes = append(newVolumes, volumes[i])
		} else {
			discarded = append(discarded, i)
		}
	}
	return
}

// ComputeMedian returns the median of @samples.
func ComputeMedian(samples []float64) (median float64) {
	var length = len(samples)
	if length > 0 {
		sort.Float64s(samples)
		if length%2 == 0 {
			median = (samples[length/2-1] + samples[length/2]) / 2
		} else {
			median = samples[(length+1)/2-1]
		}
	}
	return
}

// Vwap returns the volume weighted average price for the slices @prices and @volumes.
func Vwap(prices []float64, volumes []float64) (float64, error) {
	//log.Info("prices, volumes: ", prices, volumes)
	if len(prices) != len(volumes) {
		return 0, errors.New("number of prices does not equal number of volumes ")
	}
	avg := float64(0)
	totalVolume := float64(0)
	for i := 0; i < len(prices); i++ {
		avg += prices[i] * math.Abs(volumes[i])
		totalVolume += math.Abs(volumes[i])
	}
	if totalVolume > 0 {
		return avg / totalVolume, nil
	} else {
		return 0, nil
	}
}
