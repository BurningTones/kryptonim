package helpers

import (
	"fmt"
)

func CalculateExchangeRate(amount float64, rate float64) float64 {
	return amount * rate
}

func CalculateExchangedAmount(amount float64, fromRate float64, toRate float64, decimalPlaces int) (string, error) {
	if fromRate == 0 || toRate == 0 {
		return "", fmt.Errorf("invalid rates")
	}
	amountInUSD := amount * fromRate
	exchangedAmount := amountInUSD / toRate
	return fmt.Sprintf("%.*f", decimalPlaces, exchangedAmount), nil
}
