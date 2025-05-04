package models

type TransactionSource struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Type   string `gorm:"type:varchar(10)" json:"type"`
	Name   string `json:"name"`
	Remark string `gorm:"type:text" json:"remark"`
}

func (TransactionSource) TableName() string {
	return "mt_transaction_sources"
}

type TransactionSourceWhere struct {
	ID   uint
	Type string
}
