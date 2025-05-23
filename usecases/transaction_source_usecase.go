package usecases

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
)

type TransactionSourceUseCase struct {
	TransactionSourceRepo *repositories.TransactionSourceRepository
}

func NewTransactionSourceUsecase(transactionSourceRepo *repositories.TransactionSourceRepository) *TransactionSourceUseCase {
	return &TransactionSourceUseCase{
		TransactionSourceRepo: transactionSourceRepo,
	}
}

func (uc *TransactionSourceUseCase) GetAllSources(query dto.TransactionSourceQueryParam) ([]models.TransactionSource, error) {
	sources, err := uc.TransactionSourceRepo.GetAll(models.TransactionSourceWhere{
		Type: query.Type,
	})
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return sources, nil
}
