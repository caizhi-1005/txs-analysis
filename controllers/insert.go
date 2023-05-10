package controllers

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/service"
)

type RouteInsert struct {
	BaseController
	insertNebulaService service.InsertNebulaService
}

func (this *RouteInsert) InsertDataToNebula() {
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

	errSync := this.insertNebulaService.SyncDataToNebula(startBlock, endBlock)
	if errSync != nil {
		beego.Error("SyncDataToNebula error")
		this.ResponseInfo(500, "SyncDataToNebula error.", errSync)
		return
	}
	this.ResponseInfo(200, "insert data into nebula succeed!", "ok")
}
