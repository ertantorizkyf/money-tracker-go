package middlewares

import (
	"net/http"
	"strings"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/gin-gonic/gin"
)

func RejectAuthorization(c *gin.Context) {
	helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, constants.ERR_MESSAGE_INVALID_CREDENTIALS)
	c.JSON(http.StatusUnauthorized, dto.SetGeneralResp(
		http.StatusUnauthorized,
		constants.ERR_MESSAGE_UNAUTHORIZED,
		true,
		nil,
	))
	c.Abort()
}

func AuthorizeUser(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		RejectAuthorization(c)

		return
	}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) < 2 {
		RejectAuthorization(c)

		return
	}

	claims, err := helpers.VerifyToken(authToken[1])
	if err != nil {
		RejectAuthorization(c)

		return
	}

	subFloat, ok := claims["sub"].(float64)
	if !ok {
		RejectAuthorization(c)
		return
	}

	userID := uint(subFloat)
	c.Set("userID", userID)

	c.Next()
}
