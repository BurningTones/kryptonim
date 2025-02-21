package services

import (
	"net/http"
	"strconv"

	"kryptonim/app/helpers"

	"github.com/gin-gonic/gin"
)

type ExchangeService interface {
	HandleExchange(c *gin.Context)
}

type exchangeServiceImpl struct{}

var rates = map[string]struct {
	DecimalPlaces int
	RateToUSD     float64
}{
	"BEER":  {18, 0.00002461},
	"FLOKI": {18, 0.0001428},
	"GATE":  {18, 6.87},
	"USDT":  {6, 0.999},
	"WBTC":  {8, 57037.22},
}

func getRatesFromAPI() (map[string]struct {
	DecimalPlaces int
	RateToUSD     float64
}, error) {
	// To be implemented
	// For now, returning the mock rates
	return rates, nil
}

func (s *exchangeServiceImpl) HandleExchange(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	amountStr := c.Query("amount")

	if from == "" || to == "" || amountStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	rates, err := getRatesFromAPI()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch rates"})
		return
	}

	fromRate, fromExists := rates[from]
	toRate, toExists := rates[to]

	if !fromExists || !toExists {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	exchangedAmount, err := helpers.CalculateExchangedAmount(amount, fromRate.RateToUSD, toRate.RateToUSD, toRate.DecimalPlaces)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "calculation error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"from":   from,
		"to":     to,
		"amount": exchangedAmount,
	})
}

func NewExchangeService() ExchangeService {
	return &exchangeServiceImpl{}
}
