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

	cryptokitties "github.com/diadata-org/diadata/config/nftContracts/cryptokitties"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	structs "github.com/fatih/structs"
)

const (
	openseaAPIWait          = 1000
	cryptoKittiesFirstBlock = uint64(4605167)
	batchSize               = 5000
)

type Kitty struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}

type CryptokittiesTraits struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
}

type CryptokittiesOutput struct {
	Traits []CryptokittiesTraits `structs:",flatten"`
	Kitty  Kitty                 `structs:",flatten"`
}

type OpenSeaCryptokittiesResponse struct {
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
			OneDayVolume          float64 `json:"one_day_volume"`
			OneDayChange          float64 `json:"one_day_change"`
			OneDaySales           float64 `json:"one_day_sales"`
			OneDayAveragePrice    float64 `json:"one_day_average_price"`
			SevenDayVolume        float64 `json:"seven_day_volume"`
			SevenDayChange        float64 `json:"seven_day_change"`
			SevenDaySales         float64 `json:"seven_day_sales"`
			SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
			ThirtyDayVolume       float64 `json:"thirty_day_volume"`
			ThirtyDayChange       float64 `json:"thirty_day_change"`
			ThirtyDaySales        float64 `json:"thirty_day_sales"`
			ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
			TotalVolume           float64 `json:"total_volume"`
			TotalSales            float64 `json:"total_sales"`
			TotalSupply           float64 `json:"total_supply"`
			Count                 float64 `json:"count"`
			NumOwners             int     `json:"num_owners"`
			AveragePrice          float64 `json:"average_price"`
			NumReports            int     `json:"num_reports"`
			MarketCap             float64 `json:"market_cap"`
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
	Decimals      interface{} `json:"decimals"`
	TokenMetadata string      `json:"token_metadata"`
	SellOrders    interface{} `json:"sell_orders"`
	Creator       struct {
		User          interface{}
		ProfileImgUrl string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
		DiscordID     string `json:"discord_id"`
	} `json:"creator"`
	Traits   []CryptokittiesTraits `json:"traits"`
	LastSale struct {
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

type CryptoKittiesScraper struct {
	nftscraper       NFTScraper
	address          common.Address
	apiURLOpensea    string
	cryptokittiesURL string
	ticker           *time.Ticker
}

func NewCryptoKittiesScraper(rdb *models.RelDB) *CryptoKittiesScraper {
	connection, err := ethclient.Dial("https://eth-mainnet.alchemyapi.io/v2/v1bo6tRKiraJ71BVGKmCtWVedAzzNTd6")
	// connection, err := ethhelper.NewETHClient()
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
	s := &CryptoKittiesScraper{
		address:          common.HexToAddress("0x06012c8cf97BEaD5deAe237070F9587f8E7A266d"),
		apiURLOpensea:    "https://api.opensea.io/api/v1/",
		cryptokittiesURL: "https://www.cryptokitties.co/",
		nftscraper:       nftScraper,
		ticker:           time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptoKittiesScraper) mainLoop() {
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
			log.Printf("Cryptokitties scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *CryptoKittiesScraper) FetchData() error {
	totalSupply, err := scraper.GetTotalSupply()
	if err != nil {
		return err
	}

	fmt.Println("total supply: ", int(totalSupply.Int64()))

	cryptokittiesNFTClass, err := scraper.nftscraper.relDB.GetNFTClass(scraper.address.Hex(), dia.Ethereum)
	if err != nil {
		log.Error("getting nftclass: ", err)
	}

	for i := 0; i < int(totalSupply.Int64()); i++ {
		var out CryptokittiesOutput
		var creatorAddress common.Address

		// 1. fetch data from onchain
		out.Kitty, err = scraper.GetKitty(big.NewInt(int64(i)))
		if err != nil {
			log.Errorf("getting kitty data: %+v", err)
			return err
		}

		creationMap, err := scraper.GetCryptokittiesCreationTime()
		if err != nil {
			log.Errorf("getting creation map: %+v", err)
			return err
		}

		// 2. fetch data from offchain
		out.Traits, creatorAddress, err = scraper.GetOpenSeaKitty(big.NewInt(int64(i)))
		if err != nil {
			log.Errorf("getting Opensea data: %+v", err)
		}
		// 3. combine both in order to fill dia.NFT
		result := structs.Map(out)
		// Set output object
		cryptoKittiesNFT := dia.NFT{
			TokenID:        strconv.Itoa(i),
			NFTClass:       cryptokittiesNFTClass,
			CreationTime:   creationMap[uint64(i)],
			CreatorAddress: creatorAddress.Hex(),
			Attributes:     result,
			URI:            scraper.cryptokittiesURL + strconv.Itoa(i),
		}
		fmt.Println("get nft: ", cryptoKittiesNFT)
		scraper.GetDataChannel() <- cryptoKittiesNFT
	}
	return nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *CryptoKittiesScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := cryptokitties.NewKittyAuctionCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// GetKitty returns the kitty attributes of the NFT from on-chain.
func (scraper *CryptoKittiesScraper) GetKitty(kittyId *big.Int) (Kitty, error) {
	contract, err := cryptokitties.NewKittyCoreCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.GetKitty(&bind.CallOpts{}, kittyId)
}

// GetOpenSeaKitty returns the scraped data from Opensea for a given kitty
func (scraper *CryptoKittiesScraper) GetOpenSeaKitty(index *big.Int) ([]CryptokittiesTraits, common.Address, error) {
	var traits []CryptokittiesTraits
	var creatorAddress common.Address
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

	traits, err = GetCryptokittiesTraits(respData)
	if err != nil {
		return traits, creatorAddress, err
	}

	creatorAddress, err = GetCryptokittiesAddress(respData)
	if err != nil {
		return traits, creatorAddress, err
	}

	return traits, creatorAddress, nil

}

// GetCryptokittiesTraits returns the traits data from Opensea for a given kitty
func GetCryptokittiesTraits(kittyResp []byte) ([]CryptokittiesTraits, error) {
	var resp OpenSeaCryptokittiesResponse
	if err := json.Unmarshal(kittyResp, &resp); err != nil {
		return nil, err
	}
	return resp.Traits, nil
}

// GetCryptokittiesAddress returns the creator address from Opensea
func GetCryptokittiesAddress(kittyResp []byte) (common.Address, error) {
	var resp OpenSeaCryptokittiesResponse
	if err := json.Unmarshal(kittyResp, &resp); err != nil {
		return common.Address{}, err
	}
	return common.HexToAddress(resp.Creator.Address), nil
}

// GetCryptokittiesCreationTime returns a map[uint64]uint64 mapping a
func (scraper *CryptoKittiesScraper) GetCryptokittiesCreationTime() (map[uint64]time.Time, error) {
	creationMap := make(map[uint64]time.Time)
	filterer, err := cryptokitties.NewKittyBaseFilterer(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		return creationMap, err
	}

	header, err := scraper.nftscraper.ethConnection.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return creationMap, err
	}
	endBlockNumber := header.Number.Uint64() - blockDelayEthereum
	startBlockNumber := cryptoKittiesFirstBlock

	for endBlockNumber <= header.Number.Uint64()-blockDelayEthereum {
		var iter *cryptokitties.KittyBaseBirthIterator
		iter, err = filterer.FilterBirth(&bind.FilterOpts{
			Start: startBlockNumber,
			End:   &endBlockNumber,
		})
		if err != nil {
			if strings.Contains(err.Error(), "query returned more than 10000 results") || strings.Contains(err.Error(), "Log response size exceeded") {
				log.Info("Got `query returned more than 10000 results` error, reduce the window size and try again...")
				endBlockNumber = startBlockNumber + (endBlockNumber-startBlockNumber)/2
				log.Info("startblock - endblock: ", startBlockNumber, endBlockNumber)
				continue
			}
			log.Error("filtering kitty birth: ", err)
			return creationMap, err
		}

		// map kitty id to timestamp of creation event.
		var blockData dia.BlockData
		for iter.Next() {
			blockData, err = ethhelper.GetBlockData(int64(iter.Event.Raw.BlockNumber), scraper.nftscraper.relDB, scraper.nftscraper.ethConnection)
			if err != nil {
				log.Errorf("getting blockdata: %+v", err)
			}
			switch blockData.Data["Time"].(type) {
			case float64:
				creationMap[iter.Event.KittyId.Uint64()] = time.Unix(int64(blockData.Data["Time"].(float64)), 0)
			case uint64:
				creationMap[iter.Event.KittyId.Uint64()] = time.Unix(int64(blockData.Data["Time"].(uint64)), 0)
			}
		}
		startBlockNumber = endBlockNumber
		endBlockNumber = endBlockNumber + batchSize
	}
	return creationMap, err
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptoKittiesScraper) GetDataChannel() chan dia.NFT {
	return scraper.nftscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptoKittiesScraper) cleanup(err error) {
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
func (scraper *CryptoKittiesScraper) Close() error {
	if scraper.nftscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.nftscraper.shutdown)
	<-scraper.nftscraper.shutdownDone
	scraper.nftscraper.errorLock.RLock()
	defer scraper.nftscraper.errorLock.RUnlock()
	return scraper.nftscraper.error
}
