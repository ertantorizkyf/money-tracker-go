package migrations

import "gorm.io/gorm"

func SeedTransactionSourcesTable(db *gorm.DB) error {
	return db.Exec(`
		INSERT INTO mt_transaction_sources(type, name, remark)
		VALUES
			("income", "Salary", "Salary of your full time/part time/freelance job"),
			("income", "Stock Dividend", "Stock dividend payout"),
			("expense", "Bank Account", "Your bank account"),
			("expense", "Credit Card", "Your credit card"),
			("expense", "Loans", "Any type of loans");
	`).Error
}
