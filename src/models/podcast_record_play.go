package models

import (
	"time"
)

type PodcastRecordPlay struct {
	Id        uint      `gorm:"column:id;type:int(10) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Canister  string    `gorm:"column:canister;type:varchar(100);NOT NULL" json:"canister"`
	PodcastId string    `gorm:"column:podcast_id;type:varchar(100);NOT NULL" json:"podcast_id"`
	CreateAt  time.Time `gorm:"column:create_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_at"`
	UpdateAt  time.Time `gorm:"column:update_at;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_at"`
}

func (m *PodcastRecordPlay) TableName() string {
	return "podcast_record_play"
}
