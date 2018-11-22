package estimators

import (
	"gonum.org/v1/gonum/stat/distuv"
)

//PDFLaplace Laplace Distribution
//https://en.wikipedia.org/wiki/Laplace_distribution
type PDFLaplace struct {
	name        string
	mu          float64
	b           float64
	initialized bool
	e           error
}

func newPDFLaplace() PDF {
	return &PDFLaplace{
		name:        Laplace,
		mu:          0,
		b:           0,
		initialized: false,
		e:           nil,
	}
}

//GetName return Laplace
func (pdf *PDFLaplace) GetName() string {
	return pdf.name
}

//GetLocation returns location parameter (mu)
func (pdf *PDFLaplace) GetLocation() float64 {
	return pdf.mu
}

//GetScale returns scale parameter or diversity (b)
func (pdf *PDFLaplace) GetScale() float64 {
	return pdf.b
}

//DidConverge Return true if the solution converged satisfactory
//TODO: Implasdement convergence check
func (pdf *PDFLaplace) DidConverge() bool {
	return pdf.initialized
}

//DidGetSolution returns true if enough points where provided to perform estimation
func (pdf *PDFLaplace) DidGetSolution() bool {
	return pdf.initialized

}

//GetError returns any error that affects the estimations
func (pdf *PDFLaplace) GetError() error {
	return pdf.e

}

//PDFEstimatorLaplace Estimates Laplace distribution from samples
type PDFEstimatorLaplace struct {
	distribution string
	pdf          PDF
	e            error
	samples      []float64
}

//NewPDFEstimatorLaplace Returns an estimator for Laplace distribution
func NewPDFEstimatorLaplace() *PDFEstimatorLaplace {
	return &PDFEstimatorLaplace{
		distribution: Laplace,
		pdf:          newPDFLaplace(),
	}
}

//GetPDF Get Laplace distribution. You need to call compute before
func (e *PDFEstimatorLaplace) GetPDF() PDF {
	return e.pdf

}

//AddSamples add samples to estimator
func (e *PDFEstimatorLaplace) AddSamples(samples []float64) {
	e.samples = append(e.samples, samples...)
}

//Compute estimate Laplace distribution parameters
func (e *PDFEstimatorLaplace) Compute() error {
	var d distuv.Laplace
	d.Fit(e.samples, nil)
	pdf := e.pdf.(*PDFLaplace)
	pdf.mu = d.Mu
	pdf.b = d.Scale
	pdf.initialized = true
	return nil
}
