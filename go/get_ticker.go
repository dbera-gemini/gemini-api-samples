package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func getTicker(symbol string) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")

	req, err := http.NewRequest("GET", baseURL+"/ticker", nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	q := req.URL.Query()
	q.Add("symbol", symbol)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
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
	if err := getTicker("BTCUSD"); err != nil {
		fmt.Println("Error:", err)
	}
}
