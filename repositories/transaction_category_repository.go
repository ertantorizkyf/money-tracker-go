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

func constructTransactionCategoryWhereCondition(query *gorm.DB, whereCondition models.TransactionCategoryWhere) *gorm.DB {
	if whereCondition.ID > 0 {
		query = query.Where("id = ?", whereCondition.ID)
	}

	if whereCondition.Type != "" {
		query = query.Where("type = ?", whereCondition.Type)
	}

	return query
}

func (r *TransactionCategoryRepository) GetAll(whereCondition models.TransactionCategoryWhere) ([]models.TransactionCategory, error) {
	var categories []models.TransactionCategory

	query := r.DB.Model(&models.TransactionCategory{})

	query = constructTransactionCategoryWhereCondition(query, whereCondition)

	if err := query.Find(&categories).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return categories, nil
}

func (r *TransactionCategoryRepository) GetByID(categoryID uint) (models.TransactionCategory, error) {
	var category models.TransactionCategory

	query := r.DB.Model(&models.TransactionCategory{})

	query = constructTransactionCategoryWhereCondition(query, models.TransactionCategoryWhere{
		ID: categoryID,
	})

	if err := query.First(&category).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return category, err
	}

	return category, nil
}
