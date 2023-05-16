package nebulaModels

import (
	"github.com/zhihu/norm"
)

type Address struct {
	norm.VModel
	Address string `norm:"address"`
	Type    int    `norm:"type"`
}

type NebulaTransaction struct {
	norm.EModel
	TxHash          string `norm:"tx_hash"`
	TxTime          string `norm:"tx_time"`
	ContractAddress string `norm:"contract_address"`
	FromAddress     string `norm:"from_address"`
	ToAddress       string `norm:"to_address"`
	Amount          string `norm:"amount"`
}

//type TransactionEdge struct {
//	TxHash          string          `norm:"tx_hash"`
//	TxTime          nebula.DateTime `norm:"tx_time"`
//	ContractAddress string          `norm:"contract_address"`
//	FromAddress     string          `norm:"from_address"`
//	ToAddress       string          `norm:"to_address"`
//	Amount          string          `norm:"amount"`
//}

var _ norm.IVertex = new(Address)
var _ norm.IEdge = new(NebulaTransaction)

func (*Address) TagName() string {
	return "address"
}

func (t *Address) GetVid() interface{} {
	return t.Address
}

func (p *NebulaTransaction) EdgeName() string {
	return "transaction"
}
