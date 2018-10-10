package estimator

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T) {
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
func TestQuartiles(t *testing.T) {
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

func TestFilter(t *testing.T) {
	cases := []struct {
		samples         []float64
		filteredSamples []float64
	}{
		{[]float64{10.2, 14.1, 14.4, 14.4, 14.4, 14.5, 14.5, 14.6, 14.7, 14.7, 14.7, 14.9, 15.1, 15.9, 16.4}, []float64{14.1, 14.4, 14.4, 14.4, 14.5, 14.5, 14.6, 14.7, 14.7, 14.7, 14.9, 15.1}},
		{[]float64{}, []float64{}},
		{[]float64{24, 25, 21, 23, 49, 29, 33}, []float64{21, 23, 24, 25, 29, 33}},
		{[]float64{0, 0, 0}, []float64{0, 0, 0}},
		{[]float64{0, 0, 0, 0}, []float64{0, 0, 0, 0}},
	}
	for i, c := range cases {
		fd := RemoveOutliers(c.samples)
		if len(fd) != len(c.filteredSamples) {
			t.Errorf("Filter sample data %d failed expected:%d samples got %d", i, len(c.filteredSamples), len(fd))
		} else {
			for j := range fd {
				if fd[j] != c.filteredSamples[j] {
					t.Errorf("Filter sample data %d failed element:%d is different expected:%f got %f", i, j, c.filteredSamples[j], fd[j])
				}
			}
		}
	}
}
