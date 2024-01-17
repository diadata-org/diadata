package source

import (
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/horizonhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/stellar/go/clients/horizonclient"
)

const (
	defaultAssetRequestCursor = ""
	defaultAssetRequestLimit  = 200
)

type StellarAssetSource struct {
	horizonClient *horizonclient.Client
	assetChannel  chan dia.Asset
	doneChannel   chan bool
	blockchain    string
	relDB         *models.RelDB
	cursor        *string
	limit         uint
}

func NewStellarAssetSource(exchange dia.Exchange) *StellarAssetSource {
	var (
		horizonClient *horizonclient.Client
		assetChannel  = make(chan dia.Asset)
		doneChannel   = make(chan bool)
	)

	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_ASSETS_CURSOR", defaultAssetRequestCursor)
	limit := utils.GetenvUint(strings.ToUpper(exchange.Name)+"_ASSETS_LIMIT", defaultAssetRequestLimit)

	switch exchange.Name {
	case dia.StellarExchange:
		horizonClient = horizonclient.DefaultPublicNetClient
	}

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("make new relational datastore: ", err)
	}

	scraper := &StellarAssetSource{
		horizonClient: horizonClient,
		assetChannel:  assetChannel,
		doneChannel:   doneChannel,
		blockchain:    exchange.BlockChain.Name,
		relDB:         relDB,
		cursor:        &cursor,
		limit:         limit,
	}

	go func() {
		scraper.fetchAssets()
	}()
	return scraper
}

func (scraper *StellarAssetSource) fetchAssets() {
	page, err := scraper.horizonClient.Assets(horizonclient.AssetRequest{
		Cursor: *scraper.cursor,
		Limit:  scraper.limit,
	})
	if err != nil {
		log.Error(err)
		return
	}

	recordsFound := len(page.Embedded.Records) > 0
	for recordsFound {
		for _, stellarAsset := range page.Embedded.Records {
			log.Infof("asset: %s", stellarAsset.PT)

			asset, err := getDIAAsset(scraper, stellarAsset.Code, stellarAsset.Issuer)
			if err != nil {
				log.Error(err)
				continue
			}

			scraper.assetChannel <- asset
		}

		nextPage, err := scraper.horizonClient.NextAssetsPage(page)
		if err != nil {
			log.Error(err)
			return
		}
		page = nextPage

		recordsFound = len(page.Embedded.Records) > 0
		log.Infof("found %d assets", len(page.Embedded.Records))
	}
	scraper.doneChannel <- true
}

func getDIAAsset(scraper *StellarAssetSource, assetCode string, assetIssuer string) (asset dia.Asset, err error) {
	logContext := logrus.Fields{"context": "StellarTomlReader"}
	assetInfoReader := &horizonhelper.StellarAssetInfo{
		Logger: log.WithFields(logContext),
	}
	asset, err = assetInfoReader.GetStellarAssetInfo(scraper.horizonClient, assetCode, assetIssuer, scraper.blockchain)
	if err != nil {
		log.Error("cannot fetch asset info: ", err)
	}
	return
}

func (scraper *StellarAssetSource) Asset() chan dia.Asset {
	return scraper.assetChannel
}

func (scraper *StellarAssetSource) Done() chan bool {
	return scraper.doneChannel
}
