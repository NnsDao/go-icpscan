package models

import "time"

type Block struct {
	Id uint // 主键
	Blockidentifier string // 当前区块
	Parentblock string // 上一个区块
	Transactiohash string // 转账hash值
	Mblockheight string // 区块metadata高度
	Mmemo string // 区块备注
	Tranidentifier string // 转账标识符
	Mtimestamp string // 区块metadata时间
	Blocktimestamp string // 当前区块时间
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
