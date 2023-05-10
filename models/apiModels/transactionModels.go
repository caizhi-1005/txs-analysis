package blockChain

//获取交易列表--请求结构体
type ReqTransactionList struct {
	UserIDSafe int64 `json:"-"`
}

//获取交易列表--响应结构体
type ResTransactionList struct {
	TradeCounts string      `json:"trade_counts"`
	TradeList   interface{} `json:"trade_list"`
}

