package estimators

//PDF Provability Density Function
type PDF interface {
	DidConverge() bool
	DidGetSolution() bool
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
	//Laplace distribution name
	Laplace = "Laplace"
	//Student distribution name
	Student = "Student"
)

//PDFEstimators gives a list of the available PDF estimators
func PDFEstimators() []string {
	return []string{
		Laplace,
		Student,
	}
}

//NewPDFEstimator returns a PDFEstimator
func NewPDFEstimator(estimator string) PDFEstimator {
	switch estimator {
	case Laplace:
		return NewPDFEstimatorLaplace()
	case Student:
		return NewPDFEstimatorStudentT()
	default:
		return nil
	}
}
