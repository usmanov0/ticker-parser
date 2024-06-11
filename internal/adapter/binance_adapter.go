package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type BinanceRepository struct {
	client *http.Client
}

type PriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func NewBinanceRepository() *BinanceRepository {
	return &BinanceRepository{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (r *BinanceRepository) FetchPrice(symbol string) (string, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	resp, err := r.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var priceResp PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResp); err != nil {
		return "", err
	}

	return priceResp.Price, nil
}
