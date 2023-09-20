package filters

import (
	"encoding/json"
	"math"
	"testing"

	"github.com/diadata-org/diadata/pkg/dia"
)

func getTrades() (trades []dia.Trade) {
	json.Unmarshal([]byte(jsonTrades), &trades)
	trades = trades[0:3]
	return
}

func TestVWAP(t *testing.T) {
	var filterPoints []dia.FilterPoint
	trades := getTrades()
	maFilter := NewFilterVWAP(dia.Asset{}, "Binance", trades[len(trades)-1].Time, dia.BlockSizeSeconds, false)

	totalVolume := 0.0
	totalPrice := 0.0

	for _, trade := range trades {
		totalVolume = totalVolume + math.Abs(trade.Volume)
		totalPrice = totalPrice + (trade.EstimatedUSDPrice * math.Abs(trade.Volume))
		maFilter.Compute(trade)
	}

	expectedAns := totalPrice / totalVolume

	maFilter.FinalCompute(trades[0].Time)
	fp := maFilter.FilterPointForBlock()
	filterPoints = append(filterPoints, *fp)

	if filterPoints[0].Value != expectedAns {
		t.Errorf("Error vwap expected %v  and got %v ", expectedAns, filterPoints[0].Value)
	}

}

//BenchmarkVWAP-8   	 1438806	       828.6 ns/op	    1104 B/op	      21 allocs/op
//BenchmarkVWAP-8   	  343592	      3280 ns/op	    5011 B/op	      84 allocs/op

func BenchmarkVWAP(b *testing.B) {

	trades := getTrades()
	for n := 0; n < b.N; n++ {
		maFilter := NewFilterVWAP(dia.Asset{}, "Binance", trades[len(trades)-1].Time, dia.BlockSizeSeconds, false)

		for _, trade := range trades {

			maFilter.Compute(trade)
		}

		maFilter.FinalCompute(trades[0].Time)
	}
}
