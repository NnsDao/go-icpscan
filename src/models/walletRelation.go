package models

import (
	"time"
)

type WalletRelation struct {
	Id        uint      `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Principal string    `gorm:"column:principal;type:varchar(64);comment:principal;NOT NULL" json:"principal"`
	Address   string    `gorm:"column:address;type:varchar(64);NOT NULL" json:"address"`
	CreateAt  time.Time `gorm:"column:create_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"`
}

func (m *WalletRelation) TableName() string {
	return "wallet_relation"
}
