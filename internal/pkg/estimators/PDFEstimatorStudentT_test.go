package estimators

import (
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat/distuv"
	"testing"
)

func TestStudentFitWithRandomSamples(t *testing.T) {
	cases := []struct {
		nSamples  int
		wantNu    float64
		wantMu    float64
		wantSigma float64
	}{
		{
			nSamples:  10000,
			wantNu:    1,
			wantMu:    0,
			wantSigma: 1,
		},
		{
			nSamples:  10000,
			wantNu:    1,
			wantMu:    10,
			wantSigma: 1,
		},
		{
			nSamples:  300000,
			wantNu:    10,
			wantMu:    0,
			wantSigma: 2,
		},
	}
	for i, test := range cases {
		src := rand.New(rand.NewSource(123))
		s := distuv.StudentsT{
			Mu:    test.wantMu,
			Nu:    test.wantNu,
			Sigma: test.wantSigma,
			Src:   src,
		}
		samples := make([]float64, test.nSamples)
		for i := range samples {
			samples[i] = s.Rand()
		}
		e := NewPDFEstimatorStudentT()
		if r, err := e.fit(samples); err != nil {
			t.Errorf(err.Error())
		} else {
			if r.Status > optimize.Failure && r.Status < optimize.Success {
				t.Errorf("Fit did not converged:%d", r.Status)
			}
			if !floats.EqualWithinAbsOrRel(e.s.Nu, s.Nu, 1e-2, 1e-2) {
				t.Errorf("unexpected Nu result for test :%d got:%f, want:%f", i, e.s.Nu, s.Nu)
			}
			if !floats.EqualWithinAbsOrRel(e.s.Mu, s.Mu, 1e-2, 1e-2) {
				t.Errorf("unexpected Mu result for test :%d got:%f, want:%f", i, e.s.Mu, s.Mu)
			}
			if !floats.EqualWithinAbsOrRel(e.s.Sigma, s.Sigma, 2e-2, 2e-2) {
				t.Errorf("unexpected Sigma result for test :%d got:%f, want:%f", i, e.s.Sigma, s.Sigma)
			}
		}
	}
}

func TestPDFEstimatorStudentComputeWithRandomSamples(t *testing.T) {
	cases := []struct {
		nSamples  int
		wantNu    float64
		wantMu    float64
		wantSigma float64
	}{
		{
			nSamples:  10000,
			wantNu:    1,
			wantMu:    0,
			wantSigma: 1,
		},
	}
	for i, test := range cases {
		src := rand.New(rand.NewSource(123))
		s := distuv.StudentsT{
			Mu:    test.wantMu,
			Nu:    test.wantNu,
			Sigma: test.wantSigma,
			Src:   src,
		}
		samples := make([]float64, test.nSamples)
		for i := range samples {
			samples[i] = s.Rand()
		}
		e := NewPDFEstimatorStudentT()
		e.AddSamples(samples)
		if err := e.Compute(); err != nil {
			t.Errorf(err.Error())
		} else {
			pdf := e.GetPDF().(*PDFStudentT)
			if !pdf.DidConverge() {
				t.Errorf("Fit did not converge")
			}
			if !floats.EqualWithinAbsOrRel(e.GetPDF().(*PDFStudentT).GetNu(), s.Nu, 1e-2, 1e-2) {
				t.Errorf("unexpected Nu result for test :%d got:%f, want:%f", i, e.GetPDF().(*PDFStudentT).GetNu(), s.Nu)
			}
			if !floats.EqualWithinAbsOrRel(e.GetPDF().(*PDFStudentT).GetLocation(), s.Mu, 1e-2, 1e-2) {
				t.Errorf("unexpected Mu result for test :%d got:%f, want:%f", i, e.GetPDF().(*PDFStudentT).GetLocation(), s.Mu)
			}
			if !floats.EqualWithinAbsOrRel(e.GetPDF().(*PDFStudentT).GetScale(), s.Sigma, 2e-2, 2e-2) {
				t.Errorf("unexpected Sigma result for test :%d got:%f, want:%f", i, e.GetPDF().(*PDFStudentT).GetScale(), s.Sigma)
			}
		}
	}
}

func TestPDFEstimatorStudentComputeNotEnoughSamples(t *testing.T) {
	e := NewPDFEstimatorStudentT()
	samples := make([]float64, e.GetMinimumNumberOfSamples()-1)
	e.AddSamples(samples)
	if err := e.Compute(); err == nil {
		t.Errorf("Error expected")
	}
}
