package nftdatascrapers

// Please implement the scraping of coingecko quotations here.

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/diadata-org/diadata/config/nftContracts/sorare"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	structs "github.com/fatih/structs"
)

type SorareCard struct {
	PlayerId     *big.Int
	Season       uint16
	Scarcity     *big.Int
	SerialNumber uint16
	Metadata     []byte
	ClubId       uint16
}

type SorarePlayer struct {
	Name         string
	YearOfBirth  uint16
	MonthOfBirth uint8
	DayOfBirth   uint8
}

type SorareClub struct {
	Name        string
	Country     string
	City        string
	YearFounded uint16
}

type Club struct {
	NameClub    string
	CountryClub string
	CityClub    string
	YearFounded uint16
}

type SorareTrait struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType interface{} `json:"display_type"`
	MaxValue    interface{} `json:"max_value"`
	TraitCount  int         `json:"trait_count"`
	Order       interface{} `json:"order"`
}

type SorareOutput struct {
	Card   SorareCard    `structs:",flatten"`
	Player SorarePlayer  `structs:",flatten"`
	Club   Club          `structs:",flatten"`
	Traits []SorareTrait `structs:",flatten"`
}

type OpenSeaResponse struct {
	ID                   int         `json:"id"`
	TokenID              string      `json:"token_id"`
	NumSales             int         `json:"num_sales"`
	BackgroundColor      string      `json:"background_color"`
	ImageURL             string      `json:"image_url"`
	ImagePreviewURL      string      `json:"image_preview_url"`
	ImageThumbnailURL    string      `json:"image_thumbnail_url"`
	ImageOriginalURL     string      `json:"image_original_url"`
	AnimationURL         interface{} `json:"animation_url"`
	AnimationOriginalURL interface{} `json:"animation_original_url"`
	Name                 string      `json:"name"`
	Description          string      `json:"description"`
	ExternalLink         string      `json:"external_link"`
	AssetContract        struct {
		Address                     string      `json:"address"`
		AssetContractType           string      `json:"asset_contract_type"`
		CreatedDate                 string      `json:"created_date"`
		Name                        string      `json:"name"`
		NftVersion                  string      `json:"nft_version"`
		OpenseaVersion              interface{} `json:"opensea_version"`
		Owner                       int         `json:"owner"`
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
		PayoutAddress               string      `json:"payout_address"`
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
			Owner                       int         `json:"owner"`
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
			PayoutAddress               string      `json:"payout_address"`
		} `json:"primary_asset_contracts"`
		Traits struct {
			XP struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"XP"`
			PrintDate struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
			} `json:"print_date"`
			Power struct {
				Min float64 `json:"min"`
				Max float64 `json:"max"`
			} `json:"power"`
			SerialNumber struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"serial_number"`
			Level struct {
				Min int `json:"min"`
				Max int `json:"max"`
			} `json:"level"`
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
		BannerImageURL string `json:"banner_image_url"`
		ChatURL        string `json:"chat_url"`
		CreatedDate    string `json:"created_date"`
		DefaultToFiat  bool   `json:"default_to_fiat"`
		Description    string `json:"description"`
		DisplayData    struct {
			CardDisplayStyle string `json:"card_display_style"`
		} `json:"display_data"`
		ExternalURL                 string      `json:"external_url"`
		Featured                    bool        `json:"featured"`
		FeaturedImageURL            interface{} `json:"featured_image_url"`
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
		PayoutAddress               string      `json:"payout_address"`
		RequireEmail                bool        `json:"require_email"`
		ShortDescription            interface{} `json:"short_description"`
		Slug                        string      `json:"slug"`
		TelegramURL                 string      `json:"telegram_url"`
		TwitterUsername             string      `json:"twitter_username"`
		InstagramUsername           interface{} `json:"instagram_username"`
		WikiURL                     interface{} `json:"wiki_url"`
	} `json:"collection"`
	Decimals   int         `json:"decimals"`
	SellOrders interface{} `json:"sell_orders"`
	Creator    struct {
		User struct {
			Username string `json:"username"`
		} `json:"user"`
		ProfileImgURL string `json:"profile_img_url"`
		Address       string `json:"address"`
		Config        string `json:"config"`
		DiscordID     string `json:"discord_id"`
	} `json:"creator"`
	Traits                  []SorareTrait `json:"traits"`
	LastSale                interface{}   `json:"last_sale"`
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

type SorareScraper struct {
	nftscraper    NFTScraper
	address       common.Address
	apiURLOpensea string
	ticker        *time.Ticker
}

func NewSorareScraper(rdb *models.RelDB) *SorareScraper {
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
	s := &SorareScraper{
		address:       common.HexToAddress("0x629A673A8242c2AC4B7B8C5D8735fbeac21A6205"),
		apiURLOpensea: "https://api.opensea.io/api/v1/",
		nftscraper:    nftScraper,
		ticker:        time.NewTicker(refreshDelay),
	}

	go s.mainLoop()
	return s
}

// mainLoop runs in a goroutine until channel s is closed.
func (scraper *SorareScraper) mainLoop() {
	err := scraper.FetchData()
	if err != nil {
		log.Error("updating nfts: ", err)
	}
	for {
		select {
		case <-scraper.ticker.C:
			err := scraper.FetchData()
			if err != nil {
				log.Error("updating nfts: ", err)
			}
		case <-scraper.nftscraper.shutdown: // user requested shutdown
			log.Printf("Sorare scraper shutting down")
			err := scraper.Close()
			scraper.cleanup(err)
			return
		}
	}
}

func (scraper *SorareScraper) FetchData() error {
	totalSupply, err := scraper.GetTotalSupply()
	if err != nil {
		return err
	}

	var creatorAddress common.Address
	var creationTime time.Time
	// NFT class from DB
	nftClassID, err := scraper.nftscraper.relDB.GetNFTClassID(scraper.address.Hex(), dia.Ethereum)
	if err != nil {
		log.Error("getting nftclass ID: ", err)
	}
	sorareNFTClass, err := scraper.nftscraper.relDB.GetNFTClassByID(nftClassID)
	if err != nil {
		log.Error("getting nft by ID: ", err)
	}

	fmt.Println("total supply: ", int(totalSupply.Int64()))
	for i := 0; i < int(totalSupply.Int64()); i++ {

		var out SorareOutput
		// 1. fetch data from onchain
		tok, err := scraper.TokenByIndex(big.NewInt(int64(i)))
		if err != nil {
			log.Errorf("Error getting token ID: %+v", err)
		}
		out.Card, err = scraper.GetCard(tok)
		if err != nil {
			log.Errorf("Error getting sorare card %d: %+v", tok, err)
		}
		out.Player, err = scraper.GetPlayer(out.Card.PlayerId)
		if err != nil {
			log.Errorf("Error getting player %d: %+v", out.Card.PlayerId, err)
		}

		sorareClub, err := scraper.GetClub(out.Card.ClubId)
		if err != nil {
			log.Errorf("Error getting club %d: %+v", out.Card.ClubId, err)
		}
		out.Club = Club{
			NameClub:    sorareClub.Name,
			CountryClub: sorareClub.Country,
			CityClub:    sorareClub.City,
			YearFounded: sorareClub.YearFounded,
		}

		tokenURI, err := scraper.GetTokenURI(tok)
		if err != nil {
			log.Errorf("Error getting token URI for %d: %+v", tok, err)
		}

		// 2. fetch data from offchain
		out.Traits, creatorAddress, creationTime, err = scraper.GetOpenSeaPlayer(tok)
		if err != nil {
			log.Errorf("Error getting Opensea data: %+v", err)
		}
		// 3. combine both in order to fill dia.NFT
		result := structs.Map(out)
		// Set output object
		sorareNFT := dia.NFT{
			TokenID:        tok.String(),
			Attributes:     result,
			CreatorAddress: creatorAddress.Hex(),
			CreationTime:   creationTime,
			NFTClass:       sorareNFTClass,
			URI:            tokenURI,
		}
		log.Info("fetched nft: ", sorareNFT)
		scraper.GetDataChannel() <- sorareNFT
	}
	return nil
}

// GetTotalSupply returns the total supply of the NFT from on-chain.
func (scraper *SorareScraper) GetTotalSupply() (*big.Int, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TotalSupply(&bind.CallOpts{})
}

// GetTokenURI returns the token URI.
func (scraper *SorareScraper) GetTokenURI(index *big.Int) (string, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TokenURI(&bind.CallOpts{}, index)
}

// TokenByIndex returns the token address from on-chain.
func (scraper *SorareScraper) TokenByIndex(index *big.Int) (*big.Int, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.TokenByIndex(&bind.CallOpts{}, index)
}

// GetCard returns data for a given card from on-chain.
func (scraper *SorareScraper) GetCard(index *big.Int) (SorareCard, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.GetCard(&bind.CallOpts{}, index)
}

// GetPlayer returns data for a given player from on-chain.
func (scraper *SorareScraper) GetPlayer(index *big.Int) (SorarePlayer, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.GetPlayer(&bind.CallOpts{}, index)
}

// GetClub returns data for a given club from on-chain.
func (scraper *SorareScraper) GetClub(index uint16) (SorareClub, error) {
	contract, err := sorare.NewSorareTokensCaller(scraper.address, scraper.nftscraper.ethConnection)
	if err != nil {
		fmt.Println("error getting contract: ", err)
	}
	return contract.GetClub(&bind.CallOpts{}, index)
}

// GetOpenSeaPlayer returns the scraped data from Opensea for a given player
func (scraper *SorareScraper) GetOpenSeaPlayer(index *big.Int) ([]SorareTrait, common.Address, time.Time, error) {
	var creatorAddress common.Address
	var creationTime time.Time
	openseaURL := scraper.apiURLOpensea + "asset/" + scraper.address.String() + "/" + index.String()
	respData, statusCode, err := utils.OpenseaGetRequest(openseaURL)
	log.Info("statusCode, err: ", statusCode, err)

	count := 0
	for statusCode != 200 && count < 40 {
		// Retry get request
		log.Infof("sleep for %v seconds", count)
		time.Sleep(time.Millisecond * time.Duration(openseaAPIWait*count))
		respData, statusCode, err = utils.GetRequestWithStatus(openseaURL)
		log.Info("statusCode, err in retry: ", statusCode, err)
		count++
	}

	traits, err := GetPlayerTraits(respData)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	creatorAddress, err = GetCreatorAddress(respData)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	creationTime, err = GetCreationTime(respData)
	if err != nil {
		return nil, creatorAddress, creationTime, err
	}

	return traits, creatorAddress, creationTime, nil
}

// GetPlayerTraits returns the parsed traits data from Opensea for a given player
func GetPlayerTraits(playerResp []byte) ([]SorareTrait, error) {
	var resp OpenSeaResponse
	if err := json.Unmarshal(playerResp, &resp); err != nil {
		return nil, err
	}
	return resp.Traits, nil
}

// GetCreatorAddress returns the creator address from Opensea
func GetCreatorAddress(playerResp []byte) (common.Address, error) {
	var resp OpenSeaResponse
	var address common.Address
	if err := json.Unmarshal(playerResp, &resp); err != nil {
		return address, err
	}

	address = common.HexToAddress(resp.Creator.Address)

	return address, nil
}

// GetCreationTime returns the creation time from Opensea
func GetCreationTime(playerResp []byte) (time.Time, error) {
	var resp OpenSeaResponse
	var t time.Time
	if err := json.Unmarshal(playerResp, &resp); err != nil {
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
func (scraper *SorareScraper) GetDataChannel() chan dia.NFT {
	return scraper.nftscraper.chanData
}

// closes all connected Scrapers. Must only be called from mainLoop
func (scraper *SorareScraper) cleanup(err error) {
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
func (scraper *SorareScraper) Close() error {
	if scraper.nftscraper.closed {
		return errors.New("scraper already closed")
	}
	close(scraper.nftscraper.shutdown)
	<-scraper.nftscraper.shutdownDone
	scraper.nftscraper.errorLock.RLock()
	defer scraper.nftscraper.errorLock.RUnlock()
	return scraper.nftscraper.error
}
