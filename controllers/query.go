package controllers

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/service"
)

type RouteQuery struct {
	BaseController
	addressService service.AddressService
	txService      service.TxService
	nebulaService      service.NebulaService
}

//提现账户地址列表
func (this *RouteQuery) WithdrawAccountList() {
	result, err := this.addressService.WithdrawAccountList()
	if err != nil {
		beego.Error("get WithdrawAccountList error")
		this.ResponseInfo(500, "get WithdrawAccountList error.", err)
		return
	}
	this.ResponseInfo(200, nil, result)
}

//根据账户地址查询其全部提现交易
func (this *RouteQuery) WithdrawTxListByAddress() {
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

//获取每个账户地址每笔交易的时间、金额、汇总金额列表
func (this *RouteQuery) TxListGroupByAddress() {
	address := this.GetString("address")
	if len(address) > 0 && len(address) != 42 {
		beego.Error("input address error. address:", address)
		this.ResponseInfo(500, "input param address error.", nil)
		return
	}
	result, err := this.txService.TxListGroupByAddress(address)
	if err != nil {
		beego.Error("get WithdrawAccountList error")
		this.ResponseInfo(500, "get WithdrawAccountList error.", err)
		return
	}
	this.ResponseInfo(200, nil, result)
}

//nebula:根据指定地址，查询所有入账记录
func (this *RouteQuery) EntryTxsByAddress() {
	address := this.GetString("address")
	if len(address) != 42 {
		beego.Error("input address error. address:", address)
		this.ResponseInfo(500, "input param address error.", nil)
		return
	}
	result, err := this.nebulaService.GetEntryTxsByAddress(address)
	if err != nil {
		beego.Error("GetEntryTxsByAddress error.", err)
		this.ResponseInfo(500, "GetEntryTxsByAddress error.", err)
		return
	}
	this.ResponseInfo(200, nil, result)
}
