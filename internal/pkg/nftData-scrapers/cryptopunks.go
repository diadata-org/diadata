package nftdatascrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
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

type CryptopunkScraper struct {
	nftscraper    NFTScraper
	address       common.Address
	apiURLOpensea string
	cryptopunkURL string
	ticker        *time.Ticker
}

type CryptopunkOutput struct {
	Traits []CryptopunkTraits `structs:",flatten"`
}

func NewCryptopunkScraper(rdb *models.RelDB) *CryptopunkScraper {
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
	s := &CryptopunkScraper{
		address:       common.HexToAddress("0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb"),
		apiURLOpensea: "https://api.opensea.io/api/v1/",
		cryptopunkURL: "https://www.larvalabs.com/cryptopunks/details/",
		nftscraper:    nftScraper,
		ticker:        time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *CryptopunkScraper) mainLoop() {
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.UpdateNFT()
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

func (scraper *CryptopunkScraper) UpdateNFT() error {
	fmt.Println("fetch data...")
	nfts, err := scraper.FetchData()
	if err != nil {
		return err
	}
	for _, nft := range nfts {
		scraper.GetDataChannel() <- nft
	}
	return nil
}

func (scraper *CryptopunkScraper) FetchData() (nfts []dia.NFT, err error) {
	// TO DO: iterate over all NFT IDs and fill []dia.NFT
	totalSupply, err := scraper.GetTotalSupply()
	if err != nil {
		return
	}

	fmt.Println("total supply: ", int(totalSupply.Int64()))

	var cryptopunkNFTs []dia.NFT
	nftClassID, err := scraper.nftscraper.relDB.GetNFTClassID(scraper.address, dia.Ethereum)
	if err != nil {
		log.Error("getting nftclass ID: ", err)
	}
	cryptopunkNFTClass, err := scraper.nftscraper.relDB.GetNFTClassByID(nftClassID)
	if err != nil {
		log.Error("getting nft by ID: ", err)
	}

	for i := 0; i < int(totalSupply.Int64()); i++ {
		var out CryptopunkOutput
		var creatorAddress common.Address
		var creationTime time.Time

		out.Traits, creatorAddress, creationTime, err = scraper.GetOpenSeaPunk(big.NewInt(int64(i)))
		if err != nil {
			log.Errorf("Error getting Opensea data: %+v", err)
		}
		// 3. combine both in order to fill dia.NFT
		result := structs.Map(out)
		cryptopunkNFTs = append(cryptopunkNFTs, dia.NFT{
			TokenID:        strconv.Itoa(i),
			NFTClass:       cryptopunkNFTClass,
			CreationTime:   creationTime,
			CreatorAddress: creatorAddress,
			Attributes:     result,
			URI:            scraper.cryptopunkURL + strconv.Itoa(i),
		})
	}
	return cryptopunkNFTs, nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *CryptopunkScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := cryptopunk.NewCryptoPunksMarketCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// TokenByIndex returns the address of a punk whose id is passed as a parameter from on-chain.
func (scraper *CryptopunkScraper) TokenByIndex(index *big.Int) (common.Address, error) {
	contract, err := cryptopunk.NewCryptoPunksMarketCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.PunkIndexToAddress(&bind.CallOpts{}, index)
}

// GetOpenSeaPunk returns the scraped data from Opensea for a given punk
func (scraper *CryptopunkScraper) GetOpenSeaPunk(index *big.Int) ([]CryptopunkTraits, common.Address, time.Time, error) {
	var creatorAddress common.Address
	var creationTime time.Time
	url := scraper.apiURLOpensea + "asset/" + scraper.address.String() + "/" + index.String()
	respData, err := utils.GetRequest(url)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	traits, err := GetCryptopunkTraits(respData)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	creatorAddress, err = GetCryptopunkAddress(respData)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	creationTime, err = GetCryptopunkCreationTime(respData)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	return traits, creatorAddress, creationTime, nil

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

// GetCryptopunkCreationTime returns the creation time from Opensea
func GetCryptopunkCreationTime(punkResp []byte) (time.Time, error) {
	var resp OpenSeaCryptopunkResponse
	var t time.Time
	if err := json.Unmarshal(punkResp, &resp); err != nil {
		return t, err
	}

	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, resp.AssetContract.CreatedDate)

	if err != nil {
		return t, err
	}

	return t, nil
}

// GetDataChannel returns the scrapers data channel.
func (scraper *CryptopunkScraper) GetDataChannel() chan dia.NFT {
	return scraper.nftscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *CryptopunkScraper) cleanup(err error) {
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
func (scraper *CryptopunkScraper) Close() error {
	if scraper.nftscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.nftscraper.shutdown)
	<-scraper.nftscraper.shutdownDone
	scraper.nftscraper.errorLock.RLock()
	defer scraper.nftscraper.errorLock.RUnlock()
	return scraper.nftscraper.error
}
