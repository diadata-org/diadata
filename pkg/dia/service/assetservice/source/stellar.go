package source

import (
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/horizonhelper"
	"github.com/diadata-org/diadata/pkg/utils"
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
	cursor        *string
	limit         uint
}

func NewStellarAssetSource(exchange dia.Exchange) *StellarAssetSource {
	cursor := utils.Getenv(strings.ToUpper(exchange.Name)+"_ASSETS_CURSOR", defaultAssetRequestCursor)
	limit := utils.GetenvUint(strings.ToUpper(exchange.Name)+"_ASSETS_LIMIT", defaultAssetRequestLimit)
	chain := utils.Getenv(strings.ToUpper(exchange.Name)+"_CHAIN", horizonhelper.PublicChain)

	var (
		assetChannel = make(chan dia.Asset)
		doneChannel  = make(chan bool)
	)

	horizonClient := horizonhelper.HorizonClient(chain)

	scraper := &StellarAssetSource{
		horizonClient: horizonClient,
		assetChannel:  assetChannel,
		doneChannel:   doneChannel,
		blockchain:    exchange.BlockChain.Name,
		cursor:        &cursor,
		limit:         limit,
	}

	go func() {
		log.WithField("context", exchange.BlockChain.Name).
			WithField("horizonUrl", horizonClient.HorizonURL).
			Info("Started")
		scraper.fetchAssets()
	}()
	return scraper
}

func (scraper *StellarAssetSource) fetchAssets() {
	scraper.assetChannel <- horizonhelper.GetStellarNativeAssetInfo(scraper.blockchain)

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

			asset, err := scraper.getDIAAsset(stellarAsset.Code, stellarAsset.Issuer)
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

func (scraper *StellarAssetSource) getDIAAsset(assetCode string, assetIssuer string) (asset dia.Asset, err error) {
	asset, err = horizonhelper.GetStellarAssetInfo(scraper.horizonClient, assetCode, assetIssuer, scraper.blockchain)
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
