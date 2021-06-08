package services

import (
	"evos-conversion/models"
	"evos-conversion/repositories"
)

type currencyService struct {
	dbCurrencyRepository repositories.DBCurrencyRepositoryContract
}

// NewCurrencyService returns new CurrencyServiceContract implementation
func NewCurrencyService(
	dbCurrencyRepository repositories.DBCurrencyRepositoryContract,
) CurrencyServiceContract {
	return &currencyService{
		dbCurrencyRepository: dbCurrencyRepository,
	}
}

// Get :nodoc:
func (service *currencyService) Get(limit int, offset int) ([]models.Currency, error) {
	currencyListFromDB, err := service.dbCurrencyRepository.Get(limit, offset)
	if err != nil {
		return []models.Currency{}, err
	}

	return currencyListFromDB, err
}

// Create :nodoc:
func (service currencyService) Create(data models.Currency) (models.Currency, error) {
	newCurrency, err := service.dbCurrencyRepository.Store(data)
	if err != nil {
		return models.Currency{}, err
	}

	return newCurrency, nil
}
