package filters

import (
	"testing"

	"github.com/diadata-org/diadata/pkg/dia"
)

func TestVWAPIR(t *testing.T) {
	var filterPoints []dia.FilterPoint
	trades := getTrades()
	maFilter := NewFilterVWAPIR(dia.Asset{}, "Binance", trades[len(trades)-1].Time, dia.BlockSizeSeconds, false)

	for _, trade := range trades {

		maFilter.Compute(trade)
	}

	maFilter.FinalCompute(trades[0].Time)
	fp := maFilter.FilterPointForBlock()
	filterPoints = append(filterPoints, *fp)

	if filterPoints[0].Value != 39873.659462014075 {
		t.Errorf("Error vwap expected %v  and got %v ", 39873.659462014075, filterPoints[0].Value)
	}

}

//BenchmarkVWAPIR-8   	  332109	      3332 ns/op	    5011 B/op	      84 allocs/op

func BenchmarkVWAPIR(b *testing.B) {

	trades := getTrades()
	for n := 0; n < b.N; n++ {
		maFilter := NewFilterVWAPIR(dia.Asset{}, "Binance", trades[len(trades)-1].Time, dia.BlockSizeSeconds, false)

		for _, trade := range trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(trades[0].Time)
	}
}
