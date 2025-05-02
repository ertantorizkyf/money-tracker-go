package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	TrxDate    time.Time      `gorm:"type:date" json:"trx_date"`
	Type       string         `gorm:"type:varchar(10)" json:"type"`
	UserID     uint           `json:"user_id"`
	SourceID   uint           `json:"source_id"`
	CategoryID uint           `json:"category_id"`
	Amount     float64        `type:"decimal(15,2)" json:"amount"`
	Purpose    string         `json:"purpose"`
	Remark     string         `gorm:"type:text" json:"remark"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`

	User     User                `json:"user,omitempty"`
	Source   TransactionSource   `gorm:"foreignKey:SourceID" json:"source,omitempty"`
	Category TransactionCategory `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

func (Transaction) TableName() string {
	return "mt_transactions"
}

type TransactionWhere struct {
	UserID     uint
	SourceID   uint
	CategoryID uint
	Purpose    string
	Remark     string
	StartDate  string
	EndDate    string
	Type       string
	Period     string
}
