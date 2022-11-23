package utils

import "errors"

const (
	UNIV2_FEE_FACTOR = 997
	UNIV2_FACTOR     = 1000
)

func UniV2AmountOut(amountIn float64, reserveIn float64, reserveOut float64) (float64, error) {
	amountInWithFee := amountIn * UNIV2_FEE_FACTOR
	if reserveIn <= 0 || amountInWithFee <= 0 {
		return 0, errors.New("reserve in and amount in are not positive numbers.")
	}
	return amountInWithFee * reserveOut / ((reserveIn * UNIV2_FACTOR) + amountInWithFee), nil
}

func UniV2PriceReserveIn(reserveIn float64, reserveOut float64) (float64, error) {
	if reserveIn == 0 {
		return 0, errors.New("reserveIn is not positive number.")
	}
	return reserveOut / reserveIn, nil
}

func UniV2PriceAfterTrade(amountIn float64, reserveIn float64, reserveOut float64) (price float64, err error) {
	if reserveIn == 0 {
		return 0, errors.New("reserveIn is not positive number.")
	}
	amountOut, err := UniV2AmountOut(amountIn, reserveIn, reserveOut)
	if err != nil {
		return 0, err
	}
	newReserveOut := reserveOut - amountOut
	return newReserveOut / (amountIn + reserveIn), nil
}
