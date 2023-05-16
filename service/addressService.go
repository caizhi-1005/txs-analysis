package service

import (
	"github.com/server/txs-analysis/models/dbModels"
)

type AddressService struct {
}

//获取所有提现账户地址列表
func (this *AddressService) WithdrawAccountList() ([]*dbModels.ResWithdrawAccountList, error) {
	addressList, err := dbModels.GetWithdrawAccountList()
	if err != nil {
		return nil, err
	}
	return addressList, nil
}

