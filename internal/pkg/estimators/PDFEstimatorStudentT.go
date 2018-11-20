package estimators

import (
	// "gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/optimize"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
	"math"
)

//PDFStudentT implements the Student's T distribution with density function:
//
//  Γ((ν+1)/2) / (sqrt(νπ) Γ(ν/2) σ) (1 + 1/ν * ((x-μ)/σ)^2)^(-(ν+1)/2)
//
// see https://en.wikipedia.org/wiki/Student%27s_t-distribution,
type PDFStudentT struct {
	name        string
	initialized bool
	converged   bool
	e           error
	nu          float64
	mu          float64
	sigma       float64
}

func newPDFStudentT() PDF {
	return &PDFStudentT{
		name:        Student,
		initialized: false,
		converged:   false,
		e:           nil,
	}
}

//GetName return Laplace
func (pdf *PDFStudentT) GetName() string {
	return pdf.name
}

//GetLocation returns location parameter (mu)
func (pdf *PDFStudentT) GetLocation() float64 {
	return pdf.mu
}

//GetScale returns scale parameter (sigma)
func (pdf *PDFStudentT) GetScale() float64 {
	return pdf.sigma
}

//GetNu returns the varying parameters nu
func (pdf *PDFStudentT) GetNu() float64 {
	return pdf.nu
}

//DidConverge Return true if the solution converged satisfactory
func (pdf *PDFStudentT) DidConverge() bool {
	return pdf.converged
}

//DidGetSolution returns true if enough points where provided to perform estimation
func (pdf *PDFStudentT) DidGetSolution() bool {
	return pdf.initialized

}

//GetError returns any error that affects the estimations
func (pdf *PDFStudentT) GetError() error {
	return pdf.e

}

//PDFEstimatorStudentT Estimates Students T distribution from samples
type PDFEstimatorStudentT struct {
	distribution string
	pdf          PDF
	e            error
	samples      []float64
	// Here until it can be moved to gonum
	s distuv.StudentsT
}

//NewPDFEstimatorStudentT Returns an estimator for Laplace distribution
func NewPDFEstimatorStudentT() *PDFEstimatorStudentT {
	return &PDFEstimatorStudentT{
		distribution: Student,
		pdf:          newPDFStudentT(),
		s:            distuv.StudentsT{},
	}
}

//GetPDF Get Students T distribution. You need to call compute before
func (e *PDFEstimatorStudentT) GetPDF() PDF {
	return e.pdf
}

//AddSamples add samples to estimator
func (e *PDFEstimatorStudentT) AddSamples(samples []float64) {
	e.samples = append(e.samples, samples...)
}
func (e *PDFEstimatorStudentT) fit(samples []float64) (*optimize.Result, error) {
	s := e.s
	mu := stat.Mean(samples, nil)
	m2 := stat.Moment(2, samples, nil)
	m4 := stat.Moment(4, samples, nil)
	nu := 4 + (6*math.Pow(m2, 2))/(m4-3*math.Pow(m2, 2))
	sigma := math.Sqrt((nu - 2) * m2 / nu)
	s.Nu = nu
	s.Mu = mu
	s.Sigma = sigma

	fOptimize := func(param []float64) float64 {
		s.Nu = math.Abs(param[0])
		s.Mu = 0
		s.Sigma = 1
		var pOptimize float64
		for i := 0; i < len(samples); i++ {
			x := (samples[i] - param[1]) / math.Abs(param[2])
			pOptimize += s.LogProb(x) - math.Log(math.Abs(param[2]))
		}
		return -pOptimize
	}
	settings := optimize.DefaultSettingsLocal()
	p := optimize.Problem{
		Func: fOptimize,
	}
	param0 := []float64{nu, mu, sigma}
	if result, err := optimize.Minimize(p, param0, settings, nil); err == nil {
		e.s.Nu = math.Abs(result.X[0])
		e.s.Mu = math.Abs(result.X[1])
		e.s.Sigma = result.X[2]
		return result, nil
	} else {
		return result, err
	}
}

//Compute estimate Laplace distribution parameters
func (e *PDFEstimatorStudentT) Compute() error {
	pdf := e.pdf.(*PDFStudentT)
	if r, err := e.fit(e.samples); err != nil {
		pdf.initialized = false
		pdf.e = err
		e.e = err
		return err
	} else {
		pdf.initialized = true
		if r.Status < optimize.Failure && r.Status > optimize.Success {
			pdf.converged = true
		} else {
			pdf.converged = false
		}
	}
	return nil
}
