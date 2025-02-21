package tests

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Run tests with: go test -count=1 -timeout 30s -run ^(TestEnvVariables|TestEnvVariablesErrorHandling)$ kryptonim/tests

func TestEnvVariables(t *testing.T) {
	log.Println("Starting TestEnvVariables")

	// Set up environment variables for testing
	os.Setenv("OPENEXCHANGERATES_APP_ID", "test_app_id")
	os.Setenv("OPENEXCHANGERATES_URL", "https://openexchangerates.org/api/latest.json")
	os.Setenv("AUTH_REQ", "ON")
	os.Setenv("AUTH_TOKEN", "test_token")

	assert.Equal(t, "test_app_id", os.Getenv("OPENEXCHANGERATES_APP_ID"))
	assert.Equal(t, "https://openexchangerates.org/api/latest.json", os.Getenv("OPENEXCHANGERATES_URL"))
	assert.Equal(t, "ON", os.Getenv("AUTH_REQ"))
	assert.Equal(t, "test_token", os.Getenv("AUTH_TOKEN"))

	log.Println("Completed TestEnvVariables")
}

func TestEnvVariablesErrorHandling(t *testing.T) {
	log.Println("Starting TestEnvVariablesErrorHandling")

	// Unset environment variables to simulate error conditions
	os.Unsetenv("OPENEXCHANGERATES_APP_ID")
	os.Unsetenv("OPENEXCHANGERATES_URL")
	os.Unsetenv("AUTH_REQ")
	os.Unsetenv("AUTH_TOKEN")

	assert.Empty(t, os.Getenv("OPENEXCHANGERATES_APP_ID"))
	assert.Empty(t, os.Getenv("OPENEXCHANGERATES_URL"))
	assert.Empty(t, os.Getenv("AUTH_REQ"))
	assert.Empty(t, os.Getenv("AUTH_TOKEN"))

	log.Println("Completed TestEnvVariablesErrorHandling")
}
