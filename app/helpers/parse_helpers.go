package helpers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func ParseCurrencies(currencies string) []string {
	return strings.Split(currencies, ",")
}

func ParseRates(body []byte, currencies []string) []gin.H {
	var response map[string]interface{}
	json.Unmarshal(body, &response)
	rates := response["rates"].(map[string]interface{})

	var result []gin.H
	loggedCurrencies := make(map[string]bool)

	for _, from := range currencies {
		fromRate, fromExists := rates[from]
		if !fromExists {
			if !loggedCurrencies[from] {
				log.Printf("Currency %s not found in rates\n", from)
				loggedCurrencies[from] = true
			}
			continue
		}
		for _, to := range currencies {
			if from != to {
				toRate, toExists := rates[to]
				if !toExists {
					if !loggedCurrencies[to] {
						log.Printf("Currency %s not found in rates\n", to)
						loggedCurrencies[to] = true
					}
					continue
				}
				rate := toRate.(float64) / fromRate.(float64)
				result = append(result, gin.H{
					"from": from,
					"to":   to,
					"rate": rate,
				})
			}
		}
	}
	return result
}

func JoinCurrencies(currencies []string) string {
	result := ""
	for _, currency := range currencies {
		if result == "" {
			result = currency
		} else {
			result += "," + currency
		}
	}
	return result
}
