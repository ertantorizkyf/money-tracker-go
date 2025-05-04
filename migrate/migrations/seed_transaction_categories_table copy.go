package migrations

import "gorm.io/gorm"

func SeedTransactionCategoriesTable(db *gorm.DB) error {
	return db.Exec(`
		INSERT INTO mt_transaction_categories(type, name, remark)
		VALUES
			("income", "Income", "Any type of income"),
			("expense", "Debts", "e.g. cc payment, installment plan"),
			("expense", "Utilities", "e.g. electric, gas, water, internet bills"),
			("expense", "Transports", "e.g. Gopay, Flazz card topup"),
			("expense", "Investments", "e.g. Mutual fund, ETF, Stocks"),
			("expense", "Beauty & Personal Care", "e.g. Skincare, Makeup"),
			("expense", "Family Support", "e.g. Helping parent's bill"),
			("expense", "Rent", "e.g. Apartment room payment"),
			("expense", "Hobbies", "e.g. Books, Video Games, Plushies"),
			("expense", "Food & Dining", "e.g. Dine out, convenience store snacks, coffees, or groceries"),
			("expense", "Lifestyle", "e.g. Movie tickets, theme park tickets, hotel bookings"),
			("expense", "Household", "e.g. Detergent, toilet cleaner"),
			("expense", "Electronics", "e.g. Phone and laptop"),
			("expense", "Health & Wellness", "e.g. Doctor appointment, prescription, and supplements"),
			("expense", "Gifts & Donations", "e.g. Wedding gifts"),
			("expense", "Fashion", "e.g. Clothes");
	`).Error
}
