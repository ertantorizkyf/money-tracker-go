package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	Username  string         `gorm:"uniqueIndex" json:"username"`
	Email     string         `gorm:"uniqueIndex" json:"email"`
	Phone     string         `gorm:"uniqueIndex" json:"phone"`
	DOB       time.Time      `gorm:"type:date" json:"dob"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions,omitempty"`
}

func (User) TableName() string {
	return "mt_users"
}

type UserWhere struct {
	Username string
	Email    string
	Phone    string
}
