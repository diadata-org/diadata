package gqlclient

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	gql "github.com/machinebox/graphql"
)

func GetGraphqlAssetQuotationFromDia(blockchain, address string, blockDuration int, assetl dia.AssetList) (price float64, tradesCount float64, tradeTime time.Time, source []string, err error) {
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-blockDuration*2) * time.Second)
	type Response struct {
		GetFeed []struct {
			Name        string    `json:"Name"`
			Time        time.Time `json:"Time"`
			Value       float64   `json:"Value"`
			Pools       []string  `json:"Pools"`
			Pairs       []string  `json:"Pairs"`
			TradesCount float64   `json:"TradesCount"`
		} `json:"GetFeed"`
	}
	client := gql.NewClient("https://api.diadata.org" + "/graphql/query")
	query := `
    query  {
		 GetFeed(
		 	Filter: "vwap", 
			BlockSizeSeconds: ` + strconv.Itoa(blockDuration) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(blockDuration) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			FeedSelection: ` + generateFeedSelectionQuery(assetl) + `,
		  ) {
				Name
				Time
				Value
				TradesCount
				Pools
				Pairs
				TradesCount

	  	}
		}`

	req := gql.NewRequest(query)

	ctx := context.Background()
	var r Response
	if err = client.Run(ctx, req, &r); err != nil {
		err = errors.New("no results")
		return
	}
	if len(r.GetFeed) == 0 {
		err = errors.New("no results")
		return
	}
	source = append(source, r.GetFeed[len(r.GetFeed)-1].Pools...)
	source = append(source, r.GetFeed[len(r.GetFeed)-1].Pairs...)

	tradesCount = r.GetFeed[len(r.GetFeed)-1].TradesCount
	price = r.GetFeed[len(r.GetFeed)-1].Value
	tradeTime = r.GetFeed[len(r.GetFeed)-1].Time

	return
}

func generateFeedSelectionQuery(assetList dia.AssetList) string {
	var jsonObjects []string

	splitted := strings.Split(assetList.AssetName, "-")

	exchangepair := convertExchangePairs(assetList.Exchanges)
	jsonObject := fmt.Sprintf(`{
			Address: %q,
			Blockchain: %q,
			Exchangepairs: %s
		}`, splitted[1], splitted[0], exchangepair)

	jsonObjects = append(jsonObjects, jsonObject)

	jsonArray := "[" + strings.Join(jsonObjects, ",") + "]"

	return jsonArray
}

func convertExchangePairs(exchangePairs []dia.ExchangeList) string {
	// Create a slice to hold JSON objects
	var jsonObjects []string

	// Iterate through the exchangePairs
	for _, exchangePair := range exchangePairs {
		jsonObject := fmt.Sprintf(`{Exchange: "%s"`, exchangePair.Name)

		var quotedArr []string

		for _, item := range exchangePair.Pairs {
			quotedArr = append(quotedArr, fmt.Sprintf("\"%s\"", strings.Replace(strings.TrimSpace(item), "/", "-", -1)))
		}

		pairs := `[` + strings.Join(quotedArr, ",") + `]`

		jsonObject += `,Pairs:` + pairs

		jsonObject += "}"

		jsonObjects = append(jsonObjects, jsonObject)
	}

	jsonArray := "[" + strings.Join(jsonObjects, ",") + "]"

	return jsonArray
}
