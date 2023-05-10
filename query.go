package controllers

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/service"
)

type RouteQuery struct {
	BaseController
	addressService service.AddressService
	txService      service.TxService
}

func (this *RouteQuery) WithdrawAccountList() {
	result, err := this.addressService.WithdrawAccountList()
	if err != nil {
		beego.Error("get WithdrawAccountList error")
		this.ResponseInfo(500, "get WithdrawAccountList error.", err)
		return
	}
	this.ResponseInfo(200, nil, result)
}

func (this *RouteQuery) WithdrawTxsList() {
	address := this.GetString("address")
	if len(address) > 0 && len(address) != 42 {
		beego.Error("input address error. address:", address)
		this.ResponseInfo(500, "input param address error.", nil)
		return
	}
	result, err := this.txService.WithdrawTxListByAddress(address)
	if err != nil {
		beego.Error("get WithdrawAccountList error")
		this.ResponseInfo(500, "get WithdrawAccountList error.", err)
		return
	}
	this.ResponseInfo(200, nil, result)
}
