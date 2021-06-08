package responses

import (
	"evos-conversion/models"
	"net/http"
)

// CurrencyListResponse represents response of multiple currency
type CurrencyListResponse struct {
	Data   []models.Currency `json:"data"`
	Limit  int               `json:"limit"`
	Offset int               `json:"offset"`
}

// CreateCurrencyResponse holds data for currency creation
type CreateCurrencyResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// CurrencyAcceptedResponse represents response for single currency
type CurrencyAcceptedResponse struct {
	Data CreateCurrencyResponse `json:"data"`
}

// NewCurrencyListResponse returns new CurrencyListResponse
func NewCurrencyListResponse(data []models.Currency, limit int, offset int) CurrencyListResponse {
	return CurrencyListResponse{
		Data:   data,
		Limit:  limit,
		Offset: offset,
	}
}

// NewCurrencyAcceptedResponse returns new CurrencyResponse
func NewCurrencyAcceptedResponse(data CreateCurrencyResponse) CurrencyAcceptedResponse {
	return CurrencyAcceptedResponse{
		Data: data,
	}
}

// Render :nodoc:
func (res CurrencyAcceptedResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// Render :nodoc:
func (res CurrencyListResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
