package handlers

import (
	"net/http"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/usecases"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionUseCase *usecases.TransactionUseCase
}

func NewTransactionHandler(transactionUseCase *usecases.TransactionUseCase) *TransactionHandler {
	return &TransactionHandler{TransactionUseCase: transactionUseCase}
}

func (h *TransactionHandler) GetAllTransactions(c *gin.Context) {
	query := dto.TransactionQueryParam{}
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

	isQueryParamValid, message := helpers.ValidateTransactionQueryParam(query)
	if !isQueryParamValid {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			message,
			true,
			nil,
		))
		return
	}

	// GET USER ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.SetGeneralResp(
			http.StatusUnauthorized,
			constants.ERR_MESSAGE_UNAUTHORIZED,
			true,
			nil,
		))
		return
	}

	transactions, err := h.TransactionUseCase.GetAllTransactions(userID.(uint), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.SetGeneralResp(
			http.StatusInternalServerError,
			"Failed to get all transactions",
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully get all transactions",
		false,
		transactions,
	))
}

func (h *TransactionHandler) GetTransactionSummary(c *gin.Context) {
	query := dto.TransactionSummaryQueryParam{}
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

	isQueryParamValid, message := helpers.ValidateTransactionSummaryQueryParam(query)
	if !isQueryParamValid {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			message,
			true,
			nil,
		))
		return
	}

	// GET USER ID
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, dto.SetGeneralResp(
			http.StatusUnauthorized,
			constants.ERR_MESSAGE_UNAUTHORIZED,
			true,
			nil,
		))
		return
	}

	transactions, err := h.TransactionUseCase.GetTransactionSummary(userID.(uint), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.SetGeneralResp(
			http.StatusInternalServerError,
			"Failed to get transaction summary",
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully get transaction summary",
		false,
		transactions,
	))
}
