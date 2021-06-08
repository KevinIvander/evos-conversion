package responses

import (
	"evos-conversion/models"
	"net/http"
)

// ConversionRateListResponse represents response of multiple conversion rate
type ConversionRateListResponse struct {
	Data   []models.ConversionRate `json:"data"`
	Limit  int                     `json:"limit"`
	Offset int                     `json:"offset"`
}

// ConversionRateAcceptedResponse represents response for single conversion rate
type ConversionRateAcceptedResponse struct {
	Data models.ConversionRate `json:"data"`
}

// ConvertResponse represents response for conversion data
type ConvertResponse struct {
	Result float32 `json:"result"`
}

// NewConversionRateListResponse returns new ConversionRateListResponse
func NewConversionRateListResponse(data []models.ConversionRate, limit int, offset int) ConversionRateListResponse {
	return ConversionRateListResponse{
		Data:   data,
		Limit:  limit,
		Offset: offset,
	}
}

// NewConversionRateAcceptedResponse returns new ConversionRateResponse
func NewConversionRateAcceptedResponse(data models.ConversionRate) ConversionRateAcceptedResponse {
	return ConversionRateAcceptedResponse{
		Data: data,
	}
}

// NewConvertResponse returns new conversion amount
func NewConvertResponse(result float32) ConvertResponse {
	return ConvertResponse{
		Result: result,
	}
}

// Render :nodoc:
func (res ConversionRateAcceptedResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// Render :nodoc:
func (res ConversionRateListResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// Render :nodoc:
func (res ConvertResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
