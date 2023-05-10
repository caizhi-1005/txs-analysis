package routers

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/controllers"
)

func init() {
	beego.Router("/txs/insert/nebula", &controllers.RouteInsert{}, "get:InsertDataToNebula")

	beego.Router("/txs/query/withdrawAccountList", &controllers.RouteQuery{}, "get:WithdrawAccountList")
	beego.Router("/txs/query/withdrawTxsList", &controllers.RouteQuery{}, "get:WithdrawTxsList")

}


