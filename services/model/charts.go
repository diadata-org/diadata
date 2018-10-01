package models

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Point struct {
	UnixTime int64
	Value    float64
}

func (db *DB) GetChartPoints(symbol string) ([]Point, error) {
	result := []Point{}
	vals, err := db.redisClient.ZRange(getKeyFilterZSET("MA120_"+symbol), 0, -1).Result()
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
