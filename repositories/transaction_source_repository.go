package repositories

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"gorm.io/gorm"
)

type TransactionSourceRepository struct {
	DB *gorm.DB
}

func NewTransactionSourceRepository() *TransactionSourceRepository {
	return &TransactionSourceRepository{
		DB: initializers.DB,
	}
}

func (r *TransactionSourceRepository) GetAll() ([]models.TransactionSource, error) {
	var sources []models.TransactionSource

	err := r.DB.Find(&sources).Error
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return sources, nil
}
