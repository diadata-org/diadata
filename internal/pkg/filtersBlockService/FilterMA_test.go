package filters

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"testing"
	"time"
)

func TestFilterMa(t *testing.T) {

	filterParam := 4
	firstPrice := 50.0

	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMA("XRP", "", d, filterParam)
	steps := filterParam
	p := firstPrice
	i := 0
	for i <= steps {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(time.Second)
		i += 1
	}
	f.finalCompute(d)
	v := f.filterPointForBlock()
	if v.Value != p {
		t.Errorf("error should be stable %v", v)
	}

	priceIncrements := 1.0
	i = 0
	for i <= steps {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		p = p + priceIncrements
		d = d.Add(time.Second)
		i += 1
	}
	f.finalCompute(d)
	v = f.filterPointForBlock()
	if v.Value != 53.25 { //TODO formulas
		t.Errorf("error should be, %v", v)
	}
}

func TestFilterMa2(t *testing.T) {

	filterParam := 10
	firstPrice := 50.0

	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMA("XRP", "", d, filterParam)
	steps := filterParam
	p := firstPrice
	i := 0
	for i <= steps {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(time.Second)
		d = d.Add(time.Second)
		i += 1
	}
	v := f.finalCompute(d)
	if v != p {
		t.Errorf("error should be stable %v", v)
	}

	priceIncrements := 1.0
	i = 0
	for i <= steps {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		p = p + priceIncrements
		d = d.Add(time.Second)
		d = d.Add(time.Second)
		i += 1
	}
	v = f.finalCompute(d)
	if v != 56.4 { //TODO formulas
		t.Errorf("error shouldnt be 57.0 %v", v)
	}

}

func TestFilterMaIgnore(t *testing.T) {

	filterParam := 10
	firstPrice := 50.0

	d := time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC)
	f := NewFilterMA("XRP", "", d, filterParam)
	steps := filterParam
	p := firstPrice
	priceIncrements := 1.0
	for i := 0; i <= steps; i++ {
		f.compute(dia.Trade{EstimatedUSDPrice: p, Time: d})
		d = d.Add(-time.Second)
		p += priceIncrements
	}
	v := f.finalCompute(d)
	if v != firstPrice {
		t.Errorf("error should be initial value:%f got:%f", firstPrice, v)
	}
}
