package nftdatascrapers

// Please implement the scraping of coingecko quotations here.

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	cryptopunk "github.com/diadata-org/diadata/config/nftContracts/cryptopunk"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	structs "github.com/fatih/structs"
)

const (
	cryptoPunksFirstBlock = uint64(3918000)
)

type OpenSeaCryptopunkResponse struct {
	ID                   int         `json:"id"`
	TokenID              string      `json:"token_id"`
	NumSales             int         `json:"num_sales"`
	BackgroundColor      interface{} `json:"background_color"`
	ImageURL             string      `json:"image_url"`
	ImagePreviewURL      string      `json:"image_preview_url"`
	ImageThumbnailURL    string      `json:"image_thumbnail_url"`
	ImageOriginalURL     string      `json:"image_original_url"`
	AnimationURL         interface{} `json:"animation_url"`
	AnimationOriginalURL interface{} `json:"animation_original_url"`
	Name                 string      `json:"name"`
	Description          interface{} `json:"description"`
	ExternalLink         string      `json:"external_link"`
	AssetContract        struct {
		Address                     string      `json:"address"`
		AssetContractType           string      `json:"asset_contract_type"`
		CreatedDate                 string      `json:"created_date"`
		Name                        string      `json:"name"`
		NftVersion                  string      `json:"nft_version"`
		OpenseaVersion              interface{} `json:"opensea_version"`
		Owner                       interface{} `json:"owner"`
		SchemaName                  string      `json:"schema_name"`
		Symbol                      string      `json:"symbol"`
		TotalSupply                 interface{} `json:"total_supply"`
		Description                 string      `json:"description"`
		ExternalLink                string      `json:"external_link"`
		ImageURL                    string      `json:"image_url"`
		DefaultToFiat               bool        `json:"default_to_fiat"`
		DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
		BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
		SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
		PayoutAddress               interface{} `json:"payout_address"`
	} `json:"asset_contract"`
	Owner struct {
		User          interface{} `json:"user"`
		ProfileImgURL string      `json:"profile_img_url"`
		Address       string      `json:"address"`
		Config        string      `json:"config"`
		DiscordID     string      `json:"discord_id"`
	} `json:"owner"`
	Permalink  string `json:"permalink"`
	Collection struct {
		PaymentTokens []struct {
			ID       int     `json:"id"`
			Symbol   string  `json:"symbol"`
			Address  string  `json:"address"`
			ImageURL string  `json:"image_url"`
			Name     string  `json:"name"`
			Decimals int     `json:"decimals"`
			EthPrice float64 `json:"eth_price"`
			UsdPrice float64 `json:"usd_price"`
		} `json:"payment_tokens"`
		PrimaryAssetContracts []struct {
			Address                     string      `json:"address"`
			AssetContractType           string      `json:"asset_contract_type"`
			CreatedDate                 string      `json:"created_date"`
			Name                        string      `json:"name"`
			NftVersion                  string      `json:"nft_version"`
			OpenseaVersion              interface{} `json:"opensea_version"`
			Owner                       interface{} `json:"owner"`
			SchemaName                  string      `json:"schema_name"`
			Symbol                      string      `json:"symbol"`
			TotalSupply                 interface{} `json:"total_supply"`
			Description                 string      `json:"description"`
			ExternalLink                string      `json:"external_link"`
			ImageURL                    string      `json:"image_url"`
			DefaultToFiat               bool        `json:"default_to_fiat"`
			DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
			DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
			OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
			OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
			OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
			BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
			SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
			PayoutAddress               interface{} `json:"payout_address"`
		} `json:"primary_asset_contracts"`
		Traits struct {
		} `json:"traits"`
		Stats struct {
			SevenDayVolume       float64 `json:"seven_day_volume"`
			TotalVolume          float64 `json:"total_volume"`
			SevenDayChange       float64 `json:"seven_day_change"`
			SevenDaySales        float64 `json:"seven_day_sales"`
			TotalSales           float64 `json:"total_sales"`
			TotalSupply          float64 `json:"total_supply"`
			Count                float64 `json:"count"`
			NumOwners            int     `json:"num_owners"`
			SevenDayAveragePrice float64 `json:"seven_day_average_price"`
			AveragePrice         float64 `json:"average_price"`
			NumReports           int     `json:"num_reports"`
			MarketCap            float64 `json:"market_cap"`
		} `json:"stats"`
		BannerImageURL          string      `json:"banner_image_url"`
		ChatURL                 interface{} `json:"chat_url"`
		CreatedDate             string      `json:"created_date"`
		DefaultToFiat           bool        `json:"default_to_fiat"`
		Description             string      `json:"description"`
		DevBuyerFeeBasisPoints  string      `json:"dev_buyer_fee_basis_points"`
		DevSellerFeeBasisPoints string      `json:"dev_seller_fee_basis_points"`
		DiscordURL              string      `json:"discord_url"`
		DisplayData             struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalURL                 string      `json:"external_url"`
		Featured                    bool        `json:"featured"`
		FeaturedImageURL            string      `json:"featured_image_url"`
		Hidden                      bool        `json:"hidden"`
		SafelistRequestStatus       string      `json:"safelist_request_status"`
		ImageURL                    string      `json:"image_url"`
		IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
		LargeImageURL               string      `json:"large_image_url"`
		MediumUsername              interface{} `json:"medium_username"`
		Name                        string      `json:"name"`
		OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
		OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
		OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
		PayoutAddress               interface{} `json:"payout_address"`
		RequireEmail                bool        `json:"require_email"`
		ShortDescription            interface{} `json:"short_description"`
		Slug                        string      `json:"slug"`
		TelegramURL                 interface{} `json:"telegram_url"`
		TwitterUsername             string      `json:"twitter_username"`
		InstagramUsername           interface{} `json:"instagram_username"`
		WikiURL                     interface{} `json:"wiki_url"`
	} `json:"collection"`
	Decimals      interface{}        `json:"decimals"`
	TokenMetadata string             `json:"token_metadata"`
	SellOrders    interface{}        `json:"sell_orders"`
	Creator       interface{}        `json:"creator"`
	Traits        []CryptopunkTraits `json:"traits"`
	LastSale      struct {
		Asset struct {
			TokenID  string      `json:"token_id"`
			Decimals interface{} `json:"decimals"`
		} `json:"asset"`
		AssetBundle    interface{} `json:"asset_bundle"`
		EventType      string      `json:"event_type"`
		EventTimestamp string      `json:"event_timestamp"`
		AuctionType    interface{} `json:"auction_type"`
		TotalPrice     string      `json:"total_price"`
		PaymentToken   struct {
			ID       int    `json:"id"`
			Symbol   string `json:"symbol"`
			Address  string `json:"address"`
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
			Decimals int    `json:"decimals"`
			EthPrice string `json:"eth_price"`
			UsdPrice string `json:"usd_price"`
		} `json:"payment_token"`
		Transaction struct {
			BlockHash   string `json:"block_hash"`
			BlockNumber string `json:"block_number"`
			FromAccount struct {
				User          interface{} `json:"user"`
				ProfileImgURL string      `json:"profile_img_url"`
				Address       string      `json:"address"`
				Config        string      `json:"config"`
				DiscordID     string      `json:"discord_id"`
			} `json:"from_account"`
			ID        int    `json:"id"`
			Timestamp string `json:"timestamp"`
			ToAccount struct {
				User          interface{} `json:"user"`
				ProfileImgURL string      `json:"profile_img_url"`
				Address       string      `json:"address"`
				Config        string      `json:"config"`
				DiscordID     string      `json:"discord_id"`
			} `json:"to_account"`
			TransactionHash  string `json:"transaction_hash"`
			TransactionIndex string `json:"transaction_index"`
		} `json:"transaction"`
		CreatedDate string `json:"created_date"`
		Quantity    string `json:"quantity"`
	} `json:"last_sale"`
	TopBid                  interface{}   `json:"top_bid"`
	ListingDate             interface{}   `json:"listing_date"`
	IsPresale               bool          `json:"is_presale"`
	TransferFeePaymentToken interface{}   `json:"transfer_fee_payment_token"`
	TransferFee             interface{}   `json:"transfer_fee"`
	RelatedAssets           []interface{} `json:"related_assets"`
	Orders                  []interface{} `json:"orders"`
	Auctions                []interface{} `json:"auctions"`
	SupportsWyvern          bool          `json:"supports_wyvern"`
	TopOwnerships           []struct {
		Owner struct {
			User          interface{} `json:"user"`
			ProfileImgURL string      `json:"profile_img_url"`
			Address       string      `json:"address"`
			Config        string      `json:"config"`
			DiscordID     string      `json:"discord_id"`
		} `json:"owner"`
		Quantity string `json:"quantity"`
	} `json:"top_ownerships"`
	Ownership              interface{} `json:"ownership"`
	HighestBuyerCommitment interface{} `json:"highest_buyer_commitment"`
}

type CryptopunkTraits struct {
	TraitType   string      `json:"trait_type"`
	Value       string      `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
}

type CryptoPunksScraper struct {
	nftscraper    NFTScraper
	address       common.Address
	apiURLOpensea string
	cryptopunkURL string
	ticker        *time.Ticker
}

type CryptopunkOutput struct {
	Traits []CryptopunkTraits `structs:",flatten"`
}

func NewCryptoPunksScraper(rdb *models.RelDB) *CryptoPunksScraper {
	connection, err := ethhelper.NewETHClient()
	if err != nil {
		log.Error("Error connecting Eth Client")
	}

	nftScraper := NFTScraper{
		shutdown:      make(chan nothing),
		shutdownDone:  make(chan nothing),
		error:         nil,
		ethConnection: connection,
		relDB:         rdb,
		chanData:      make(chan dia.NFT),
	}
	s := &CryptoPunksScraper{
		address:       common.HexToAddress("0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB"),
		apiURLOpensea: "https://api.opensea.io/api/v1/",
		cryptopunkURL: "https://www.larvalabs.com/cryptopunks/details/",
		nftscraper:    nftScraper,
		ticker:        time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoPunksScraper) mainLoop() {
	err := scraper.FetchData()
	if err != nil {
		log.Error("error updating NFT: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchData()
			if err != nil {
				log.Error("error updating NFT: ", err)
			}
		case <-scraper.nftscraper.shutdown: // user requested shutdown
			log.Printf("Cryptopunk scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoPunksScraper) FetchData() (err error) {
	totalSupply, err := scraper.GetTotalSupply()
	if err != nil {
		return
	}

	fmt.Println("total supply: ", int(totalSupply.Int64()))

	nftClassID, err := scraper.nftscraper.relDB.GetNFTClassID(scraper.address.Hex(), dia.Ethereum)
	if err != nil {
		log.Error("getting nftclass ID: ", err)
	}
	cryptopunkNFTClass, err := scraper.nftscraper.relDB.GetNFTClassByID(nftClassID)
	if err != nil {
		log.Error("getting nft by ID: ", err)
	}

	creationTimeMap, creatorMap, err := scraper.GetCreationEvents()
	if err != nil {
		log.Error("getting nft creation data: ", err)
	}

	for i := 0; i < int(totalSupply.Int64()); i++ {
		var out CryptopunkOutput

		out.Traits, err = scraper.GetOpenSeaPunk(big.NewInt(int64(i)))
		if err != nil {
			log.Errorf("getting Opensea data: %+v", err)
		}
		// 3. combine both in order to fill dia.NFT
		result := structs.Map(out)
		nft := dia.NFT{
			TokenID:        strconv.Itoa(i),
			NFTClass:       cryptopunkNFTClass,
			CreationTime:   creationTimeMap[uint64(i)],
			CreatorAddress: creatorMap[uint64(i)].Hex(),
			Attributes:     result,
			URI:            scraper.cryptopunkURL + strconv.Itoa(i),
		}
		scraper.GetDataChannel() <- nft
	}
	return nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *CryptoPunksScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := cryptopunk.NewCryptoPunksMarketCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// TokenByIndex returns the address of a punk whose id is passed as a parameter from on-chain.
func (scraper *CryptoPunksScraper) TokenByIndex(index *big.Int) (common.Address, error) {
	contract, err := cryptopunk.NewCryptoPunksMarketCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.PunkIndexToAddress(&bind.CallOpts{}, index)
}

// GetOpenSeaPunk returns the scraped data from Opensea for a given punk
func (scraper *CryptoPunksScraper) GetOpenSeaPunkOld(index *big.Int) ([]CryptopunkTraits, error) {
	var traits []CryptopunkTraits
	url := scraper.apiURLOpensea + "asset/" + scraper.address.String() + "/" + index.String()

	respData, statusCode, err := utils.GetRequestWithStatus(url)
	log.Info("statusCode, err: ", statusCode, err)

	count := 0
	for statusCode != 200 && count < 40 {
		// Retry get request
		log.Infof("sleep for %v seconds", count)
		time.Sleep(time.Millisecond * time.Duration(openseaAPIWait*count))
		respData, statusCode, err = utils.GetRequestWithStatus(url)
		log.Info("statusCode, err in retry: ", statusCode, err)
		count++
	}

	traits, err = GetCryptopunkTraits(respData)
	if err != nil {
		return nil, err
	}

	return traits, nil

}

func (scraper *CryptoPunksScraper) GetOpenSeaPunk(index *big.Int) ([]CryptopunkTraits, error) {
	var traits []CryptopunkTraits
	openseaURL := scraper.apiURLOpensea + "asset/" + scraper.address.String() + "/" + index.String()

	respData, statusCode, err := utils.OpenseaGetRequest(openseaURL)
	if err != nil {
		return traits, err
	}

	count := 0
	for statusCode != 200 && count < 40 {
		// Retry get request
		log.Infof("sleep for %v seconds", count)
		time.Sleep(time.Millisecond * time.Duration(openseaAPIWait*count))
		respData, statusCode, err = utils.OpenseaGetRequest(openseaURL)
		log.Info("statusCode, err in retry: ", statusCode, err)
		count++
	}

	traits, err = GetCryptopunkTraits(respData)
	if err != nil {
		return nil, err
	}
	return traits, nil
}

// GetCreationEvents returns maps for creation time and creator address by filtering 'assign punk' events.
func (scraper *CryptoPunksScraper) GetCreationEvents() (map[uint64]time.Time, map[uint64]common.Address, error) {
	log.Info("fetching creation events ...")
	creationTimeMap := make(map[uint64]time.Time)
	creatorAddressMap := make(map[uint64]common.Address)
	filterer, err := cryptopunk.NewCryptoPunksMarketFilterer(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		return creationTimeMap, creatorAddressMap, err
	}

	header, err := scraper.nftscraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return creationTimeMap, creatorAddressMap, err
	}

	endBlockNumber := header.Number.Uint64() - blockDelayEthereum
	startBlockNumber := cryptoPunksFirstBlock

	for endBlockNumber <= header.Number.Uint64()-blockDelayEthereum {
		var iter *cryptopunk.CryptoPunksMarketAssignIterator
		fmt.Printf("startblock -- endblock: %v -- %v ...\n", startBlockNumber, endBlockNumber)
		iter, err = filterer.FilterAssign(&bind.FilterOpts{
			Start: startBlockNumber,
			End:   &endBlockNumber,
		}, nil)
		if err != nil {
			if strings.Contains(err.Error(), "query returned more than 10000 results") || strings.Contains(err.Error(), "Log response size exceeded") {
				log.Info("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = startBlockNumber + (endBlockNumber-startBlockNumber)/2
				continue
			}
			log.Error("filtering assign punk: ", err)
			return creationTimeMap, creatorAddressMap, err
		}

		// map punk index to timestamp of creation event and to creator address.
		var blockData dia.BlockData
		for iter.Next() {

			blockData, err = ethhelper.GetBlockData(int64(iter.Event.Raw.BlockNumber), scraper.nftscraper.relDB, scraper.nftscraper.ethConnection)
			if err != nil {
				log.Errorf("getting blockdata: %+v", err)
			}
			switch blockData.Data["Time"].(type) {
			case float64:
				creationTimeMap[iter.Event.PunkIndex.Uint64()] = time.Unix(int64(blockData.Data["Time"].(float64)), 0)
			case uint64:
				creationTimeMap[iter.Event.PunkIndex.Uint64()] = time.Unix(int64(blockData.Data["Time"].(uint64)), 0)
			}
			creatorAddressMap[iter.Event.PunkIndex.Uint64()] = iter.Event.To
		}
		startBlockNumber = endBlockNumber
		endBlockNumber = endBlockNumber + batchSize
	}
	return creationTimeMap, creatorAddressMap, err
}

// GetCryptopunkTraits returns the parsed traits data from Opensea for a given punk
func GetCryptopunkTraits(punkResp []byte) ([]CryptopunkTraits, error) {
	var resp OpenSeaCryptopunkResponse
	if err := json.Unmarshal(punkResp, &resp); err != nil {
		return nil, err
	}
	return resp.Traits, nil
}

// GetCryptopunkAddress returns the creator address from Opensea
func GetCryptopunkAddress(punkResp []byte) (common.Address, error) {

	return common.HexToAddress("0xc352b534e8b987e036a93539fd6897f53488e56a"), nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptoPunksScraper) GetDataChannel() chan dia.NFT {
	return scraper.nftscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoPunksScraper) cleanup(err error) {
	scraper.nftscraper.errorLock.Lock()
	defer scraper.nftscraper.errorLock.Unlock()
	scraper.ticker.Stop()
	if err != nil {
		scraper.nftscraper.error = err
	}
	scraper.nftscraper.closed = true
	close(scraper.nftscraper.shutdownDone) // signal that shutdown is complete
}

// Close closes any existing API connections
func (scraper *CryptoPunksScraper) Close() error {
	if scraper.nftscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.nftscraper.shutdown)
	<-scraper.nftscraper.shutdownDone
	scraper.nftscraper.errorLock.RLock()
	defer scraper.nftscraper.errorLock.RUnlock()
	return scraper.nftscraper.error
}
