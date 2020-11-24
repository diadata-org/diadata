package supplyBlockService

import (
	"fmt"
	"github.com/cnf/structhash"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	refreshDelay = time.Second * 60
)

type nothing struct {
}

type SupplyScraper struct {
	// signaling channels
	shutdown          chan nothing
	shutdownDone      chan nothing
	ticker            *time.Ticker
	closed            bool
	currentSupplies   map[string]dia.Supply
	chanSuppliesBlock chan *dia.SuppliesBlock
	datastore         models.Datastore
}

// NewSupplyScraper returns a new SupplyScraper initialized with default values.
// The instance is asynchronously scraping as soon as it is created.
func NewSupplyScraper() *SupplyScraper {
	s := &SupplyScraper{
		shutdown:          make(chan nothing),
		shutdownDone:      make(chan nothing),
		currentSupplies:   map[string]dia.Supply{},
		chanSuppliesBlock: make(chan *dia.SuppliesBlock),
		ticker:            time.NewTicker(refreshDelay),
	}

	d, err := models.NewDataStore()
	if err == nil {
		s.datastore = d
	} else {
		log.Fatal("NewDataStore failure")
	}
	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (s *SupplyScraper) mainLoop() {
	s.update()
	for {
		select {
		case <-s.ticker.C:
			s.update()
		case <-s.shutdown: // user requested shutdown
			log.Println("SupplyScraper shutting down")
			s.cleanup(nil)
			return
		}
	}
}

// must only be called from mainLoop
func (s *SupplyScraper) cleanup(err error) {
	s.ticker.Stop()
	s.closed = true
	close(s.shutdownDone) // signal that shutdown is complete
}

func (ps *SupplyScraper) Channel() chan *dia.SuppliesBlock {
	return ps.chanSuppliesBlock
}

func (s *SupplyScraper) finaliseBlock() {
	log.Printf("finaliseBlock")

	currentBlock := dia.SuppliesBlock{
		BlockData: dia.SuppliesBlockData{
			Supplies: []dia.Supply{},
			Time:     time.Now(),
		},
	}

	for _, v := range s.currentSupplies {
		currentBlock.BlockData.Supplies = append(currentBlock.BlockData.Supplies, v)
	}

	hash, err := structhash.Hash(currentBlock.BlockData, 1)
	if err != nil {
		log.Printf("error on hash")
		hash = "hashError"
	}

	currentBlock.BlockHash = hash

	s.chanSuppliesBlock <- &currentBlock

	log.Println(currentBlock)

}

func getSupplyFromDia(symbol string) (*dia.Supply, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/supply/" + symbol)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api")
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%s\n", string(contents))

		var b dia.Supply
		err = b.UnmarshalBinary(contents)
		if err == nil {
			log.Info(b)
			return &b, nil
		}
		return nil, err
	}
}

func (s *SupplyScraper) update() error {
	l, err := dia.GetSymbolsList("http://api.blockstate.com/")
	if err == nil {
		for _, symbol := range l {
			supply, errDia := getSupplyFromDia(symbol)
			if errDia == nil && supply != nil {
				s.currentSupplies[symbol] = *supply
				s.datastore.SetSupply(supply)
			}
		}
	}
	s.finaliseBlock()
	return nil
}
