package indexBlockService

import (
	"errors"
	"github.com/cnf/structhash"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/jinzhu/now"
	log "github.com/sirupsen/logrus"
	"math"
	"sort"
	"sync"
	"time"
)

var (
	filterUsed                 = "MA120"
	magicNumber                = 2676316995.365064
	numberTokenettesInOneToken = 10000.0
	usdPerPointsOfIndex        = 0.1
)

// values in second in the past for which we calculate something
// max value a the end !

type nothing struct{}

type IndexBlockService struct {
	shutdown            chan nothing
	shutdownDone        chan nothing
	chanFiltersBlock    chan *dia.FiltersBlock
	chanIndexBlock      chan *dia.IndexBlock
	chanIndexBlock2     chan *dia.IndexBlock
	chanIndexBlockDaily chan *dia.IndexBlock
	chanSuppliesBlock   chan *dia.SuppliesBlock
	//chanVolatilityBlock chan *priceCollection.VolatilityBlock
	errorLock           sync.RWMutex
	error               error
	closed              bool
	points              map[string]dia.FilterPoint
	suppliesBlock       *dia.SuppliesBlock
	supplies            map[string]dia.Supply
	//volatilityBlock     *priceCollection.VolatilityBlock
	filterNameUsed      string
	lastUpdate          time.Time
	lastDailyUpdate     time.Time
}

func (s *IndexBlockService) shouldGenerateNewBlock() bool {

	/*
	   	now.BeginningOfDay()           // 2013-11-18 00:00:00 Mon
	   now.BeginningOfWeek()          // 2013-11-17 00:00:00 Sun
	   now.BeginningOfMonth()         // 2013-11-01 00:00:00 Fri
	*/
	// return true

	result := false
	t := now.BeginningOfMonth()
	if t.Month() != s.lastUpdate.Month() {
		result = true
		s.lastUpdate = t
	}
	log.Info("Should update:", result)
	return result
}

/*func (s *IndexBlockService) shouldGenerateNewDailyBlock() bool {
	result := false
	t := now.BeginningOfDay()
	if t.Day() != s.lastDailyUpdate.Day() {
		result = true
		s.lastDailyUpdate = t
	}
	log.Info("Should update:", result)
	return result
}*/

func (s *IndexBlockService) ChanIndexBlock() chan *dia.IndexBlock {
	return s.chanIndexBlock
}

func (s *IndexBlockService) ChanIndexBlock2() chan *dia.IndexBlock {
	return s.chanIndexBlock2
}

func (s *IndexBlockService) ChanIndexBlockDaily() chan *dia.IndexBlock {
	return s.chanIndexBlockDaily
}

func (s *IndexBlockService) processSupplyBlock(sb *dia.SuppliesBlock) {
	s.suppliesBlock = sb
	s.supplies = make(map[string]dia.Supply)
	for _, su := range sb.BlockData.Supplies {
		s.supplies[su.Symbol] = su
	}
}

func NewIndexBlockService(sb *dia.SuppliesBlock) *IndexBlockService {

	var filter = filterUsed

	s := &IndexBlockService{
		shutdown:            make(chan nothing),
		shutdownDone:        make(chan nothing),
		chanFiltersBlock:    make(chan *dia.FiltersBlock),
		chanIndexBlock:      make(chan *dia.IndexBlock),
		chanIndexBlock2:     make(chan *dia.IndexBlock),
		chanIndexBlockDaily: make(chan *dia.IndexBlock),
		chanSuppliesBlock:   make(chan *dia.SuppliesBlock),
		//chanVolatilityBlock: make(chan *priceCollection.VolatilityBlock),
		error:               nil,
		points:              make(map[string]dia.FilterPoint),
		suppliesBlock:       sb,
		supplies:            nil,
		//volatilityBlock:     vb,
		filterNameUsed:      filter,
	}

	if sb != nil {
		s.processSupplyBlock(s.suppliesBlock)
	}
	s.shouldGenerateNewBlock() // used to init
	//s.shouldGenerateNewDailyBlock() // used to init
	go s.mainLoop()
	return s
}

/*func (s *IndexBlockService) ProcessVolatilityBlock(vb *priceCollection.VolatilityBlock) {
	s.chanVolatilityBlock <- vb
}*/

func (s *IndexBlockService) ProcessFiltersBlock(filters *dia.FiltersBlock) {
	s.chanFiltersBlock <- filters
}

func (s *IndexBlockService) ProcessSuppliesBlock(supplies *dia.SuppliesBlock) {
	s.chanSuppliesBlock <- supplies
}

func (s *IndexBlockService) Close() error {
	if s.closed {
		return errors.New("IndexBlockService: Already closed")
	}
	close(s.shutdown)
	<-s.shutdownDone
	return s.error
}

// must only be called from mainLoop
func (s *IndexBlockService) cleanup(err error) {

	s.errorLock.Lock()
	defer s.errorLock.Unlock()

	if err != nil {
		s.error = err
	}
	s.closed = true

	close(s.shutdownDone) // signal that shutdown is complete
}

func calculateIndexValue(composition dia.IndexBlock) float64 {
	/*
		r := calculateIndexValueWithoutMagicNumber(composition)
		m := r / 1000
		log.Info("calculateIndexValue: index = ", r/m)
		fmt.Printf("magic number should be %f\n", m)
		return 1000.0
	*/
	return calculateIndexValueWithoutMagicNumber(composition) / magicNumber
}

func calculateIndexValueWithoutMagicNumber(composition dia.IndexBlock) float64 {
	result := 0.0
	for _, v := range composition.IndexBlockData.IndexElements {
		result += v.FilteredPoint.Value * v.Supply.CirculatingSupply * v.Percentage
	}
	return result
}

func valueTokenette(indexValue float64) float64 {
	//10000 tokenettes = 1 token,
	//1 token = indexValue USD
	return valueToken(indexValue) / numberTokenettesInOneToken
}

func valueToken(indexValue float64) float64 {
	return usdPerPointsOfIndex * indexValue
}

type compositeNode struct {
	Symbol        string
	CoinMarketCap float64
}

func calculateFirstIndexElement(sortedCompositeNodes []compositeNode, maxCapPercentage float64) dia.IndexElement {
	marketCap := 0.0
	for _, v := range sortedCompositeNodes {
		marketCap += v.CoinMarketCap
	}
	node := sortedCompositeNodes[0]

	percentage := 0.0

	if marketCap != 0.0 {
		percentage = node.CoinMarketCap * 100 / marketCap
	} else {
		log.Error("error MarketCap = 0 for ", node.Symbol)
	}

	result := dia.IndexElement{
		Name:       helpers.NameForSymbol(node.Symbol),
		Symbol:     node.Symbol,
		Percentage: percentage,
	}

	if result.Percentage > maxCapPercentage {
		result.Percentage = maxCapPercentage
	}
	return result
}

func maxCap(cap float64, index int) float64 {
	result := cap / math.Pow((1-cap/100), float64(index))
	return result
}

func CalculateComposite(time time.Time, points map[string]dia.FilterPoint, supplies map[string]dia.Supply, numberElementsInComposite int, cap float64) dia.IndexBlock {
	var result dia.IndexBlock
	result.IndexBlockData.Time = time

	var ss []compositeNode
	for _, v := range points {

		supply := supplies[v.Symbol]
		ss = append(ss, compositeNode{v.Symbol, v.Value * supply.CirculatingSupply})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].CoinMarketCap > ss[j].CoinMarketCap
	})

	if len(ss) > numberElementsInComposite {
		ss = ss[:numberElementsInComposite] //ignore coins out of index
	}

	remainingPercentage := 1.0
	remainingPercentageLast := 100.0

	for i, _ := range ss {
		nodes := ss[i:]
		indexNode := calculateFirstIndexElement(nodes, maxCap(cap, i))
		indexNode.FilteredPoint = points[indexNode.Symbol]

		indexNode.Supply = supplies[indexNode.Symbol]
		remainingPercentage = remainingPercentage - indexNode.Percentage/100
		indexNode.Percentage = indexNode.Percentage * remainingPercentageLast / 100
		remainingPercentageLast = remainingPercentageLast - indexNode.Percentage
		result.IndexBlockData.IndexElements = append(result.IndexBlockData.IndexElements, indexNode)
	}

	result.IndexBlockData.IndexValue = calculateIndexValue(result)
	result.IndexBlockData.ValueTokenette = valueTokenette(result.IndexBlockData.IndexValue)
	result.IndexBlockData.ValueToken = valueToken(result.IndexBlockData.IndexValue)
	result.IndexBlockData.USDPerPointsOfIndex = usdPerPointsOfIndex
	result.IndexBlockData.IndexElementsNumber = len(result.IndexBlockData.IndexElements)
	return result
}

// runs in a goroutine until s is closed
func (s *IndexBlockService) mainLoop() {
	for {
		select {
		case <-s.shutdown:
			log.Println("IndexBlockService shutting down")
			s.cleanup(nil)
			return

		case sb, ok := <-s.chanSuppliesBlock:
			if ok == false {
				log.Println("Error reading channel")
				break
			}

			s.processSupplyBlock(sb)

		case fb, ok := <-s.chanFiltersBlock:
			if ok == false {
				log.Println("Error reading channel")
				break
			}

			s.points = make(map[string]dia.FilterPoint)

			for _, filter := range fb.FiltersBlockData.FilterPoints {
				if filter.Name == s.filterNameUsed {
					s.points[filter.Symbol] = filter
				}
			}
			if s.supplies == nil {
				log.Println("cant calculate new index cause no supplies block received yet.", fb.BlockHash, fb)
			} else {


				result := CalculateComposite(fb.FiltersBlockData.EndTime, s.points, s.supplies, 15, 30.0)
				result.IndexBlockData.FiltersBlockHash = fb.BlockHash
				result.IndexBlockData.SuppliesBlockHash = s.suppliesBlock.BlockHash
				/*if s.volatilityBlock != nil {
					result.IndexBlockData.VolatilityBlockHash = s.volatilityBlock.BlockHash
				}*/
				hash, err := structhash.Hash(result.IndexBlockData, 1)
				if err != nil {
					log.Printf("error on hash")
					hash = "hashError"
				}
				result.BlockHash = hash

				if len(result.IndexBlockData.IndexElements) != 0 {
					if s.shouldGenerateNewBlock() {
						s.chanIndexBlock <- &result
						s.chanIndexBlockDaily <- &result
					} else {
						/*if s.shouldGenerateNewDailyBlock() {
							s.chanIndexBlockDaily <- &result
						}*/
						s.chanIndexBlock2 <- &result
					}
				} else {
					log.Println("ignoring result IndexBlockData because empty", result)
				}
			}
		}
	}
}
