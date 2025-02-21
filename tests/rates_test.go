package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"kryptonim/app/container"
	"kryptonim/app/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetRates(t *testing.T) {
	log.Println("Starting TestGetRates")

	// Set up environment variables for testing
	os.Setenv("OPENEXCHANGERATES_APP_ID", "5ca430924b12402f881b5c5088e220b3")
	os.Setenv("OPENEXCHANGERATES_URL", "https://openexchangerates.org/api/latest.json")
	os.Setenv("AUTH_TOKEN", "kryp+0N!m") // Ensure AUTH_TOKEN is set if required

	router := gin.Default()
	cont := container.BuildContainer()
	router.GET("/rates", handlers.GetRates(cont))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates?currencies=USD,GBP,EUR", nil)
	req.Header.Set("Auth", "kryp+0N!m") // Set authorization header if required
	router.ServeHTTP(w, req)

	log.Printf("Response Code: %d", w.Code)
	log.Printf("Response Body: %s", w.Body.String())

	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions as needed

	log.Println("Completed TestGetRates")
}

func TestGetRatesMissingCurrencies(t *testing.T) {
	log.Println("Starting TestGetRatesMissingCurrencies")

	// Set up environment variables for testing
	os.Setenv("OPENEXCHANGERATES_APP_ID", "5ca430924b12402f881b5c5088e220b3")
	os.Setenv("OPENEXCHANGERATES_URL", "https://openexchangerates.org/api/latest.json")

	router := gin.Default()
	cont := container.BuildContainer()
	router.GET("/rates", handlers.GetRates(cont))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/rates", nil)
	router.ServeHTTP(w, req)

	log.Printf("Response Code: %d", w.Code)
	log.Printf("Response Body: %s", w.Body.String())

	assert.Equal(t, http.StatusBadRequest, w.Code)
	// Add more assertions as needed

	log.Println("Completed TestGetRatesMissingCurrencies")
}
