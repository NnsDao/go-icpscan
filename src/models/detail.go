package models

import "time"

type Detail struct {
	Id uint // 主键
	Tranidentifier string // 转账hash值
	Oindex string // 当前这笔转账的索引
	Otype string // 转账类型
	Ostatus string // 转账状态
	Oaccountaddress string // 转账账户地址
	Oamountvalue string // 转账价值，可能为负数
	Oamountcurrencydecimals string //转账代币小数点
	Oamountcurrencysymbol string
	Mblockheight string // 区块metadata高度
	Mmemo string // 区块备注
	Mtimestamp string // 区块备注
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
}
