package filters

import (
	"fmt"
	"math"
	"time"
	"errors"

	scrapers "github.com/diadata-org/diadata/internal/pkg/exchange-scrapers"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// expiration related -------------------------------------

// Eod - gives you t's midnight time
func Eod(t time.Time, timezone string) (time.Time, error) {
	nilTime := time.Time{}
	d, err := time.ParseDuration("24h")
	if err != nil {
		return nilTime, err
	}
	l, err := time.LoadLocation(timezone)
	if err != nil {
		return nilTime, err
	}
	// next day
	year, month, day := t.Add(d).Date()
	return time.Date(year, month, day, 0, 0, 0, 0, l), nil
}

// Bod - gives you today's beginning of day
func Bod(t time.Time, timezone string) (time.Time, error) {
	nilTime := time.Time{}
	l, err := time.LoadLocation(timezone)
	if err != nil {
		return nilTime, err
	}
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, l), nil
}

// MinutesUntilMidnight - how many minutes until midnight
func MinutesUntilMidnight(timezone string) (float64, error) {
	const nilTime float64 = 0
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return nilTime, err
	}
	now := time.Now().In(loc)

	t, err := processMinutesUntilMidnight(now, timezone)
	if err != nil {
		return nilTime, err
	}
	return t, nil
}
func processMinutesUntilMidnight(now time.Time, timezone string) (float64, error) {
	const nilTime float64 = 0
	eod, err := Eod(now, timezone)
	if err != nil {
		return nilTime, err
	}
	// returns the time until eod in eod's location (timezone)
	return eod.Sub(now).Minutes(), nil
}

// MinutesUntilSettlement - how many minutes from midnight until settlement; If on the settlement day, however and after midnight, then count how many minutes until settlement (rather than count how many minutes from midnight until settlement)
func MinutesUntilSettlement(settlement scrapers.OptionSettlement, timezone string) (float64, error) {
	const nilTime float64 = 0
	loc, err := time.LoadLocation(timezone)
	now := time.Now().In(loc)
	mid, err := Eod(now, timezone)
	if err != nil {
		return nilTime, err
	}
	// add more timezones, with different expirations
	switch timezone {
	case "CET":
		switch settlement {
		case scrapers.RegularOptionSettlement:
			// 09:00 AM CET expiration for regular
			d, err := time.ParseDuration("9h")
			if err != nil {
				return nilTime, err
			}
			t := mid.Add(d)
			return t.Sub(mid).Minutes(), nil
		case scrapers.WeeklyOptionSettlement:
			// 05:30 PM CET expiration for weekly
			d, err := time.ParseDuration("17h30m")
			if err != nil {
				return nilTime, err
			}
			t := mid.Add(d)
			return t.Sub(mid).Minutes(), nil
		default:
			return nilTime, fmt.Errorf("unknown option settlement type")
		}
	default:
		return nilTime, fmt.Errorf("unknown timezone provided")
	}
}

// MinutesBetweenTwoDays - given two days: t1 and t2, this method calculates how many minutes there are between the midnight of t1 and the midnight of the day immediately before t2. So this is an exclusive time difference measured in minutes between the two dates. The order of the dates given does not matter.
func MinutesBetweenTwoDays(t1 time.Time, t2 time.Time) (float64, error) {
	const nilTime float64 = 0
	// if the date is the same. not using func (t Time) Before(u Time) bool
	// because the time can be diffrent but for our purposes this is the same day
	if t1.Day() == t2.Day() && t1.Year() == t2.Year() && t1.Month() == t2.Month() {
		return nilTime, nil
	}
	t1Beforet2 := t1.Before(t2)
	// time order invariant
	switch t1Beforet2 {
	case true:
		t1End, err := Eod(t1, "")
		if err != nil {
			return nilTime, err
		}
		t2Begin, err := Bod(t2, "")
		if err != nil {
			return nilTime, err
		}
		mins := t2Begin.Sub(t1End).Minutes()
		return mins, nil
	case false:
		t1End, err := Eod(t2, "")
		if err != nil {
			return nilTime, err
		}
		t2Begin, err := Bod(t1, "")
		if err != nil {
			return nilTime, err
		}
		mins := t2Begin.Sub(t1End).Minutes()
		return mins, nil
	}
	// ! unreachable code. required to silence go compiler
	return nilTime, nil
}

// MinutesInYear - returns how many minutes there were in a year
func MinutesInYear(year int) (float64, error) {
	if year < 0 {
		return 0, fmt.Errorf("negative years not allowed")
	}
	// does not matter what timezone is used
	l, _ := time.LoadLocation("UTC")
	firstD := time.Date(year, 1, 1, 0, 0, 0, 0, l)
	lastD := time.Date(year+1, 1, 1, 0, 0, 0, 0, l)
	return lastD.Sub(firstD).Minutes(), nil
}

// --------------------------------------------------------

func TimeToMaturity(option dia.OptionMetaForward) float64 {
	endtime := option.ExpirationTime
	starttime := time.Now()

	return endtime.Sub(starttime).Minutes()
}

// Get the forward option meta information for near and next term
func GetNextTermOptionMeta(baseCurrency string) ([]dia.OptionMetaForward, error) {
	result := []dia.OptionMetaForward{}
	timeResult := dia.OptionMeta{}
	ds, err := models.NewDataStore()
	if err != nil {
		return result, err
	}

	optionsMeta, err := ds.GetOptionMeta(baseCurrency)
	if err != nil {
		return result, err
	}

	// Create date so late it is out of target range
	earliestDate := time.Now().Add(62 * 24 * time.Hour)
	// Determine date of NextTerm
	for _, optionMeta := range(optionsMeta) {
		if optionMeta.ExpirationTime.Sub(time.Now()) > 62 * 24 * time.Hour {
			continue
		}
		if optionMeta.ExpirationTime.Sub(time.Now()) < 20 * 24 * time.Hour {
			continue
		}
		if optionMeta.ExpirationTime.Unix() < earliestDate.Unix() {
			earliestDate = optionMeta.ExpirationTime
			timeResult = optionMeta
		}
	}
	if timeResult.InstrumentName == "" {
		err = errors.New("No matching instrument found.")
	}
	result, err = generateForwardMeta(ds, timeResult, optionsMeta)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Find last option with expiration date before NextTermOption
func GetNearTermOptionMeta(baseCurrency string, expirationNextTerm time.Time) ([]dia.OptionMetaForward, error) {
	log.Info("GetNearTermOptionMeta")
	timeResult := dia.OptionMeta{}
	result := []dia.OptionMetaForward{}
	ds, err := models.NewDataStore()
	if err != nil {
		return result, err
	}

	optionsMeta, err := ds.GetOptionMeta(baseCurrency)
	if err != nil {
		return result, err
	}
	latestFoundDate := time.Now().Add(1 * time.Second)
	for _, optionMeta := range(optionsMeta) {
		if optionMeta.ExpirationTime.Unix() >= expirationNextTerm.Unix() {
			continue
		}
		if optionMeta.ExpirationTime.Unix() > latestFoundDate.Unix() {
			latestFoundDate = optionMeta.ExpirationTime
			timeResult = optionMeta
		}
	}
	if timeResult.InstrumentName == "" {
		err = errors.New("No matching instrument found.")
	}
	result, err = generateForwardMeta(ds, timeResult, optionsMeta)
	if err != nil {
		return result, err
	}
	return result, nil
	return result, err
}

func generateForwardMeta(ds *models.DB, timeResult dia.OptionMeta, optionsMeta []dia.OptionMeta) ([]dia.OptionMetaForward, error) {
	result := []dia.OptionMetaForward{}
	type OptionMetaOrderbook struct {
		orderbookDataPut  dia.OptionOrderbookDatum
		orderbookDataCall dia.OptionOrderbookDatum
		optionMetaPut     dia.OptionMeta
		optionMetaCall    dia.OptionMeta
	}

	orderbookDataStore := make(map[string]OptionMetaOrderbook)

	// Add all options that expire at the same date into one object of type OptionMetaOrderbook
	for _, optionMeta := range(optionsMeta) {
		if optionMeta.ExpirationTime.Equal(timeResult.ExpirationTime) {
			orderbookData, err := ds.GetOptionOrderbookDataInflux(optionMeta)
			if err != nil {
				return result, err
			}
			if orderbookData.InstrumentName != "" {
				generalizedInstrumentName := optionMeta.InstrumentName[:16]
				switch optionMeta.OptionType {
				case dia.CallOption:
					if val, ok := orderbookDataStore[generalizedInstrumentName]; ok {
						orderbookDataStore[generalizedInstrumentName] = OptionMetaOrderbook {
							orderbookDataCall: orderbookData,
							optionMetaCall:    optionMeta,
							orderbookDataPut:  val.orderbookDataPut,
							optionMetaPut:     val.optionMetaPut,
						}
					} else {
						orderbookDataStore[generalizedInstrumentName] = OptionMetaOrderbook {
							orderbookDataCall: orderbookData,
							optionMetaCall:    optionMeta,
						}
					}
				case dia.PutOption:
					if val, ok := orderbookDataStore[generalizedInstrumentName]; ok {
						orderbookDataStore[generalizedInstrumentName] = OptionMetaOrderbook {
							orderbookDataPut:  orderbookData,
							optionMetaPut:     optionMeta,
							orderbookDataCall: val.orderbookDataCall,
							optionMetaCall:    val.optionMetaCall,
						}
					} else {
						orderbookDataStore[generalizedInstrumentName] = OptionMetaOrderbook {
							orderbookDataPut: orderbookData,
							optionMetaPut:    optionMeta,
						}
					}
				}
			}
		}
	}

	for _, orderbookData := range orderbookDataStore {
		result = append(result, dia.OptionMetaForward {
			ExpirationTime: orderbookData.optionMetaCall.ExpirationTime,
			StrikePrice: orderbookData.optionMetaCall.StrikePrice,
			CallPrice:   math.Abs(orderbookData.orderbookDataCall.AskPrice - orderbookData.orderbookDataCall.BidPrice),
			PutPrice:    math.Abs(orderbookData.orderbookDataPut.AskPrice - orderbookData.orderbookDataPut.BidPrice),
		})
	}
	return result, nil
}

// variance / index - calculation related

// VarianceIndex is used to calculate variance for near term and next term options; later on these two values are used in interpolation to obtain a CVI value; r - risk free rate; t - time to expiration; f - forward index level; k0 - strike price just below the forward index level; LaTeX equation for the output of this function is: \sigma^2_j = \frac{1}{T_j} \left(2 \sum_i \frac{\Delta K_i}{K_i^2} \exp{(RT_j)} \cdot Q(K_i) - \left( \frac{F_j}{K_0} - 1 \right)^2 \right), \forall \ j \in \{1,2\}
func VarianceIndex(optionsMeta []dia.OptionMetaIndex, r float64, t float64, f float64, k0 float64) (float64, error) {
	if len(optionsMeta) <= 3 {
		return 0, fmt.Errorf("not enough options to compute the CVI")
	}

	var (
		lh     float64 = 0
		deltaK float64 = 0
	)

	// left & right hand side terms
	// as explained in the issue on GitHub: https://github.com/diadata-org/diadata/issues/193
	// it is not possible to calculate \Delta K_i for the option with the lowest and the option
	// with the highest strike. They are not included in the calculation
	for ix, option := range optionsMeta {
		if ix == 0 || ix == len(optionsMeta)-1 {
			continue
		}
		deltaK = ((optionsMeta[ix-1].StrikePrice + optionsMeta[ix+1].StrikePrice) / 2)
		lh += deltaK * ((option.AskPrice + option.BidPrice) / 2)
	}

	return (1 / t) * (2*math.Exp(r*t)*lh - math.Pow(f/k0-1, 2)), nil
}

// ForwardIndexLevel calculates the forward level; used to compute the forward level for near-term & next-term options; r - risk free rate; t - time to expiration; LaTeX equation for the forward index level is: F_j = \texttt{Strike Price}_j + \exp{(R_j T)} \cdot (\texttt{Call Price}_j - \texttt{Put Price}_j)
func ForwardIndexLevel(optionsMeta []dia.OptionMetaForward, r float64, t float64) (float64, error) {
	if len(optionsMeta) <= 1 {
		return 0, fmt.Errorf("not enough data points to compute the forward level")
	}
	minAbsolutePriceDiff := math.Abs(optionsMeta[0].CallPrice - optionsMeta[0].PutPrice)
	minAbsolutePriceDiffIx := 0

	for ix, optionMeta := range optionsMeta {
		if ix == 0 {
			continue
		}
		absDiff := math.Abs(optionMeta.CallPrice - optionMeta.PutPrice)
		if absDiff < minAbsolutePriceDiff {
			minAbsolutePriceDiff = absDiff
			minAbsolutePriceDiffIx = ix
		}
	}

	option := optionsMeta[minAbsolutePriceDiffIx]
	return option.StrikePrice + math.Exp(r*t)*(option.CallPrice-option.PutPrice), nil
}

// CVI computes the crypto volatility index
func CVI(sigma1 float64, sigma2 float64, t1 float64, t2 float64, year int) (float64, error) {
	const (
		nilCVI float64 = 0
		N30    float64 = 43_200
		N365   float64 = 525_600
	)

	minutesInCalculationYear, err := MinutesInYear(year)
	if err != nil {
		return nilCVI, err
	}

	var (
		Nt1 float64 = t1 * minutesInCalculationYear
		Nt2 float64 = t2 * minutesInCalculationYear
	)

	if (Nt2 - Nt1) == 0 {
		return nilCVI, fmt.Errorf("t2 has to be later than t1")
	}

	w1 := (Nt2 - N30) / (Nt2 - Nt1)
	w2 := (N30 - Nt1) / (Nt2 - Nt1)

	cvi := 100 * math.Sqrt((t1*math.Pow(sigma1, 2)*w1+t2*math.Pow(sigma2, 2)*w2)*(N365/N30))
	return cvi, nil
}

// CVIFiltering is the actual filtering algorithm; computedCVIs is the channel through which we receive the calculated CVIs, filteredCVIs is the channel through which we send the filtered CVIs
func CVIFiltering(computedCVIs scrapers.ComputedCVIs, filteredCVIs chan<- scrapers.ComputedCVI) {
	// it is the responsibility of the function that filters the CVIs to close the channel through which it communicates these values
	defer close(filteredCVIs)
	var baseline float64 = 0
	var lastCVItime time.Time = time.Time{}
	var absDiff float64 = 0
	var noChangeCVItime time.Duration = 0 // tracks for how long there has been no change to CVI

	// it is the responsibility of thr computedCVIs creator to send the computed CVI values as often as it needs to
	for v := range computedCVIs {
		// first CVI value becomes baseline
		if baseline == 0 {
			baseline = v.CVI
			lastCVItime = v.CalculationTime
			filteredCVIs <- scrapers.ComputedCVI{CVI: baseline, CalculationTime: time.Now()}
			continue
		}
		// if no changes for 2 minutes or more
		noChangeCVItime = time.Now().Sub(lastCVItime)
		if noChangeCVItime.Minutes() >= 2 {
			baseline = v.CVI
			lastCVItime = v.CalculationTime
			filteredCVIs <- scrapers.ComputedCVI{CVI: baseline, CalculationTime: time.Now()}
			continue
		}
		// ignoring or updating the CVI logic
		absDiff = math.Abs(v.CVI - baseline)
		if absDiff <= 0.49 {
			baseline = v.CVI
			lastCVItime = v.CalculationTime
			filteredCVIs <- scrapers.ComputedCVI{CVI: baseline, CalculationTime: time.Now()}
			continue
		}
		if absDiff > 0.49 {
			// do not update the CVI time so that we can check whether we have had the baseline for two minutes or more
			filteredCVIs <- scrapers.ComputedCVI{CVI: baseline, CalculationTime: time.Now()}
			continue
		}
	}
}
