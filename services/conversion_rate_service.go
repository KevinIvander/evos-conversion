package services

import (
	"evos-conversion/models"
	"evos-conversion/repositories"
)

type conversionRateService struct {
	dbConversionRateRepository repositories.DBConversionRateRepositoryContract
}

// NewConversionRateService returns new ConversionRateServiceContract implementation
func NewConversionRateService(
	dbConversionRateRepository repositories.DBConversionRateRepositoryContract,
) ConversionRateServiceContract {
	return &conversionRateService{
		dbConversionRateRepository: dbConversionRateRepository,
	}
}

// Get :nodoc:
func (service *conversionRateService) Get(limit int, offset int) ([]models.ConversionRate, error) {
	conversionRateListFromDB, err := service.dbConversionRateRepository.Get(limit, offset)
	if err != nil {
		return []models.ConversionRate{}, err
	}

	return conversionRateListFromDB, err
}

// Create :nodoc:
func (service conversionRateService) Create(data models.ConversionRate) (models.ConversionRate, error) {
	newConversionRate, err := service.dbConversionRateRepository.Store(data)
	if err != nil {
		return models.ConversionRate{}, err
	}

	return newConversionRate, nil
}

// Convert :nodoc:
func (service conversionRateService) Convert(idFrom int, idTo int, amount float32) (float32, error) {
	result, err := service.dbConversionRateRepository.Convert(idFrom, idTo, amount)
	if err != nil {
		return 0.0, err
	}

	return result, nil
}
