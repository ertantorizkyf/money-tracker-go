package repositories

import (
	"fmt"

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
	if whereCondition.Type != "" {
		query = query.Where("type = ?", whereCondition.Type)
	}

	return query
}

func (r *TransactionCategoryRepository) GetAll(whereCondition models.TransactionCategoryWhere) ([]models.TransactionCategory, error) {
	var categories []models.TransactionCategory

	fmt.Println("=== DEBUG: ", whereCondition.Type)

	query := r.DB

	query = constructTransactionCategoryWhereCondition(query, whereCondition)

	if err := query.Find(&categories).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return categories, nil
}
