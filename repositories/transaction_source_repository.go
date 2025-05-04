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

func constructTransactionSourceWhereCondition(query *gorm.DB, whereCondition models.TransactionSourceWhere) *gorm.DB {
	if whereCondition.ID > 0 {
		query = query.Where("id = ?", whereCondition.ID)
	}

	if whereCondition.Type != "" {
		query = query.Where("type = ?", whereCondition.Type)
	}

	return query
}

func (r *TransactionSourceRepository) GetAll(whereCondition models.TransactionSourceWhere) ([]models.TransactionSource, error) {
	var sources []models.TransactionSource

	query := r.DB.Model(&models.TransactionSource{})

	query = constructTransactionSourceWhereCondition(query, whereCondition)

	if err := query.Find(&sources).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return sources, nil
}

func (r *TransactionSourceRepository) GetByID(sourceID uint) (models.TransactionSource, error) {
	var source models.TransactionSource

	query := r.DB.Model(&models.TransactionSource{})

	query = constructTransactionSourceWhereCondition(query, models.TransactionSourceWhere{
		ID: sourceID,
	})

	if err := query.First(&source).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return source, err
	}

	return source, nil
}
