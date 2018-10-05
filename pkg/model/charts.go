package models

import (
	"fmt"
	"github.com/diadata-org/api-golang/dia"
	log "github.com/sirupsen/logrus"
)

func (db *DB) GetChartPoints(symbol string) ([]Point, error) {
	result := []Point{}
	vals, err := db.redisClient.ZRange(getKeyFilterZSET(dia.FilterKing+"_"+symbol), 0, -1).Result()
	if err == nil {
		var p Point
		for _, v := range vals {
			fmt.Sscanf(v, "%f %d", &p.Value, &p.UnixTime)
			result = append(result, p)
		}
	}
	log.Println(symbol, "GetChartPoints:%v", result)
	return result, err
}
