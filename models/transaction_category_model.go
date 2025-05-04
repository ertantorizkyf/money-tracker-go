package models

type TransactionCategory struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Type   string `gorm:"type:varchar(10)" json:"type"`
	Name   string `json:"name"`
	Remark string `gorm:"type:text" json:"remark"`
}

func (TransactionCategory) TableName() string {
	return "mt_transaction_categories"
}

type TransactionCategoryWhere struct {
	Type string
}
