package handlers

import (
	"net/http"
	"strconv"
	"strings"

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

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	req := dto.CreateTransactionRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			constants.ERR_MESSAGE_BAD_REQUEST,
			true,
			nil,
		))
		return
	}

	isRequestValid, message := helpers.ValidateCreateTransactionRequest(req)
	if !isRequestValid {
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

	transaction, err := h.TransactionUseCase.CreateTransaction(userID.(uint), req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "Failed to create transaction"
		if strings.Contains(err.Error(), constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE) {
			statusCode = http.StatusBadRequest
			errMessage = constants.ERR_MESSAGE_BAD_REQUEST
		}

		c.JSON(statusCode, dto.SetGeneralResp(
			statusCode,
			errMessage,
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusCreated, dto.SetGeneralResp(
		http.StatusCreated,
		"Successfully created transaction",
		false,
		*transaction,
	))
}

func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			constants.ERR_MESSAGE_BAD_REQUEST,
			true,
			nil,
		))
		return
	}

	req := dto.UpdateTransactionRequest{}
	err = c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			constants.ERR_MESSAGE_BAD_REQUEST,
			true,
			nil,
		))
		return
	}

	isRequestValid, message := helpers.ValidateUpdateTransactionRequest(req)
	if !isRequestValid {
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

	transaction, err := h.TransactionUseCase.UpdateTransaction(userID.(uint), uint(id), req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "Failed to update transaction"
		if strings.Contains(err.Error(), constants.ERR_MESSAGE_INVALID_TRANSACTION_TYPE) {
			statusCode = http.StatusBadRequest
			errMessage = constants.ERR_MESSAGE_BAD_REQUEST
		}
		if strings.Contains(err.Error(), constants.ERR_MESSAGE_RECORD_NOT_FOUND) {
			statusCode = http.StatusNotFound
			errMessage = constants.ERR_MESSAGE_RECORD_NOT_FOUND
		}

		c.JSON(statusCode, dto.SetGeneralResp(
			statusCode,
			errMessage,
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully updated transaction",
		false,
		*transaction,
	))
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			constants.ERR_MESSAGE_BAD_REQUEST,
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

	if err := h.TransactionUseCase.DeleteTransaction(userID.(uint), uint(id)); err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := "Failed to update transaction"
		if strings.Contains(err.Error(), constants.ERR_MESSAGE_RECORD_NOT_FOUND) {
			statusCode = http.StatusNotFound
			errMessage = constants.ERR_MESSAGE_RECORD_NOT_FOUND
		}

		c.JSON(statusCode, dto.SetGeneralResp(
			statusCode,
			errMessage,
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusNoContent, dto.SetGeneralResp(
		http.StatusNoContent,
		"Successfully deleted transaction",
		false,
		nil,
	))
}
