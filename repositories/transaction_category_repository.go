package repositories

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"gorm.io/gorm"
)

type TransactionCategoryRepository struct {
	DB *gorm.DB
}

func NewTransactionCategoryRepository() *TransactionCategoryRepository {
	return &TransactionCategoryRepository{
		DB: initializers.DB,
	}
}

func (r *TransactionCategoryRepository) GetAll() ([]models.TransactionCategory, error) {
	var categories []models.TransactionCategory

	err := r.DB.Find(&categories).Error
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return categories, nil
}
