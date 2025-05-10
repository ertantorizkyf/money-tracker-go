package router

import (
	"github.com/ertantorizkyf/money-tracker-go/handlers"
	"github.com/ertantorizkyf/money-tracker-go/middlewares"
	"github.com/gin-gonic/gin"
)

func transactionRoutes(
	apiGroup *gin.RouterGroup,
	transactionHandler *handlers.TransactionHandler,
	transactionSourceHandler *handlers.TransactionSourceHandler,
	transactionCategoryHandler *handlers.TransactionCategoryHandler,
) {
	transactionGroup := apiGroup.Group("/transactions", middlewares.AuthorizeUser)
	{
		// TRX SOURCE GROUP
		transactionSourceGroup := transactionGroup.Group("/sources")
		{
			transactionSourceGroup.GET("/", transactionSourceHandler.GetAllSources)
		}

		// TRX CATEGORY GROUP
		transactionCategoryGroup := transactionGroup.Group("/categories")
		{
			transactionCategoryGroup.GET("/", transactionCategoryHandler.GetAllCategories)
		}

		// TRX GROUP
		transactionGroup.GET("/", transactionHandler.GetAllTransactions)
		transactionGroup.POST("/", transactionHandler.CreateTransaction)
		transactionGroup.PATCH("/:id", transactionHandler.UpdateTransaction)
		transactionGroup.DELETE("/:id", transactionHandler.DeleteTransaction)
		transactionGroup.GET("/summary", transactionHandler.GetTransactionSummary)
	}
}
