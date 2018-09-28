package filters

import (
	"github.com/diadata-org/api-golang/dia"
	"testing"
	"time"
)

func TestFilterMa(t *testing.T) {

	filterParam := 10
	firstPrice := 50.0

	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMA("test", d, filterParam)
	steps := filterParam
	p := firstPrice
	i := 0
	for i <= steps {
		f.compute(&dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(time.Second)
		i += 1
	}
	v := f.filterPoint(d)
	if v.Value != p {
		t.Errorf("error should be stable %v", v)
	}

	priceIncrements := 1.0
	i = 0
	for i <= steps {
		f.compute(&dia.Trade{EstimatedUSDPrice: p, Time: d})
		p = p + priceIncrements
		d = d.Add(time.Second)
		i += 1
	}
	v = f.filterPoint(d)

	if v.Value != firstPrice+(priceIncrements*float64((steps-1))/2) {
		t.Errorf("error shouldnt be %v", v)
	}

}

func TestFilterMa2(t *testing.T) {

	filterParam := 10
	firstPrice := 50.0

	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMA("test", d, filterParam)
	steps := filterParam
	p := firstPrice
	i := 0
	for i <= steps {
		f.compute(&dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(time.Second)
		d = d.Add(time.Second)
		i += 1
	}
	v := f.filterPoint(d)
	if v.Value != p {
		t.Errorf("error should be stable %v", v)
	}

	priceIncrements := 1.0
	i = 0
	for i <= steps {
		f.compute(&dia.Trade{EstimatedUSDPrice: p, Time: d})
		p = p + priceIncrements
		d = d.Add(time.Second)
		d = d.Add(time.Second)
		i += 1
	}
	v = f.filterPoint(d)

	if v.Value != 57.0 { //TODO formulas
		t.Errorf("error shouldnt be 57.0 %v", v)
	}

}
