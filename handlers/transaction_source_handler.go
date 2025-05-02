package handlers

import (
	"net/http"

	"github.com/ertantorizkyf/money-tracker-go/dto"
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
	sources, err := h.TransactionSourceUseCase.GetAllSources()
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
