package estimators

import (
	"math"
	"testing"
)

func compareNumbers(t *testing.T, expected float64, got float64, i int, test string) {
	if math.Abs(float64(expected-got)) > 1e-4 {
		t.Errorf(test+" was incorrect, got: %f, expected: %f for set:%d", got, expected, i)
	}
}
func TestPDFEstimatorLaplace(t *testing.T) {
	cases := []struct {
		samples []float64
		weights []float64
		mu      float64
		b       float64
	}{
		{[]float64{8.749301274757595, 6.873071190482888, 10.523406072893454, 7.531696772945777, 5.250203726375541, 9.392686695859132, 14.353503126218946, 8.091053148128688, 9.678239747448012, 10.014582029749249},
			nil, 8.7493, 1.7467},
		{[]float64{1, 1},
			[]float64{1, 0}, 1, 0},
	}
	for i, c := range cases {
		e := NewPDFEstimatorLaplace()
		e.AddSamples(c.samples) //, c.weights)
		e.Compute()
		pdf := e.GetPDF().(*PDFLaplace)
		compareNumbers(t, c.mu, pdf.mu, i, "Location")
		compareNumbers(t, c.b, pdf.b, i, "Scale")
	}
}
func TestPDFEstimatorLaplaceComputeNotEnoughSamples(t *testing.T) {
	e := NewPDFEstimatorLaplace()
	e.AddSamples([]float64{})
	if err := e.Compute(); err == nil {
		t.Errorf("Error expected")
	}
}
