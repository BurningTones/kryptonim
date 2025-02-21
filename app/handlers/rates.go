package handlers

import (
	"kryptonim/app/container"
	"kryptonim/app/helpers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRates(container *container.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		currencies := c.Query("currencies")
		if currencies == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "currencies query parameter is required"})
			return
		}

		currencyList := helpers.ParseCurrencies(currencies)
		rates, err := container.RatesService.GetRates(c, currencyList)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, rates)
	}
}
