package nebulaModels

import (
	"github.com/zhihu/norm"
	"time"
)

type Address struct {
	norm.VModel
	Address string  `norm:"address"`
	Type    int     `norm:"type"`
	//Value   float64 `norm:"value"`
}

type NebulaTransaction struct {
	norm.EModel
	TxHash          string    `norm:"tx_hash"`
	TxTime          time.Time `norm:"tx_time"`
	ContractAddress string    `norm:"contract_address"`
	FromAddress     string    `norm:"from_address"`
	ToAddress       string    `norm:"to_address"`
	Amount          float64   `norm:"amount"`
}

//type NebulaTransaction struct {
//	norm.EModel
//	//BlockId         int64     `norm:"block_id"`
//	TxHash          string    `norm:"tx_hash"`
//	//TxType          int       `norm:"tx_type"`
//	TxTime          time.Time `norm:"tx_time"`
//	ContractAddress string    `norm:"contract_address"`
//	From            string    `norm:"from"`
//	To              string    `norm:"to"`
//	//Value           string    `norm:"value"`
//	Amount          float64   `norm:"amount"`
//	//TxStatus        int       `norm:"tx_status"`
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
