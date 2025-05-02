package usecases

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
)

type TransactionUseCase struct {
	TransactionRepo *repositories.TransactionRepository
}

func NewTransactionUsecase(transactionRepo *repositories.TransactionRepository) *TransactionUseCase {
	return &TransactionUseCase{
		TransactionRepo: transactionRepo,
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
