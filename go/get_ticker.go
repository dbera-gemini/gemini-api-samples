package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getBaseURL() string {
	// Load .env file (ignore error if file doesn't exist)
	_ = godotenv.Load()

	baseURL := os.Getenv("GEMINI_BASE_URL")
	if baseURL == "" {
		baseURL = "https://api.gemini.com/v1"
	}
	return baseURL
}

func getTicker(symbol string) error {
	// Symbol should be lowercase for Gemini API
	symbolLower := strings.ToLower(symbol)
	url := fmt.Sprintf("%s/pubticker/%s", getBaseURL(), symbolLower)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	fmt.Println("Ticker data:", result)
	return nil
}

func main() {
	args := os.Args[1:]

	// Show help message
	if len(args) > 0 && (args[0] == "--help" || args[0] == "-h") {
		fmt.Println("Usage: go run get_ticker.go [symbol]")
		fmt.Println("Example: go run get_ticker.go ethusd")
		fmt.Println("Default symbol: btcusd")
		os.Exit(0)
	}

	// Get symbol from command line or use default
	symbol := "btcusd"
	if len(args) > 0 {
		symbol = args[0]
	}

	if err := getTicker(symbol); err != nil {
		fmt.Println("Error:", err)
	}
}
