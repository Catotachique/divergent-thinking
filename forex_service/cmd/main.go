package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Define a structure that matches the expected JSON format
type ForexResponse struct {
	Success   bool   `json:"success"`
	Base      string `json:"base"`
	Timestamp int64  `json:"timestamp"`
	Rates     struct {
		EUR float64 `json:"EUR"`
	} `json:"rates"`
}

// LoadEnv manually loads variables from the .env file into the environment
func loadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || line[0] == '#' { // Skip empty lines or comments
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return scanner.Err()
}

func main() {
	// Load environment variables from .env file
	if err := loadEnv(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the API key from environment variables
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is required")
	}

	// Build the URL with parameters
	url := fmt.Sprintf("https://api.forexrateapi.com/v1/latest?api_key=%s&base=USD&currencies=EUR", apiKey)

	// Make the GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var forexResponse ForexResponse
	if err := json.NewDecoder(resp.Body).Decode(&forexResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Check if the API call was successful
	if forexResponse.Success {
		fmt.Printf("Base Currency: %s\n", forexResponse.Base)
		fmt.Printf("Exchange Rate (EUR): %.6f\n", forexResponse.Rates.EUR)
	} else {
		fmt.Println("Error fetching forex data.")
	}
}
