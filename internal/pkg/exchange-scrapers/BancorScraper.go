package scrapers

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"sync"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/jjjjpppp/quoinex-go-client/v2"
)

const (
	// Set delay to one day and divide by number of pairs. In this way, we can
	// use 24h volume as averaged trade volume
	BancorApiDelay = 60 * 60 * 24
)

type BancorTicker struct {
	Data struct {
		Name            string  `json:"name"`
		Symbol          string  `json:"symbol"`
		Code            string  `json:"code"`
		Decimals        int     `json:"decimals"`
		TotalSupply     string  `json:"totalSupply"`
		TotalSupplyD    string  `json:"totalSupplyD"`
		Volume24H       string  `json:"volume24h"`
		Volume24HD      float64 `json:"volume24hD"`
		DisplayCurrency string  `json:"displayCurrency"`
		Price24H        string  `json:"price24h"`
		Price           float64 `json:"price"`
	} `json:"data"`
}

type BancorAssetPairs struct {
	Data struct {
		ETH     string `json:"ETH"`
		EOS     string `json:"EOS"`
		STX     string `json:"STX"`
		BNB     string `json:"BNB"`
		BAT     string `json:"BAT"`
		OMG     string `json:"OMG"`
		DICE    string `json:"DICE"`
		BLACK   string `json:"BLACK"`
		IQ      string `json:"IQ"`
		MANA    string `json:"MANA"`
		SRN     string `json:"SRN"`
		LOC     string `json:"LOC"`
		MKR     string `json:"MKR"`
		ENJ     string `json:"ENJ"`
		SNT     string `json:"SNT"`
		GTO     string `json:"GTO"`
		HORUS   string `json:"HORUS"`
		POWR    string `json:"POWR"`
		EMCO    string `json:"EMCO"`
		KNC     string `json:"KNC"`
		XDCE    string `json:"XDCE"`
		XPAT    string `json:"XPAT"`
		PLR     string `json:"PLR"`
		EPRA    string `json:"EPRA"`
		AMN     string `json:"AMN"`
		MEV     string `json:"MEV"`
		MEETONE string `json:"MEETONE"`
		CEEK    string `json:"CEEK"`
		MYB     string `json:"MYB"`
		TKN     string `json:"TKN"`
		X8X     string `json:"X8X"`
		SPD     string `json:"SPD"`
		XNK     string `json:"XNK"`
		AGRI    string `json:"AGRI"`
		REAL    string `json:"REAL"`
		RBLX    string `json:"RBLX"`
		ONG     string `json:"ONG"`
		ZIPT    string `json:"ZIPT"`
		CAN     string `json:"CAN"`
		WAND    string `json:"WAND"`
		DTRC    string `json:"DTRC"`
		OCT     string `json:"OCT"`
		VEE     string `json:"VEE"`
		POA20   string `json:"POA20"`
		REQ     string `json:"REQ"`
		MFT     string `json:"MFT"`
		RVT     string `json:"RVT"`
		ELF     string `json:"ELF"`
		RDN     string `json:"RDN"`
		ANT     string `json:"ANT"`
		WINGS   string `json:"WINGS"`
		EURS    string `json:"EURS"`
		GNO     string `json:"GNO"`
		AID     string `json:"AID"`
		TNS     string `json:"TNS"`
		TAEL    string `json:"TAEL"`
		GRID    string `json:"GRID"`
		ZINC    string `json:"ZINC"`
		HVT     string `json:"HVT"`
		SVD     string `json:"SVD"`
		MRG     string `json:"MRG"`
		LOCI    string `json:"LOCI"`
		BETR    string `json:"BETR"`
		LDC     string `json:"LDC"`
		MFG     string `json:"MFG"`
		AUC     string `json:"AUC"`
		BOXX    string `json:"BOXX"`
		DRT     string `json:"DRT"`
		WLK     string `json:"WLK"`
		RLC     string `json:"RLC"`
		RCN     string `json:"RCN"`
		DATA    string `json:"DATA"`
		MRPH    string `json:"MRPH"`
		IND     string `json:"IND"`
		CAT     string `json:"CAT"`
		LUME    string `json:"LUME"`
		VIB     string `json:"VIB"`
		SCL     string `json:"SCL"`
		REM     string `json:"REM"`
		AIX     string `json:"AIX"`
		TBX     string `json:"TBX"`
		STAC    string `json:"STAC"`
		ABX     string `json:"ABX"`
		COT     string `json:"COT"`
		PEOS    string `json:"PEOS"`
		FTX     string `json:"FTX"`
		ZOS     string `json:"ZOS"`
		BCS     string `json:"BCS"`
		REF     string `json:"REF"`
		J8T     string `json:"J8T"`
		FLIXX   string `json:"FLIXX"`
		EQUA    string `json:"EQUA"`
		MNTP    string `json:"MNTP"`
		MTL     string `json:"MTL"`
		MDT     string `json:"MDT"`
		SXL     string `json:"SXL"`
		XBP     string `json:"XBP"`
		SAN     string `json:"SAN"`
		UP      string `json:"UP"`
		DAPP    string `json:"DAPP"`
		EOSDT   string `json:"EOSDT"`
		FINX    string `json:"FINX"`
		EMT     string `json:"EMT"`
		NUT     string `json:"NUT"`
		EFOOD   string `json:"EFOOD"`
		HEDG    string `json:"HEDG"`
		PIXEOS  string `json:"PIXEOS"`
		CHEX    string `json:"CHEX"`
		USDQ    string `json:"USDQ"`
		QDAO    string `json:"QDAO"`
		ELET    string `json:"ELET"`
		AMPL    string `json:"AMPL"`
		USDB    string `json:"USDB"`
		SENSE   string `json:"SENSE"`
		USDT    string `json:"USDT"`
		NMR     string `json:"NMR"`
		MET     string `json:"MET"`
		CUSD    string `json:"CUSD"`
		UBT     string `json:"UBT"`
		ANK     string `json:"ANK"`
		DAI     string `json:"DAI"`
		GRIG    string `json:"GRIG"`
		UPT     string `json:"UPT"`
		RPL     string `json:"RPL"`
		PBTC    string `json:"PBTC"`
		JRT     string `json:"JRT"`
		STM     string `json:"STM"`
		INVOX   string `json:"INVOX"`
		XRT     string `json:"XRT"`
	} `json:"data"`
}

type BancorScraper struct {
	client       *quoinex.Client
	exchangeName string

	// channels to signal events
	run          bool
	initDone     chan nothing
	shutdown     chan nothing
	shutdownDone chan nothing

	errorLock sync.RWMutex
	error     error
	closed    bool

	pairScrapers   map[string]*BancorPairScraper
	productPairIds map[string]int
	chanTrades     chan *dia.Trade
}

func NewBancorScraper(exchangeName string) *BancorScraper {
	scraper := &BancorScraper{
		exchangeName:   exchangeName,
		initDone:       make(chan nothing),
		shutdown:       make(chan nothing),
		shutdownDone:   make(chan nothing),
		productPairIds: make(map[string]int),
		pairScrapers:   make(map[string]*BancorPairScraper),
		chanTrades:     make(chan *dia.Trade),
	}

	go scraper.mainLoop()
	return scraper
}

func (scraper *BancorScraper) mainLoop() {
	scraper.run = true
	pairs, _ := scraper.FetchAvailablePairs()
	numPairs := len(pairs)
	for scraper.run {
		if len(scraper.pairScrapers) == 0 {
			scraper.error = errors.New("bancor: No pairs to scrap provided")
			log.Error(scraper.error.Error())
			break
		}
		for pair, pairScraper := range scraper.pairScrapers {
			// Sleep such that each pair is scraped once per day
			time.Sleep(time.Duration(BancorApiDelay/numPairs) * time.Second)

			var ticker BancorTicker
			tickerData, err := utils.GetRequest("https://api.bancor.network/0.1/currencies/" + pairScraper.pair.Symbol + "/ticker?fromCurrencyCode=BNT")
			if err != nil {
				log.Error("Error getting ticker: ", err)
				continue
			}
			err = json.Unmarshal(tickerData, &ticker)
			if err != nil {
				log.Error("Error unmarshalling ticker: ", err)
				continue
			}

			price, err := strconv.ParseFloat(ticker.Data.Price24H, 64)
			if err != nil {
				log.Error("error parsing price24H: ", err)
			}

			trade := &dia.Trade{
				Symbol:         pairScraper.pair.Symbol,
				Pair:           pair,
				Price:          price,
				Volume:         ticker.Data.Volume24HD,
				Time:           time.Now(),
				ForeignTradeID: "",
				Source:         scraper.exchangeName,
			}
			log.Info("Got Trade: ", trade)

			pairScraper.parent.chanTrades <- trade
		}
	}
	if scraper.error == nil {
		scraper.error = errors.New("Main loop terminated by Close().")
	}
	scraper.cleanup(nil)
}

func (scraper *BancorScraper) FetchAvailablePairs() (pairs []dia.Pair, err error) {
	assets, err := scraper.readAssets()
	if err != nil {
		log.Error("Couldn't obtain Bancor product ids:", err)
	}
	v := reflect.ValueOf(assets.Data)
	typeOfS := v.Type()
	pairs = make([]dia.Pair, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		pairs = append(pairs, dia.Pair{
			Symbol:      typeOfS.Field(i).Name,
			ForeignName: typeOfS.Field(i).Name + "-" + v.Field(i).Interface().(string),
			Exchange:    scraper.exchangeName,
		})
	}
	return
}

func (scraper *BancorScraper) readAssets() (BancorAssetPairs, error) {
	var pair BancorAssetPairs

	pairs, err := utils.GetRequest("https://api.bancor.network/0.1/currencies/convertiblePairs")
	if err != nil {
		return pair, err
	}
	err = json.Unmarshal(pairs, &pair)
	if err != nil {
		log.Error("Error reading json", err)

	}
	return pair, nil

}

func (scraper *BancorScraper) ScrapePair(pair dia.Pair) (PairScraper, error) {
	scraper.errorLock.RLock()
	defer scraper.errorLock.RUnlock()

	if scraper.error != nil {
		return nil, scraper.error
	}

	if scraper.closed {
		return nil, errors.New("BancorScraper is closed")
	}

	pairScraper := &BancorPairScraper{
		parent: scraper,
		pair:   pair,
	}

	scraper.pairScrapers[pair.ForeignName] = pairScraper

	return pairScraper, nil
}
func (scraper *BancorScraper) cleanup(err error) {
	scraper.errorLock.Lock()
	defer scraper.errorLock.Unlock()
	if err != nil {
		scraper.error = err
	}
	scraper.closed = true
	close(scraper.shutdownDone)
}

func (scraper *BancorScraper) Close() error {
	// close the pair scraper channels
	scraper.run = false
	for _, pairScraper := range scraper.pairScrapers {
		pairScraper.closed = true
	}

	close(scraper.shutdown)
	<-scraper.shutdownDone
	return nil
}

type BancorPairScraper struct {
	parent *BancorScraper
	pair   dia.Pair
	closed bool
}

func (pairScraper *BancorPairScraper) Pair() dia.Pair {
	return pairScraper.pair
}

func (scraper *BancorScraper) Channel() chan *dia.Trade {
	return scraper.chanTrades
}

func (pairScraper *BancorPairScraper) Error() error {
	s := pairScraper.parent
	s.errorLock.RLock()
	defer s.errorLock.RUnlock()
	return s.error
}

func (pairScraper *BancorPairScraper) Close() error {
	pairScraper.parent.errorLock.RLock()
	defer pairScraper.parent.errorLock.RUnlock()
	pairScraper.closed = true
	return nil
}
