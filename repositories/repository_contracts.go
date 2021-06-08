package repositories

import "evos-conversion/models"

// DBCurrencyRepositoryContract represents sets of functions for
// manipulating currency data on database
type DBCurrencyRepositoryContract interface {
	// Get returns currency data with pagination
	Get(limit int, offset int) ([]models.Currency, error)

	// Store saves new currency
	Store(data models.Currency) (models.Currency, error)
}

// DBConversionRateRepositoryContract represents sets of functions for
// manipulating conversion rate data on database
type DBConversionRateRepositoryContract interface {
	// Get returns conversion rate data with pagination
	Get(limit int, offset int) ([]models.ConversionRate, error)

	// Store saves new conversion rate
	Store(data models.ConversionRate) (models.ConversionRate, error)

	// Convert data with conversion rate
	Convert(idFrom int, idTo int, amount float32) (float32, error)
}
