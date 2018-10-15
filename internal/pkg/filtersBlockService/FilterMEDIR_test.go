package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"math"
	"testing"
	"time"
)

func TestFilterMEDIRInternalRemoveOutliers(t *testing.T) {
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

func TestFilterMEDIRMedianCleanOutliers(t *testing.T) {
	cases := []struct {
		samples []float64
		mean    float64
	}{

		{[]float64{10.2, 14.1, 14.4, 14.4, 14.4, 14.5, 14.5, 14.6, 14.7, 14.7, 14.7, 14.9, 15.1, 15.9, 16.4}, 14.55},
		{[]float64{}, 0.0},
		{[]float64{24, 25, 21, 23, 49, 29, 33}, 24.5},
		{[]float64{0, 0, 0}, 0},
		{[]float64{0, 0, 0, 0}, 0},
		{[]float64{50, 50, 50, 200, 50, 50, 50}, 50},
	}
	for i, c := range cases {
		memory := 20
		d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
		f := NewFilterMEDIR("XRP", "", d, memory)
		for _, p := range c.samples {
			f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
			d = d.Add(time.Second)
		}
		v := f.finalCompute(d)
		if math.Abs(float64(v-c.mean)) > 1e-4 {
			t.Errorf("Median was incorrect, got: %f, expected: %f for set:%d", v, c.mean, i)
		}
	}
}
