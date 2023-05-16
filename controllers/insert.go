package controllers

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/models/nebulaModels"
	"github.com/server/txs-analysis/service"
)

type RouteInsert struct {
	BaseController
	insertNebulaService service.InsertNebulaService
}

func (this *RouteInsert) InitNebula() {
	beego.Debug("InitNebula start------------>")
	nebulaDB := nebulaModels.Init()
	//准备tag
	err := nebulaModels.PrepareAddress(nebulaDB)
	if err != nil {
		beego.Error("nebulaModels.PrepareAddress error: ", err)
		//this.ResponseInfo(500, "PrepareAddress error.", err)
		//return
	}
	//准备edge
	err = nebulaModels.PrepareTxs(nebulaDB)
	if err != nil {
		beego.Error("nebulaModels.PrepareTxs error: ", err)
		this.ResponseInfo(500, "PrepareTxs error.", err)
		return
	}
	nebulaDB.Close()
	beego.Info("nebula prepare success")
	this.ResponseInfo(200, "nebula prepare success succeed!", "ok")
}

func (this *RouteInsert) SyncTxDataToNebula() {
	startBlock, err := this.GetInt("startBlock")
	if err != nil {
		beego.Error("input startBlock error. startBlock:", startBlock)
		this.ResponseInfo(500, "input param startBlock error.", nil)
		return
	}
	endBlock, err := this.GetInt("endBlock")
	if err != nil {
		beego.Error("input endBlock error. endBlock:", endBlock)
		this.ResponseInfo(500, "input param endBlock error.", nil)
		return
	}

	beego.Debug("SyncTxDataToNebula start------------>")
	errSync := this.insertNebulaService.SyncTxDataToNebula(startBlock, endBlock)
	if errSync != nil {
		beego.Error("SyncTxDataToNebula error")
		this.ResponseInfo(500, "SyncTxDataToNebula error.", errSync)
		return
	}
	this.ResponseInfo(200, "SyncTxDataToNebula succeed!", "ok")
}

func (this *RouteInsert) SyncAddressDataToNebula() {

	beego.Debug("SyncAddressDataToNebula start------------>")
	errSync := this.insertNebulaService.SyncAddressDataToNebula()
	if errSync != nil {
		beego.Error("SyncAddressDataToNebula error")
		this.ResponseInfo(500, "SyncAddressDataToNebula error.", errSync)
		return
	}
	this.ResponseInfo(200, "SyncAddressDataToNebula succeed!", "ok")
}
