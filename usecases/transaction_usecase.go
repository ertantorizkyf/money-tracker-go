package usecases

import (
	"fmt"
	"time"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
	"gorm.io/gorm"
)

type TransactionUseCase struct {
	TransactionRepo         *repositories.TransactionRepository
	TransactionCategoryRepo *repositories.TransactionCategoryRepository
	TransactionSourceRepo   *repositories.TransactionSourceRepository
}

func NewTransactionUsecase(
	transactionRepo *repositories.TransactionRepository,
	transactionCategoryRepo *repositories.TransactionCategoryRepository,
	transactionSourceRepo *repositories.TransactionSourceRepository,
) *TransactionUseCase {
	return &TransactionUseCase{
		TransactionRepo:         transactionRepo,
		TransactionCategoryRepo: transactionCategoryRepo,
		TransactionSourceRepo:   transactionSourceRepo,
	}
}

func (uc *TransactionUseCase) GetAllTransactions(userID uint, query dto.TransactionQueryParam) ([]models.Transaction, error) {
	transactions, err := uc.TransactionRepo.GetAll(models.TransactionWhere{
		UserID:     userID,
		SourceID:   query.SourceID,
		CategoryID: query.CategoryID,
		Purpose:    query.Purpose,
		Remark:     query.Remark,
		StartDate:  query.StartDate,
		EndDate:    query.EndDate,
		Type:       query.Type,
	}, models.TransactionPreload{
		IncludeSource:   true,
		IncludeCategory: true,
	}, query.Order)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return transactions, nil
}

func (uc *TransactionUseCase) GetTransactionSummary(userID uint, query dto.TransactionSummaryQueryParam) (dto.TransactionSummaryData, error) {
	if query.Period == "" {
		query.Period = time.Now().Format("2006-01")
	}

	summary, err := uc.TransactionRepo.SummarizeByPeriod(models.TransactionWhere{
		UserID: userID,
		Period: query.Period,
	})
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return summary, err
	}

	return summary, nil
}

func (uc *TransactionUseCase) CreateTransaction(userID uint, req dto.CreateTransactionRequest) (*models.Transaction, error) {
	if req.TrxDate == "" {
		req.TrxDate = time.Now().Format("2006-01-02")
	}

	// GET CATEGORY
	category, err := uc.TransactionCategoryRepo.GetByID(req.CategoryID)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	// GET SOURCE
	source, err := uc.TransactionSourceRepo.GetByID(req.SourceID)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	// VALIDATE REQ TYPE
	if req.Type != category.Type || req.Type != source.Type || category.Type != source.Type {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE)
		return nil, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE)
	}

	// CONSTRUCT TRANSACTION
	trxDate, err := time.Parse("2006-01-02", req.TrxDate)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	transaction := models.Transaction{
		TrxDate:    trxDate,
		Type:       req.Type,
		UserID:     userID,
		SourceID:   req.SourceID,
		CategoryID: req.CategoryID,
		Amount:     req.Amount,
		Purpose:    req.Purpose,
		Remark:     req.Remark,
	}

	// CREATE TRANSACTION
	if err := uc.TransactionRepo.CreateTransaction(transaction); err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return &transaction, nil
}

func (uc *TransactionUseCase) UpdateTransaction(userID uint, trxID uint, req dto.UpdateTransactionRequest) (*models.Transaction, error) {
	if req.TrxDate == "" {
		req.TrxDate = time.Now().Format("2006-01-02")
	}

	// GET TRANSACTION BY ID
	transaction, err := uc.TransactionRepo.GetByID(trxID)
	if err != nil && err != gorm.ErrRecordNotFound {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}
	if (err != nil && err == gorm.ErrRecordNotFound) || transaction.UserID != userID {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_RECORD_NOT_FOUND)
		return nil, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_RECORD_NOT_FOUND)
	}

	// GET CATEGORY
	category, err := uc.TransactionCategoryRepo.GetByID(req.CategoryID)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	// GET SOURCE
	source, err := uc.TransactionSourceRepo.GetByID(req.SourceID)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	// VALIDATE REQ TYPE
	if req.Type != category.Type || req.Type != source.Type || category.Type != source.Type {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE)
		return nil, fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE)
	}

	// CONSTRUCT TRANSACTION
	trxDate, err := time.Parse("2006-01-02", req.TrxDate)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	transaction.TrxDate = trxDate
	transaction.Type = req.Type
	transaction.SourceID = req.SourceID
	transaction.CategoryID = req.CategoryID
	transaction.Amount = req.Amount
	transaction.Purpose = req.Purpose
	transaction.Remark = req.Remark

	// UPDATE TRANSACTION
	if err := uc.TransactionRepo.UpdateTransaction(transaction); err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return &transaction, nil
}

func (uc *TransactionUseCase) DeleteTransaction(userID uint, trxID uint) error {
	// GET TRANSACTION BY ID
	transaction, err := uc.TransactionRepo.GetByID(trxID)
	if err != nil && err != gorm.ErrRecordNotFound {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}
	if (err != nil && err == gorm.ErrRecordNotFound) || transaction.UserID != userID {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_RECORD_NOT_FOUND)
		return fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_RECORD_NOT_FOUND)
	}

	// DELETE TRANSACTION
	if err := uc.TransactionRepo.DeleteTransaction(trxID); err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}
