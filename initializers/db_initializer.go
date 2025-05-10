package initializers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ertantorizkyf/money-tracker-go/constants"
	"github.com/ertantorizkyf/money-tracker-go/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// DB SAMPLE IS USING MYSQL DRIVER
	var err error
	var dsn string

	// GET DB CONFIG FROM ENV VARS
	dbUsingPass, _ := strconv.ParseBool(os.Getenv("DB_USING_PASS"))
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// CONFIGURE DB CONN
	if dbUsingPass {
		dsn = fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	} else {
		dsn = fmt.Sprintf("%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbHost, dbName)
	}

	// INIT DB CONN
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_ERROR, "[ERR] Failed to connect to DB")
	} else {
		helpers.LogWithSeverity(constants.LOGGER_SEVERITY_INFO, "[INFO] Connected to DB")
	}
}
