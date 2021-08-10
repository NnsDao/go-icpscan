package models

import "time"

type Users struct {
	ID          int64     `gorm:"column:id" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	PrincipalId string    `gorm:"column:principal_id" json:"principal_id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
