package models

import (
	"fmt"
	"github.com/diadata-org/diadata/pkg/dia"
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

func (db *DB) GetChartPoints7Days(symbol string) ([]float64, error) {
	result := []float64{}
	vals, err := db.redisClient.ZRange(getKeyFilterZSET(dia.FilterKing+"_"+symbol), 0, -1).Result()
	if err == nil {
		if len(vals) != 0 {
			var lastUnixTime int64
			var lastValue float64
			var secondsInOnePoint int64 = 3600 * 24
			var numberOfPoints int64 = 8
			fmt.Sscanf(vals[len(vals)-1], "%f %d", &lastValue, &lastUnixTime)
			lastTime := lastUnixTime - secondsInOnePoint*numberOfPoints
			var value float64
			var unixTime int64
			for _, v := range vals {
				fmt.Sscanf(v, "%f %d", &value, &unixTime)
				if unixTime-lastTime > secondsInOnePoint {
					lastTime = unixTime
					result = append(result, value)
				}
			}
			result[len(result)-1] = lastValue
		}
	}
	log.Println(symbol, "GetChartPoints:%v", result)
	return result, err
}
