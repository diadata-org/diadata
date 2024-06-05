package utils

import (
	"context"
	"errors"
	"strconv"
	"time"

	gql "github.com/machinebox/graphql"
)

func GetGraphqlAssetQuotationFromDia(blockchain, address string, blockDuration int) (float64, error) {
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-blockDuration*2) * time.Second)
	type Response struct {
		GetFeed []struct {
			Name  string    `json:"Name"`
			Time  time.Time `json:"Time"`
			Value float64   `json:"Value"`
		} `json:"GetFeed"`
	}
	client := gql.NewClient("https://api.diadata.org" + "/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetFeed(
		 	Filter: "mair", 
			BlockSizeSeconds: ` + strconv.Itoa(blockDuration) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(blockDuration) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			FeedSelection: [{
			  Address: "` + address + `", 
			  Blockchain: "` + blockchain + `"
		  },
		  ],
		  ) {
				Name
				Time
				Value
	  	}
		}`)

	ctx := context.Background()
	var r Response
	if err := client.Run(ctx, req, &r); err != nil {
		return 0.0, err
	}
	if len(r.GetFeed) == 0 {
		return 0.0, errors.New("no results")
	}
	return r.GetFeed[len(r.GetFeed)-1].Value, nil
}
