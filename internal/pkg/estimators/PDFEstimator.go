package estimators

//PDF Provability Density Function
type PDF interface {
	DidConverge() (bool, error)
	DidGetSolution() (bool, error)
	GetName() string
	GetError() error
}

//PDFEstimator estimate a PDF based on samples
type PDFEstimator interface {
	GetPDF() PDF
	AddSamples(samples []float64)
	Compute() error
}

const (
	Laplace = "Laplace"
)

//PDFEstimators gives a list of the available PDF estimators
func PDFEstimators() []string {
	return []string{
		Laplace,
	}
}

//NewPDFEstimator returns a PDFEstimator
func NewPDFEstimator(estimator string) PDFEstimator {
	switch estimator {
	case Laplace:
		return NewPDFEstimatorLaplace()
	default:
		return nil
	}
}
