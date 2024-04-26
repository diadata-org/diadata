package utils

import (
	"math/big"
	"regexp"

	"github.com/shopspring/decimal"
)

func StringToFloat64(value string, decimals int64) (float64, bool) {
	bigInt := new(big.Int)
	re := regexp.MustCompile(`0x`)
	v := re.ReplaceAllString(value, "")
	bigInt.SetString(v, 10)

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(bigInt.String())
	result := num.Div(mul)

	return result.Float64()
}
