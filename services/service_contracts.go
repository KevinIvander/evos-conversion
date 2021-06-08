package services

import "evos-conversion/models"

// CurrencyServiceContract handles business logic for currency
type CurrencyServiceContract interface {
	// Get returns currency with pagination
	Get(limit int, offset int) ([]models.Currency, error)

	// Create saves new currency
	Create(data models.Currency) (models.Currency, error)
}

// ConversionRateServiceContract handles business logic for conversion rate
type ConversionRateServiceContract interface {
	// Get returns conversion rate with pagination
	Get(limit int, offset int) ([]models.ConversionRate, error)

	// Create saves new ConversionRate
	Create(data models.ConversionRate) (models.ConversionRate, error)

	// Convert transform data with conversion rate
	Convert(idFrom int, idTo int, amount float32) (float32, error)
}
