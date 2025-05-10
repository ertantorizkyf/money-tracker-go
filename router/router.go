package router

import (
	"fmt"
	"net/http"

	"github.com/ertantorizkyf/money-tracker-go/dto"
	"github.com/ertantorizkyf/money-tracker-go/handlers"
	"github.com/ertantorizkyf/money-tracker-go/repositories"
	"github.com/ertantorizkyf/money-tracker-go/usecases"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// SETUP REPOS
	userRepo := repositories.NewUserRepository()
	transactionRepo := repositories.NewTransactionRepository()
	transactionSourceRepo := repositories.NewTransactionSourceRepository()
	transactionCategoryRepo := repositories.NewTransactionCategoryRepository()
	transactionRedisRepo := repositories.NewTransactionRedisRepository()

	// SETUP USECASES
	userUseCase := usecases.NewUserUsecase(userRepo)
	transactionUseCase := usecases.NewTransactionUsecase(
		transactionRepo,
		transactionCategoryRepo,
		transactionSourceRepo,
		transactionRedisRepo,
	)
	transactionSourceUseCase := usecases.NewTransactionSourceUsecase(transactionSourceRepo)
	transactionCategoryUseCase := usecases.NewTransactionCategoryUsecase(transactionCategoryRepo)

	// SETUP HANDLERS
	userHandler := handlers.NewUserHandler(userUseCase)
	transactionHandler := handlers.NewTransactionHandler(transactionUseCase)
	transactionSourceHandler := handlers.NewTransactionSourceHandler(transactionSourceUseCase)
	transactionCategoryHandler := handlers.NewTransactionCategoryHandler(transactionCategoryUseCase)

	// DEFINE API
	apiGroup := router.Group("/api")

	// REGISTER ROUTES
	userRoutes(apiGroup, userHandler)
	transactionRoutes(
		apiGroup,
		transactionHandler,
		transactionSourceHandler,
		transactionCategoryHandler,
	)

	router.GET("/ping", func(c *gin.Context) {
		response := dto.SetGeneralResp(
			http.StatusOK,
			"Pong",
			true,
			nil,
		)
		c.JSON(http.StatusOK, response)
	})

	router.NoRoute(func(c *gin.Context) {
		errMethod := c.Request.Method
		errPath := c.Request.URL.Path
		errMessage := fmt.Sprintf("Path [%s] %s not found!", errMethod, errPath)

		response := dto.SetGeneralResp(
			http.StatusNotFound,
			errMessage,
			true,
			nil,
		)
		c.JSON(response.StatusCode, response)
	})

	return router
}
