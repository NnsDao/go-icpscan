package response

type BlockNewRes struct {
	Oaccountaddress  string
	Tranidentifier string
	Blocktimestamp string
	Transactiohash string
	Ostatus string
	Mmemo string
	Osum string
	Mblockheight string
}

type SearchDetailRes struct {
	Account string // 账户地址
	Type uint32 // 1: Transaction 2:Account
	Status string // 按正常来说应该是枚举，但是数据库存放的是字符串英文
	BlockHeight string // 区块高度
	Timestamp string // 区块时间
	From string // 发起转账地址
	To string // 接收地址
	Amount string // 转账金额
	Fee string // 转账手续费
	Memo string // 区块备注
	Balance string // 账户余额
}
