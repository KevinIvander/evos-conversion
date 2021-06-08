package requests

import (
	"errors"
	"net/http"
)

// CreateConversionRateRequest holds data for conversion rate creation
type CreateConversionRateRequest struct {
	CurrencyFromID int     `json:"currency_from_id"`
	CurrencyToID   int     `json:"currency_to_id"`
	Rate           float32 `json:"rate"`
}

// ConvertConversionRateRequest holds data for conversion rate creation
type ConvertConversionRateRequest struct {
	CurrencyFromID int     `json:"currency_from_id"`
	CurrencyToID   int     `json:"currency_to_id"`
	Amount         float32 `json:"amount"`
}

// Bind :nodoc:
func (req CreateConversionRateRequest) Bind(request *http.Request) error {
	if req.CurrencyFromID == 0 {
		return errors.New("`currency_from_id` is required")
	}
	if req.CurrencyToID == 0 {
		return errors.New("`currency_to_id` is required")
	}
	if req.Rate == 0 {
		return errors.New("`rate` is required")
	}
	return nil
}

// Bind :nodoc:
func (req ConvertConversionRateRequest) Bind(request *http.Request) error {
	if req.CurrencyFromID == 0 {
		return errors.New("`currency_from_id` is required")
	}
	if req.CurrencyToID == 0 {
		return errors.New("`currency_to_id` is required")
	}
	if req.Amount == 0 {
		return errors.New("`amount` is required")
	}
	return nil
}
