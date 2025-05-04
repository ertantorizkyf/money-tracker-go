package migrations

import "gorm.io/gorm"

func CreateTransactionsTable(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS mt_transactions (
			id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			trx_date DATE NOT NULL,
			type VARCHAR(10) NOT NULL,
			user_id BIGINT UNSIGNED NOT NULL,
			source_id BIGINT UNSIGNED NOT NULL,
			category_id BIGINT UNSIGNED NOT NULL,
			amount DECIMAL(15,2) NOT NULL,
			purpose VARCHAR(255) NOT NULL,
			remark TEXT,
			created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			deleted_at DATETIME NULL,
			PRIMARY KEY (id),
			INDEX idx_transactions_deleted_at (deleted_at)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
	`).Error
}
