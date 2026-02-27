package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type OrderParams struct {
	Symbol   string  `json:"symbol"`
	Side     string  `json:"side"`
	Type     string  `json:"type"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price,omitempty"`
}

func placeOrder(params OrderParams) error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	baseURL := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	jsonData, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	req, err := http.NewRequest("POST", baseURL+"/order", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("X-API-Key", apiKey)
	req.Header.Add("X-API-Secret", apiSecret)
	req.Header.Add("Content-Type", "application/json")

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

	fmt.Println("Order placed:", result)
	return nil
}

func main() {
	order := OrderParams{
		Symbol:   "BTCUSD",
		Side:     "buy",
		Type:     "limit",
		Quantity: 0.01,
		Price:    50000,
	}

	if err := placeOrder(order); err != nil {
		fmt.Println("Error:", err)
	}
}
