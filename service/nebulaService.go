package service

import (
	"github.com/server/txs-analysis/models/dbModels"
)

type NebulaService struct {
}

//根据指定地址，查询所有入账记录
func (this *NebulaService) GetEntryTxsByAddress(string) ([]*dbModels.ResWithdrawAccountList, error) {
	addressList, err := dbModels.GetWithdrawAccountList()
	if err != nil {
		return nil, err
	}
	return addressList, nil
}


