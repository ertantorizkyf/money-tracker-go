package handlers

import (
	"net/http"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/usecases"
	"github.com/gin-gonic/gin"
)

type TransactionCategoryHandler struct {
	TransactionCategoryUseCase *usecases.TransactionCategoryUseCase
}

func NewTransactionCategoryHandler(transactionCategoryUseCase *usecases.TransactionCategoryUseCase) *TransactionCategoryHandler {
	return &TransactionCategoryHandler{TransactionCategoryUseCase: transactionCategoryUseCase}
}

func (h *TransactionCategoryHandler) GetAllCategories(c *gin.Context) {
	query := dto.TransactionCategoryQueryParam{}
	err := c.BindQuery(&query)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			constants.ERR_MESSAGE_BAD_REQUEST,
			true,
			nil,
		))
		return
	}

	categories, err := h.TransactionCategoryUseCase.GetAllCategories(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.SetGeneralResp(
			http.StatusInternalServerError,
			"Failed to get all categories",
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully get all categories",
		false,
		categories,
	))
}
