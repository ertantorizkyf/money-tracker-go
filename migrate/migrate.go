package main

import (
	"github.com/ertantorizkyf/money-tracker-go/initializers"
	"github.com/ertantorizkyf/money-tracker-go/migrate/migrations"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectDB()
}

func main() {
	// MANUAL MIGRATE TABLE
	migrations.CreateUsersTable(initializers.DB)
	migrations.CreateTransactionCategoriesTable(initializers.DB)
	migrations.CreateTransactionSourcesTable(initializers.DB)
	migrations.CreateTransactionsTable(initializers.DB)
}
