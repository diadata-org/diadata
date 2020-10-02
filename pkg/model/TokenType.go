package models

import (
	"time"
)

type Token struct {
	Symbol      string
	Name        string
	TotalSupply float64
	Decimals    int
	Source      string
	Time        time.Time
}
