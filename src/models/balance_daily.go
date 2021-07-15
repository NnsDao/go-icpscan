package models

import "time"

type BalanceDaily struct {
	ID int64 `gorm:"column:id;primaryKey" json:"id"`
	Oaccountaddress string `gorm:"column:oaccountaddress" json:"oaccountaddress"`
	Balance string `gorm:"column:balance" json:"balance"`
	Dt int64 `gorm:"column:dt" json:"dt"`
	CreatedAt time.Time `gorm:"-" json:"created_at"`
	UpdatedAt time.Time `gorm:"-" json:"updated_at"`
}
