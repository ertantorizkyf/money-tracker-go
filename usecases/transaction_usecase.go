package usecases

import (
	"fmt"
	"time"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
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
	})
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

func (uc *TransactionUseCase) CreateTransaction(userID uint, req dto.CreateTransactionRequest) error {
	if req.TrxDate == "" {
		req.TrxDate = time.Now().Format("2006-01-02")
	}

	category, err := uc.TransactionCategoryRepo.GetByID(req.CategoryID)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	source, err := uc.TransactionSourceRepo.GetByID(req.SourceID)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	if req.Type != category.Type || req.Type != source.Type || category.Type != source.Type {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE)
		return fmt.Errorf("an error has occurred: %s", constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE)
	}

	trxDate, err := time.Parse("2006-01-02", req.TrxDate)
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
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
	if err := uc.TransactionRepo.CreateTransaction(transaction); err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return err
	}

	return nil
}
