package repositories

import (
	"time"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{
		DB: initializers.DB,
	}
}

func constructTransactionWhereCondition(query *gorm.DB, whereCondition models.TransactionWhere) {
	if whereCondition.UserID > 0 {
		query.Where("user_id = ?", whereCondition.UserID)
	}
	if whereCondition.SourceID > 0 {
		query.Where("source_id = ?", whereCondition.SourceID)
	}
	if whereCondition.CategoryID > 0 {
		query.Where("category_id = ?", whereCondition.CategoryID)
	}
	if whereCondition.Purpose != "" {
		query.Where("purpose = ?", whereCondition.Purpose)
	}
	if whereCondition.Remark != "" {
		query.Where("remark = ?", whereCondition.Remark)
	}
	if whereCondition.StartDate != "" && whereCondition.EndDate != "" {
		startDate, _ := time.Parse("2006-01-02", whereCondition.StartDate)
		endDate, _ := time.Parse("2006-01-02", whereCondition.EndDate)

		query.Where("trx_date BETWEEN ? AND ?", startDate, endDate)
	}
	if whereCondition.Type != "" {
		query.Where("type = ?", whereCondition.Type)
	}
}

func (r *TransactionRepository) GetAll(whereCondition models.TransactionWhere) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := r.DB

	constructTransactionWhereCondition(query, whereCondition)

	err := query.Find(&transactions).Error
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) SummarizeByPeriod(whereCondition models.TransactionWhere) (dto.TransactionSummaryData, error) {
	var summary dto.TransactionSummaryData
	period := whereCondition.Period
	summary.Period = period

	startDate, err := time.Parse("2006-01", period)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return summary, err
	}
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

	err = r.DB.
		Model(&models.Transaction{}).
		Select(`
			? AS period,
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) AS income_amount,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) AS expense_amount
		`, period).
		Where("user_id = ? AND trx_date BETWEEN ? AND ?", whereCondition.UserID, startDate, endDate).
		Scan(&summary).Error
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return summary, err
	}

	return summary, nil
}
