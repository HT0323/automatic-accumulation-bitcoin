package utils

import "math"

func RoundDecimal(num float64) float64 {
	return math.Round(num)
}

func roundUp(num, places float64) float64 {
	shift := math.Pow(10, places)
	return RoundDecimal(num*shift) / shift
}

func CalcAmount(price, budget, minimumAmount, places float64) float64 {
	amount := roundUp(budget/price, places)
	if amount < minimumAmount {
		return minimumAmount
	} else {
		return amount
	}
}
