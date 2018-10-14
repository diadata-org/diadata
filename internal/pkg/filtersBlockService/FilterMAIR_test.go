package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"math"
	"testing"
	"time"
)

func TestFilterMAIRInternalMean(t *testing.T) {
	cases := []struct {
		samples []float64
		mean    float64
	}{
		{[]float64{59.4896, 26.2212, 60.2843, 71.1216, 22.1747, 11.7418, 29.6676, 31.8778, 42.4167, 50.7858}, 40.5781},
		{[]float64{6, 3, 2, 4, 5, 1}, 3.5},
		{[]float64{1}, 1.0},
		{[]float64{}, 0.0},
	}
	for i, c := range cases {
		mean := computeMean(c.samples)
		if math.Abs(float64(mean-c.mean)) > 1e-4 {
			t.Errorf("Mean was incorrect, got: %f, expected: %f for set:%d", mean, c.mean, i)
		}
	}
}
func TestFilterMAIRInternalMedian(t *testing.T) {
	cases := []struct {
		samples []float64
		median  float64
	}{
		{[]float64{59.4896, 26.2212, 60.2843, 71.1216, 22.1747, 11.7418, 29.6676, 31.8778, 42.4167, 50.7858}, 37.1472},
		{[]float64{26.2212, 60.2843, 71.1216, 22.1747, 11.7418, 29.6676, 31.8778, 42.4167, 50.7858}, 31.8778},
		{[]float64{5, 3, 4, 2, 1}, 3.0},
		{[]float64{6, 3, 2, 4, 5, 1}, 3.5},
		{[]float64{1}, 1.0},
		{[]float64{}, 0.0},
	}
	for i, c := range cases {
		median := computeMedian(c.samples)
		if math.Abs(float64(median-c.median)) > 1e-4 {
			t.Errorf("Median was incorrect, got: %f, expected: %f for set:%d", median, c.median, i)
		}
	}

}
func TestFilterMAIRInternalQuartiles(t *testing.T) {
	cases := []struct {
		samples []float64
		Q1      float64
		Q3      float64
	}{
		{[]float64{59.4896, 26.2212, 60.2843, 71.1216, 22.1747, 11.7418, 29.6676, 31.8778, 42.4167, 50.7858}, 26.2212, 59.4896},
		{[]float64{6, 7, 39, 15, 36, 40, 41, 42, 47, 43, 49}, 15, 43},
		{[]float64{7, 15, 36, 39, 40, 41}, 15, 40},
	}
	for i, c := range cases {
		Q1, Q3 := computeQuartiles(c.samples)
		if math.Abs(float64(Q1-c.Q1)) > 1e-5 || math.Abs(float64(Q3-c.Q3)) > 1e-5 {
			t.Errorf("Quartiles was incorrect, got: [%f,%f] expected: [%f,%f] for set:%d", Q1, Q3, c.Q1, c.Q3, i)
		}
	}
}
func TestFilterMAIRInternalRemoveOutliers(t *testing.T) {
	cases := []struct {
		samples         []float64
		filteredSamples []float64
	}{
		{[]float64{10.2, 14.1, 14.4, 14.4, 14.4, 14.5, 14.5, 14.6, 14.7, 14.7, 14.7, 14.9, 15.1, 15.9, 16.4}, []float64{14.1, 14.4, 14.4, 14.4, 14.5, 14.5, 14.6, 14.7, 14.7, 14.7, 14.9, 15.1}},
		{[]float64{}, []float64{}},
		{[]float64{24, 25, 21, 23, 49, 29, 33}, []float64{21, 23, 24, 25, 29, 33}},
		{[]float64{0, 0, 0}, []float64{0, 0, 0}},
		{[]float64{0, 0, 0, 0}, []float64{0, 0, 0, 0}},
		{[]float64{50, 50, 50, 200, 50, 50, 50}, []float64{50, 50, 50, 50, 50, 50}},
	}
	for i, c := range cases {
		cd := removeOutliers(c.samples)
		if len(cd) != len(c.filteredSamples) {
			t.Errorf("Filter sample data %d failed expected:%d samples got %d", i, len(c.filteredSamples), len(cd))
		} else {
			for j := range cd {
				if cd[j] != c.filteredSamples[j] {
					t.Errorf("Filter sample data %d failed element:%d is different expected:%f got %f", i, j, c.filteredSamples[j], cd[j])
				}
			}
		}
	}
}
func TestFilterMAIRIgnore(t *testing.T) {
	filterParam := 10
	firstPrice := 50.0
	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	steps := filterParam
	f := NewFilterMAIR("XRP", "", d, filterParam)
	p := firstPrice
	priceIncrements := 1.0
	for i := 0; i <= steps; i++ {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(-time.Second)
		p += priceIncrements
	}
	f.finalComputeEndOfBlock(d)
	v := f.filterPointForBlock()
	if v.Value != firstPrice {
		t.Errorf("error should be initial value:%f got:%f", firstPrice, v.Value)
	}
}
func TestFilterMAIRAverage(t *testing.T) {
	filterParam := 30
	firstPrice := 50.0
	avg := 0.
	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMAIR("XRP", "", d, filterParam)
	p := firstPrice
	priceIncrements := 1.0
	samples := 15
	for i := 0; i < samples; i++ {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(time.Second)
		avg += p
		p += priceIncrements
	}
	// append last value twice. Same as filter
	avg += p - priceIncrements
	avg = avg / float64(samples+1)
	f.finalComputeEndOfBlock(d)
	v := f.filterPointForBlock()
	if v.Value != avg {
		t.Errorf("error should be average value:%f got:%f", avg, v.Value)
	}
}
func TestFilterMAIRAverageOutsideRange(t *testing.T) {
	memory := 5
	firstPrice := 50.0
	avg := 0.
	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMAIR("XRP", "", d, memory)
	p := firstPrice
	priceIncrements := 1.0
	samples := 15
	for i := 0; i < samples; i++ {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(time.Second)
		if samples-i <= memory {
			avg += p
		}
		p += priceIncrements
	}
	// append last value twice. Same as filter
	avg = (avg + priceIncrements*float64(memory-1)) / float64(memory)
	f.finalComputeEndOfBlock(d)
	v := f.filterPointForBlock()
	if v.Value != avg {
		t.Errorf("error should be average value:%f got:%f", avg, v.Value)
	}
}
func TestFilterMAIRAverageCleanOutliers(t *testing.T) {
	cases := []struct {
		samples []float64
		mean    float64
	}{
		{[]float64{50, 50, 50, 200, 50, 50, 50}, 50},
		{[]float64{6, 3, 2, 4, 5, 1}, 3.1428},
		{[]float64{1}, 1.0},
		{[]float64{}, 0.0},
		{[]float64{10.2, 14.1, 14.4, 14.4, 14.4, 14.5, 14.5, 14.6, 14.7, 14.7, 14.7, 14.9, 15.1, 15.9, 16.4}, 14.5833},
		{[]float64{}, 0},
		{[]float64{24, 25, 21, 23, 49, 29, 33}, 26.8571},
		{[]float64{0, 0, 0}, 0},
		{[]float64{0, 0, 0, 0}, 0},
	}
	for i, c := range cases {
		memory := 20
		d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
		f := NewFilterMAIR("XRP", "", d, memory)
		for _, p := range c.samples {
			f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
			d = d.Add(time.Second)
		}
		f.finalComputeEndOfBlock(d)
		v := f.filterPointForBlock()
		if math.Abs(float64(v.Value-c.mean)) > 1e-4 {
			t.Errorf("Mean was incorrect, got: %f, expected: %f for set:%d", v.Value, c.mean, i)
		}
	}
}
