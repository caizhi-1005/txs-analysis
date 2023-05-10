package service

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/server/txs-analysis/models/dbModels"
	"github.com/server/txs-analysis/models/nebulaModels"
)

type InsertNebulaService struct {
}

//获取通证详情信息
func (this *InsertNebulaService) SyncDataToNebula(startNum, endNum int) error {
	ormer := orm.NewOrm()
	nebulaDB := nebulaModels.Init()
	for currentBlockId := startNum; currentBlockId <= endNum; currentBlockId += 1001 {
		currentEndBlockId := currentBlockId + 1000
		if currentEndBlockId > endNum {
			currentEndBlockId = endNum
		}

		//查询调用提现合约的交易 (to=0x871fcb6b836db1b5d6ee64901fb17245cd403e6d or input_data=0x86d1a69f)
		//txs1 := []dbModels.TbTransaction{}
		//sq1 := "select * from tb_transaction where tx_status = 1 and input_data='0x86d1a69f' or `to`= '0x871fcb6b836db1b5d6ee64901fb17245cd403e6d';"
		//ormer.Raw(sq1).QueryRows(&txs1)

		txs1, err := dbModels.GetWithdrawTxList()
		if err != nil {
			beego.Error("dbModels.GetWithdrawTxList error: ", err)
			return err
		}
		for _, v := range txs1 {
			tx := &nebulaModels.NebulaTransaction{
				TxHash:          v.TxHash,
				TxTime:          v.TxTime,
				ContractAddress: v.ContractAddress,
				FromAddress:     v.From,
				ToAddress:       v.To,
				Amount:          v.Amount,
			}
			err3 := nebulaModels.InsertTxn(nebulaDB, tx)
			if err3 != nil {
				beego.Error("nebulaModels.InsertTxn error: ", err3)
				return err3
			}

			//根据调用合约的账户地址，查询账户类型
			var accountType int
			sq2 := "select account_type from tb_account_info where account_address = '" + v.From + "'"
			ormer.Raw(sq2).QueryRow(&accountType)
			address := &nebulaModels.Address{
				Address: v.From,
				Type:    accountType,
			}
			err1 := nebulaModels.InsertAddress(nebulaDB, address)
			if err1 != nil {
				beego.Error("nebulaModels.InsertAddress error: ", err1)
				return err1
			}

			//根据调用合约的账户地址，查询当前账户发起的交易
			txs2 := []dbModels.TbTransaction{}
			sq3 := "select * from tb_transaction where `from` = '" + v.From + "'"
			ormer.Raw(sq3).QueryRows(&txs2)
		}
	}
	nebulaDB.Close()
	return nil
}
