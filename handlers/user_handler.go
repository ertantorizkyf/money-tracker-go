package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/ertantorizkyf/money-tracker-go/usecases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	req := dto.RegisterReq{}
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

	isReqValid, message := helpers.ValidateRegisterReq(req)
	if !isReqValid {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			message,
			true,
			nil,
		))
		return
	}

	jwtToken, err := h.UserUseCase.RegisterUser(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := constants.ERR_MESSAGE_INTERNAL_SERVER_ERROR

		if strings.Contains(err.Error(), constants.ERR_MESSAGE_DATA_TAKEN) {
			statusCode = http.StatusConflict
			errMessage = constants.ERR_MESSAGE_DATA_TAKEN
		}

		if strings.Contains(err.Error(), constants.ERR_MESSAGE_BAD_REQUEST) {
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

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully register user",
		false,
		map[string]string{
			"token": *jwtToken,
		},
	))
}

func (h *UserHandler) Login(c *gin.Context) {
	req := dto.LoginReq{}
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

	isReqValid, message := helpers.ValidateLoginReq(req)
	if !isReqValid {
		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			http.StatusBadRequest,
			message,
			true,
			nil,
		))
		return
	}

	jwtToken, err := h.UserUseCase.Login(c, req)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errMessage := constants.ERR_MESSAGE_INTERNAL_SERVER_ERROR
		if errors.Is(err, gorm.ErrRecordNotFound) ||
			strings.Contains(err.Error(), constants.ERR_MESSAGE_INVALID_CREDENTIALS) {
			statusCode = http.StatusUnauthorized
			errMessage = constants.ERR_MESSAGE_INVALID_CREDENTIALS
		}

		c.JSON(http.StatusBadRequest, dto.SetGeneralResp(
			statusCode,
			errMessage,
			true,
			nil,
		))
		return
	}

	c.JSON(http.StatusOK, dto.SetGeneralResp(
		http.StatusOK,
		"Successfully logged in",
		false,
		map[string]string{
			"token": *jwtToken,
		},
	))
}
