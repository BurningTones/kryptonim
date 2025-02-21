package services

import (
	"fmt"
	"net/http"
	"os"

	"kryptonim/app/errors"
	"kryptonim/app/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type RatesService interface {
	GetRates(c *gin.Context, currencies []string) ([]gin.H, error)
}

type ratesServiceImpl struct{}

func NewRatesService() RatesService {
	return &ratesServiceImpl{}
}

func (s *ratesServiceImpl) GetRates(c *gin.Context, currencies []string) ([]gin.H, error) {
	client := resty.New()
	appID := os.Getenv("OPENEXCHANGERATES_APP_ID")
	apiURL := os.Getenv("OPENEXCHANGERATES_URL")
	if appID == "" {
		err := fmt.Errorf("OPENEXCHANGERATES_APP_ID not set")
		errors.HandleError(c, err, http.StatusInternalServerError)
		return nil, err
	}
	if apiURL == "" {
		err := fmt.Errorf("OPENEXCHANGERATES_URL not set")
		errors.HandleError(c, err, http.StatusInternalServerError)
		return nil, err
	}

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"app_id":  appID,
			"symbols": helpers.JoinCurrencies(currencies),
		}).
		Get(apiURL)
	if err != nil {
		errors.HandleError(c, err, http.StatusInternalServerError)
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		err := fmt.Errorf("failed to fetch rates")
		errors.HandleError(c, err, resp.StatusCode())
		return nil, err
	}

	rates := helpers.ParseRates(resp.Body(), currencies)
	return rates, nil
}
