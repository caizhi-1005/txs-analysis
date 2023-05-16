package routers

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/controllers"
)

func init() {
	//初始化nebula (tag:address, edge:transaction)
	beego.Router("/api/insert/init", &controllers.RouteInsert{}, "get:InitNebula")
	//写入交易数据
	beego.Router("/api/insert/txs", &controllers.RouteInsert{}, "get:SyncTxDataToNebula")
	//写入地址数据
	beego.Router("/api/insert/address", &controllers.RouteInsert{}, "get:SyncAddressDataToNebula")

	//提现地址列表
	beego.Router("/api/query/withdrawAccountList", &controllers.RouteQuery{}, "get:WithdrawAccountList")
	//根据给定地址查其提现交易记录
	beego.Router("/api/query/withdrawTxListByAddress", &controllers.RouteQuery{}, "get:WithdrawTxListByAddress")
	//每个账户地址的每笔交易的时间、金额、汇总金额列表
	beego.Router("/api/query/txListGroupByAddress", &controllers.RouteQuery{}, "get:TxListGroupByAddress")
	//nebula:根据指定地址，查询所有入账记录
	beego.Router("/api/query/entryTxsByAddress", &controllers.RouteQuery{}, "get:EntryTxsByAddress")

}
