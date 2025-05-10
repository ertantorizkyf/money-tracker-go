package handlers

import (
	"net/http"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/usecases"
	"github.com/gin-gonic/gin"
)

type TransactionSourceHandler struct {
	TransactionSourceUseCase *usecases.TransactionSourceUseCase
}

func NewTransactionSourceHandler(transactionSourceUseCase *usecases.TransactionSourceUseCase) *TransactionSourceHandler {
	return &TransactionSourceHandler{TransactionSourceUseCase: transactionSourceUseCase}
}

func (h *TransactionSourceHandler) GetAllSources(c *gin.Context) {
	query := dto.TransactionSourceQueryParam{}
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

	isTrxTypeValid, message := helpers.ValidateTrxType(query.Type)
	if !isTrxTypeValid {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			message,
			true,
			nil,
		))
		return
	}

	sources, err := h.TransactionSourceUseCase.GetAllSources(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.SetGeneralResp(
			http.StatusInternalServerError,
			"Failed to get all sources",
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully get all sources",
		false,
		sources,
	))
}
