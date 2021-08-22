package models

import "time"

type TransationDaily struct {
	ID               int64     `gorm:"column:id;primaryKey" json:"id"`
	TransationNum    int64     `gorm:"column:transation_num" json:"transation_num"`
	TransationAmount float64   `gorm:"column:transation_amount" json:"transation_amount"`
	DestroyAmount    float64   `gorm:"column:destroy_amount" json:"destroy_amount"`
	MintAmount       float64   `gorm:"column:mint_amount" json:"mint_amount"`
	BlockNum         int64     `gorm:"column:block_num" json:"block_num"`
	Dt               int64     `gorm:"column:dt" json:"dt"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
}
