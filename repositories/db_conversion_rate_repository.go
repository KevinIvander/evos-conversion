package repositories

import (
	"errors"
	"evos-conversion/models"
	"strconv"

	"github.com/jinzhu/gorm"
)

// dbConversionRateRepository implements DBConversionRateRepositoryContract
type dbConversionRateRepository struct {
	db *gorm.DB
}

// NewDBConversionRateRepository returns new DBConversionRateRepositoryContract implementation
func NewDBConversionRateRepository(db *gorm.DB) DBConversionRateRepositoryContract {
	return &dbConversionRateRepository{
		db: db,
	}
}

// Get :nodoc:
func (repository *dbConversionRateRepository) Get(limit int, offset int) ([]models.ConversionRate, error) {
	var conversionRates []models.ConversionRate
	if limit == 0 {
		repository.db.
			Find(&conversionRates)
	} else {
		repository.db.
			Limit(limit).
			Offset(offset).
			Find(&conversionRates)
	}
	if len(conversionRates) <= 0 {
		return []models.ConversionRate{}, errors.New("conversion rate not found")
	}
	return conversionRates, nil
}

// Store :nodoc:
func (repository *dbConversionRateRepository) Store(data models.ConversionRate) (models.ConversionRate, error) {
	if err := repository.db.Save(&data).Error; err != nil {
		return models.ConversionRate{}, err
	}
	return data, nil
}

// Convert :nodoc:
func (repository *dbConversionRateRepository) Convert(idFrom int, idTo int, amount float32) (float32, error) {
	check := ""
	if idFrom < idTo {
		check = strconv.Itoa(idFrom) + strconv.Itoa(idTo)
	} else if idFrom > idTo {
		check = strconv.Itoa(idTo) + strconv.Itoa(idFrom)
	} else {
		return amount, nil
	}
	var conversionRate models.ConversionRate
	repository.db.First(&conversionRate, "conversion_rates.check = ?", check)
	if conversionRate.ID == 0 {
		return 0.0, errors.New("conversion rate not found")
	}
	result := 0.0
	if conversionRate.CurrencyFromID == idFrom {
		result = float64(amount) * float64(conversionRate.Rate)
	} else {
		result = float64(amount) / float64(conversionRate.Rate)
	}
	return float32(result), nil
}
