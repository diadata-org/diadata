package dia

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zekroTJA/timedmap"
)

// GetBaseToken returns the base token of a trading pair
// TO DO (20/11/2020): This method is no longer needed once we switch to new Token/Trade structs
func (t *Trade) GetBaseToken() string {

	if t.BaseToken.Symbol != "" {
		return t.BaseToken.Symbol
	}

	pair := strings.ToUpper(t.Pair)
	if len(pair) > 3 {
		switch t.Source {
		case KrakenExchange:
			if pair[len(pair)-3:] == "XBT" {
				return "BTC"
			}
		case BitfinexExchange:
			if pair[len(pair)-3:] == "USD" {
				return "USDT"
			}
		case HitBTCExchange:
			if pair[len(pair)-3:] == "USD" {
				return "USDT"
			}
		}
	}

	second := strings.TrimPrefix(pair, strings.ToUpper(t.Symbol)+"_")
	if second != pair {
		return second
	}
	second = strings.TrimPrefix(pair, strings.ToUpper(t.Symbol)+"-")
	if second != pair {
		return second
	}
	second = strings.TrimPrefix(pair, strings.ToUpper(t.Symbol)+"/")
	if second != pair {
		return second
	}

	return strings.TrimPrefix(pair, strings.ToUpper(t.Symbol))
}

func SwapTradeSeparator(t Trade, separator string) (Trade, error) {
	if t.Price == 0 {
		return t, errors.New("zero price. cannot swap trade")
	}
	t.BaseToken, t.QuoteToken = t.QuoteToken, t.BaseToken
	t.Symbol = t.QuoteToken.Symbol
	t.Pair = t.QuoteToken.Symbol + separator + t.BaseToken.Symbol
	t.Volume = -t.Price * t.Volume
	t.Price = 1 / t.Price

	return t, nil
}

// SwapTrade swaps base and quote token of a trade and inverts the price accordingly
func SwapTrade(t Trade) (Trade, error) {
	return SwapTradeSeparator(t, "-")
}

// IdentifyDuplicateFull returns true in case a trade is fully identical to one stored in the timed map @falseDuplicateTrades.
func (t *Trade) IdentifyDuplicateFull(falseDuplicateTrades *timedmap.TimedMap, memory time.Duration) (discardTrade bool) {
	if _, ok := falseDuplicateTrades.GetValue(t.TradeIdentifierFull()).(int); !ok {
		falseDuplicateTrades.Set(t.TradeIdentifierFull(), 1, memory)
	} else {
		discardTrade = true
	}
	return
}

// IdentifyDuplicateTagset identifies trades with identical timestamps and tagsets and add a Nanosecond to
// the timestamp in order for the trade not to be overwritten in Influx.
func (t *Trade) IdentifyDuplicateTagset(duplicateTrades *timedmap.TimedMap, memory time.Duration) {
	if val, ok := duplicateTrades.GetValue(t.TradeIdentifierTagset()).(int); !ok {
		duplicateTrades.Set(t.TradeIdentifierTagset(), 1, memory)
	} else {
		// Set trade identifier as key with value incremented.
		duplicateTrades.Set(t.TradeIdentifierTagset(), val+1, memory)
		// Add old value as Nanosecond to trade time.
		t.Time = t.Time.Add(time.Duration(val) * time.Nanosecond)
	}
}

// TradeIdentifierFull returns an identifier with respect to all fields of a trade.
func (t *Trade) TradeIdentifierFull() string {
	timeString := strconv.Itoa(int(t.Time.UnixNano()))
	priceString := fmt.Sprintf("%f", t.Price)
	volumeString := fmt.Sprintf("%f", t.Volume)
	return timeString +
		priceString +
		volumeString +
		t.ForeignTradeID +
		t.Source +
		t.QuoteToken.Address +
		t.QuoteToken.Blockchain +
		t.BaseToken.Address +
		t.BaseToken.Blockchain
}

// TradeIdentifierTagset returns an identifier with respect to the tagset of a trade in Influx.
// In other words, a trade with this same tagset is overwritten in Influx trades table.
func (t *Trade) TradeIdentifierTagset() string {
	timeString := strconv.Itoa(int(t.Time.UnixNano()))
	return timeString +
		t.Symbol +
		t.Pair +
		t.Source +
		t.QuoteToken.Address +
		t.QuoteToken.Blockchain +
		t.BaseToken.Address +
		t.BaseToken.Blockchain
}
