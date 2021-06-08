package models

// ConversionRate represents conversion rate data
type ConversionRate struct {
	ID             uint    `json:"id"`
	CurrencyFromID int     `json:"currency_from_id"`
	CurrencyToID   int     `json:"currency_to_id"`
	Rate           float32 `json:"rate"`
}
