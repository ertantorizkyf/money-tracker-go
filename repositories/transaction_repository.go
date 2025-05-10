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

func constructTransactionWhereCondition(query *gorm.DB, whereCondition models.TransactionWhere) *gorm.DB {
	if whereCondition.UserID > 0 {
		query = query.Where("user_id = ?", whereCondition.UserID)
	}
	if whereCondition.SourceID > 0 {
		query = query.Where("source_id = ?", whereCondition.SourceID)
	}
	if whereCondition.CategoryID > 0 {
		query = query.Where("category_id = ?", whereCondition.CategoryID)
	}
	if whereCondition.Purpose != "" {
		query = query.Where("purpose = ?", whereCondition.Purpose)
	}
	if whereCondition.Remark != "" {
		query = query.Where("remark = ?", whereCondition.Remark)
	}
	if whereCondition.StartDate != "" && whereCondition.EndDate != "" {
		startDate, _ := time.Parse("2006-01-02", whereCondition.StartDate)
		endDate, _ := time.Parse("2006-01-02", whereCondition.EndDate)

		query = query.Where("trx_date BETWEEN ? AND ?", startDate, endDate)
	}
	if whereCondition.Type != "" {
		query = query.Where("type = ?", whereCondition.Type)
	}

	return query
}

func constructTransactionPreload(query *gorm.DB, preload models.TransactionPreload) *gorm.DB {
	if preload.IncludeUser {
		query = query.Preload("User")
	}
	if preload.IncludeCategory {
		query = query.Preload("Category")
	}
	if preload.IncludeSource {
		query = query.Preload("Source")
	}

	return query
}

func constructTransactionOrder(query *gorm.DB, order string) *gorm.DB {
	if order != "" {
		if order == constants.TRANSACTION_ORDER_OLDEST {
			order = "trx_date ASC"
		} else if order == constants.TRANSACTION_ORDER_NEWEST {
			order = "trx_date DESC, created_at DESC"
		}

		query = query.Order(order)
	}

	return query
}

func (r *TransactionRepository) GetAll(
	whereCondition models.TransactionWhere,
	preload models.TransactionPreload,
	order string,
) ([]models.Transaction, error) {
	var transactions []models.Transaction

	query := r.DB

	query = constructTransactionWhereCondition(query, whereCondition)
	query = constructTransactionPreload(query, preload)
	query = constructTransactionOrder(query, order)

	if err := query.Find(&transactions).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepository) GetByUserAndID(userID uint, id uint) (models.Transaction, error) {
	var transaction models.Transaction

	query := r.DB.Where("user_id = ?", userID).Where("id = ?", id)

	if err := query.First(&transaction).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return transaction, err
	}

	return transaction, nil
}

func (r *TransactionRepository) GetSummaryByPeriod(whereCondition models.TransactionWhere) (dto.TransactionSummaryData, error) {
	var summary dto.TransactionSummaryData
	period := whereCondition.Period
	summary.Period = period

	startDate, err := time.Parse("2006-01", period)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return summary, err
	}
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Nanosecond)

	query := r.DB.
		Model(&models.Transaction{}).
		Select(`
			? AS period,
			COALESCE(SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END), 0) AS income_amount,
			COALESCE(SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END), 0) AS expense_amount
		`, period).
		Where("user_id = ? AND trx_date BETWEEN ? AND ?", whereCondition.UserID, startDate, endDate)

	if err := query.Scan(&summary).Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return summary, err
	}

	return summary, nil
}

func (r *TransactionRepository) CreateTransaction(transaction *models.Transaction) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(transaction).Error
	if err != nil {
		tx.Rollback()
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	if err = tx.Commit().Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}

func (r *TransactionRepository) UpdateTransaction(transaction models.Transaction) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Model(&models.Transaction{}).
		Where("id = ?", transaction.ID).
		Updates(&transaction).Error
	if err != nil {
		tx.Rollback()
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	if err = tx.Commit().Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}

func (r *TransactionRepository) DeleteTransaction(trxID uint) error {
	tx := r.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("id = ?", trxID).Delete(&models.Transaction{}).Error
	if err != nil {
		tx.Rollback()
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	if err = tx.Commit().Error; err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}
