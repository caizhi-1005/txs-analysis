package service

import (
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/models/dbModels"
)

type TxService struct {
}

//查询当前账户地址的全部提现交易
func (this *TxService) WithdrawTxListByAddress(address string) ([]*dbModels.TbTransaction, error) {
	addressList, err := dbModels.WithdrawTxsListByAddress(address)
	if err != nil {
		return nil, err
	}
	return addressList, nil
}

//获取每个账户地址每笔交易的时间、金额、汇总金额列表
func (this *TxService) TxListGroupByAddress(address string) ([]dbModels.ResTxListByAddress, error) {
	txListAll := []dbModels.ResTxListByAddress{}
	if len(address) == 0 {
		addressList, err := dbModels.AccountInfoList()
		if err != nil {
			return nil, err
		}
		//根据地址获取交易列表
		for _, v := range addressList {
			txListByAddress := dbModels.ResTxListByAddress{}
			txList, err := dbModels.TransactionListGetByAddress(v.AccountAddress)
			if err != nil {
				beego.Error("dbModels.TransactionListGetByAddress error. account_address:", v.AccountAddress)
				return nil, err
			}
			txListByAddress.TxList = txList

			//获取入账总金额
			inTotalValue, err := dbModels.GetInTotalValueByAddress(v.AccountAddress)
			txListByAddress.InTotalValue = inTotalValue

			//获取出账总金额
			outTotalValue, err := dbModels.GetOutTotalValueByAddress(v.AccountAddress)
			txListByAddress.OutTotalValue = outTotalValue
			txListAll = append(txListAll, txListByAddress)
		}

	}else {
		txListByAddress := dbModels.ResTxListByAddress{}
		txList, err := dbModels.TransactionListGetByAddress(address)
		if err != nil {
			beego.Error("dbModels.TransactionListGetByAddress error. account_address:", address)
			return nil, err
		}

		txListByAddress.TxList = txList

		//获取入账总金额
		inTotalValue, err := dbModels.GetInTotalValueByAddress(address)
		txListByAddress.InTotalValue = inTotalValue

		//获取出账总金额
		outTotalValue, err := dbModels.GetOutTotalValueByAddress(address)
		txListByAddress.OutTotalValue = outTotalValue
		txListAll = append(txListAll, txListByAddress)
	}

	return txListAll, nil
}
