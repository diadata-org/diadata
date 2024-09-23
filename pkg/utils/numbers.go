package utils

import (
	"math/big"

	"github.com/shopspring/decimal"
)

func StringToFloat64(value string, decimals int64) (float64, bool) {
	bigInt := new(big.Int)
	bigInt.SetString(value, 10)
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(bigInt.String())
	result := num.Div(mul)
	return result.Float64()
}

func Min(n1 uint64, n2 uint64) uint64 {
	if n1 <= n2 {
		return n1
	}
	return n2
}
