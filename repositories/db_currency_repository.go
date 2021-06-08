package repositories

import (
	"errors"
	"evos-conversion/models"

	"github.com/jinzhu/gorm"
)

// dbCurrencyRepository implements DBCurrencyRepositoryContract
type dbCurrencyRepository struct {
	db *gorm.DB
}

// NewDBCurrencyRepository returns new DBCurrencyRepositoryContract implementation
func NewDBCurrencyRepository(db *gorm.DB) DBCurrencyRepositoryContract {
	return &dbCurrencyRepository{
		db: db,
	}
}

// Get :nodoc:
func (repository *dbCurrencyRepository) Get(limit int, offset int) ([]models.Currency, error) {
	var currencies []models.Currency
	if limit == 0 {
		repository.db.
			Find(&currencies)
	} else {
		repository.db.
			Limit(limit).
			Offset(offset).
			Find(&currencies)
	}
	if len(currencies) <= 0 {
		return []models.Currency{}, errors.New("currency not found")
	}

	return currencies, nil
}

// Store :nodoc:
func (repository *dbCurrencyRepository) Store(data models.Currency) (models.Currency, error) {
	if err := repository.db.Save(&data).Error; err != nil {
		return models.Currency{}, err
	}
	return data, nil
}
