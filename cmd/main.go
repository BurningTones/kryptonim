package main

import (
	"kryptonim/app/container"
	"kryptonim/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Build the container
	cont := container.BuildContainer()

	// Middleware for routing and token-based authorization
	r := routes.SetupRouter(cont)

	log.Println("Starting server on :8080")
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
