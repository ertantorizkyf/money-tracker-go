package usecases

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/models"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
)

type TransactionCategoryUseCase struct {
	TransactionCategoryRepo *repositories.TransactionCategoryRepository
}

func NewTransactionCategoryUsecase(transactionCategoryRepo *repositories.TransactionCategoryRepository) *TransactionCategoryUseCase {
	return &TransactionCategoryUseCase{
		TransactionCategoryRepo: transactionCategoryRepo,
	}
}

func (uc *TransactionCategoryUseCase) GetAllCategories() ([]models.TransactionCategory, error) {
	categories, err := uc.TransactionCategoryRepo.GetAll()
	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, err)
		return nil, err
	}

	return categories, nil
}
