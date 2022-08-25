package response

type JSONResult struct {
	Success bool        `json:"success" `
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type BlockNewRes struct {
	Oaccountaddress string
	Tranidentifier  string
	Blocktimestamp  string
	Transactiohash  string
	Ostatus         string
	Mmemo           string
	Osum            string
	Mblockheight    string
}

type SearchDetailRes struct {
	Account        string // 账户地址
	Type           uint32 // 1: Transaction 2:Account
	Status         string // 按正常来说应该是枚举，但是数据库存放的是字符串英文
	BlockHeight    string // 区块高度
	Timestamp      string // 区块时间
	From           string // 发起转账地址
	To             string // 接收地址
	Amount         string // 转账金额
	Fee            string // 转账手续费
	Memo           string // 区块备注
	Balance        string // 账户余额
	Tranidentifier string // 转账hash
	Symbol         string // 转账货币符号
	Decimals       string // 转账货币位数
}

type AccountDealDetail struct {
	Account        string // 账户地址
	Timestamp      string // 区块时间
	From           string // 发起转账地址
	To             string // 接收地址
	Amount         string // 转账金额
	Tranidentifier string // 转账hash
	Fee            string // 转账手续费
}

type AccountDealDetailRes struct {
	Detail   []AccountDealDetail
	Total    int64
	Page     int
	PageSize int
}

type GetAccountBalanceCurveRes struct {
	Date    []string
	Balance []string
}

type DAUAccountRes struct {
	Dt     []int64
	Number []int64
}

type LoginRes struct {
	Name string
}

type GetDayTransationCountRes struct {
	Dt    []int64
	Count []int64
}

type GetDayBlockCountRes struct {
	Dt    []int64
	Count []int64
}

type GetDayTransationAmountRes struct {
	Dt    []int64
	Count []float64
}

type GetDayDestroyAmountRes struct {
	Dt    []int64
	Count []float64
}

type GetDayMintAmountRes struct {
	Dt    []int64
	Count []float64
}

type WalletRelationRes struct {
	Principal string `json:"principal"`
	Address   string `json:"address"`
}
