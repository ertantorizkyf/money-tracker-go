package migrations

import "gorm.io/gorm"

func CreateTransactionCategoriesTable(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS mt_transaction_categories (
			id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			type VARCHAR(10) NOT NULL,
			name VARCHAR(255) NOT NULL,
			remark TEXT,
			PRIMARY KEY (id),
			UNIQUE KEY uq_transaction_categories_type_name (type, name)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
	`).Error
}
