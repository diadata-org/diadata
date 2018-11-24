package estimators

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat/distuv"
	"testing"
	"time"
)

func TestShutdown(t *testing.T) {
	e := NewEstimatorsService(nil)
	time.Sleep(time.Second)
	e.Close()
}

func TestCreateEstimator(t *testing.T) {
	e := NewEstimatorsService(nil)
	e.CreateEstimator(PDFEstimators()[0], "symbol", "exchange")
	if len(e.estimators) != 1 {
		t.Errorf("Estimator was not added")
	}
	e.CreateEstimator("Estimator not implemented", "symbol", "exchange")
	if len(e.estimators) != 1 {
		t.Errorf("Estimator was added")
	}
	e.Close()
}

type distribution interface {
	Rand() float64
}

func TestProcessTradesBlock(t *testing.T) {
	cases := []struct {
		nSamples       int
		wantParameters []float64
		estimator      string
	}{
		{
			nSamples:       10000,
			wantParameters: []float64{1, 0, 1},
			estimator:      "Student",
		},
		{
			nSamples:       100000,
			wantParameters: []float64{10, 1},
			estimator:      "Laplace",
		},
	}
	for i, test := range cases {
		c := make(chan PDF)
		exchange := "none"
		symbol := "test"
		src := rand.New(rand.NewSource(1))
		trades := dia.TradesBlock{}
		trades.TradesBlockData.Trades = make([]dia.Trade, test.nSamples)
		var d distribution
		switch test.estimator {
		case "Laplace":
			d = distuv.Laplace{
				Mu:    test.wantParameters[0],
				Scale: test.wantParameters[1],
				Src:   src,
			}
		case "Student":
			d = distuv.StudentsT{
				Nu:    test.wantParameters[0],
				Mu:    test.wantParameters[1],
				Sigma: test.wantParameters[2],
				Src:   src,
			}
		}
		for i := range trades.TradesBlockData.Trades {
			trades.TradesBlockData.Trades[i] = dia.Trade{
				Symbol:            symbol,
				Source:            exchange,
				EstimatedUSDPrice: d.Rand(),
			}
		}
		s := NewEstimatorsService(c)
		s.CreateEstimator(test.estimator, symbol, exchange)
		s.ProcessTradesBlock(&trades)
		pdf := <-c
		gotParameter := make([]float64, len(test.wantParameters))
		switch test.estimator {
		case "Laplace":
			gotParameter[0] = pdf.(*PDFLaplace).GetLocation()
			gotParameter[1] = pdf.(*PDFLaplace).GetScale()
		case "Student":
			gotParameter[0] = pdf.(*PDFStudentT).GetNu()
			gotParameter[1] = pdf.(*PDFStudentT).GetLocation()
			gotParameter[2] = pdf.(*PDFStudentT).GetScale()
		}

		if !pdf.DidGetSolution() {
			t.Errorf("Estimator did not get solution for test %d", i)
		}
		if !pdf.DidConverge() {
			t.Errorf("Estimator did not converged for test %d", i)
		}
		for j := range test.wantParameters {
			if !floats.EqualWithinAbsOrRel(gotParameter[j], test.wantParameters[j], 1e-1, 1e-1) {
				t.Errorf("unexpected parameter:%d result for test :%d got:%f, want:%f", j, i, gotParameter[j], test.wantParameters[j])
			}
		}
	}
}
