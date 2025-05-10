package initializers

import (
	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	// LOAD .ENV FILE
	err := godotenv.Load()

	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Error loading .env file")
	} else {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_INFO, "[INFO] .env file loaded")
	}
}
