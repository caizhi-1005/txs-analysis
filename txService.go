package service

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/models/dbModels"
)

type TxService struct {
}

//查询当前账户地址的全部提现交易
func (this *TxService) WithdrawTxListByAddress(address string) ([]*dbModels.TbTransaction, error) {
	//inputData := "0x86d1a69f"
	//contract := "0x871fcb6b836db1b5d6ee64901fb17245cd403e6d"
	inputData := beego.AppConfig.String("withdraw::inputdata")
	contract := beego.AppConfig.String("withdraw::contract")
	addressList, err := dbModels.WithdrawTxsListByAddress(address, inputData, contract)
	if err != nil {
		return nil, err
	}
	return addressList, nil
}

func (this *TxService) TxListByAddress(address string) ([]dbModels.ResTxListByAddress, error) {
	if len(address) == 0 {

	}
	txListAll := []dbModels.ResTxListByAddress{}
	addressList, err := dbModels.AccountInfoList()
	//根据地址获取交易列表
	for _, v := range addressList {
		txListByAddress := dbModels.ResTxListByAddress{}
		txList, err := dbModels.TransactionListGetByAddress(v.AccountAddress)
		if err != nil {
			beego.Error("dbModels.TransactionListGetByAddress error. account_address:", v.AccountAddress)
		}
		txListByAddress.TxList = txList
		txListAll = append(txListAll, txListByAddress)
	}

	if err != nil {
		return nil, err
	}
	return txListAll, nil
}
