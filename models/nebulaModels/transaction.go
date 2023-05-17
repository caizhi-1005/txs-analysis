package nebulaModels

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"github.com/server/txs-analysis/models/dbModels"
	"github.com/vesoft-inc/nebula-go/v3/nebula"
	"github.com/zhihu/norm"
	"github.com/zhihu/norm/constants"
	"math/big"
)

func PrepareTxs(db *norm.DB) error {
	createSchema :=
		"CREATE EDGE IF NOT EXISTS transaction(tx_hash string, tx_time string, contract_address string, from_address string, to_address string, amount string);" +
			"CREATE EDGE INDEX txs_index on transaction();"
	_, err := db.Execute(createSchema)
	return err
}

func InsertTxn(db *norm.DB, tx dbModels.ResTransaction) error {
	rank := getTxRank(tx.TxHash)
	txNebula := NebulaTransaction{
		EModel: norm.EModel{
			Src:       tx.From,
			SrcPolicy: constants.PolicyNothing,
			Dst:       tx.To,
			DstPolicy: constants.PolicyNothing,
			Rank:      rank,
		},
		TxHash:      tx.TxHash,
		TxTime:      tx.TxTime.String(),
		FromAddress: tx.From,
		ToAddress:   tx.To,
		Amount:      tx.Amount,
	}
	fmt.Println("tx.TxTime.String():", tx.TxTime.String())
	err := db.InsertEdge(&txNebula)
	if err != nil {
		beego.Error("InsertTxn err: ", err, " txHash:", tx.TxHash)
	} else {
		beego.Info("InsertTxn success! txHash:", tx.TxHash)
	}
	return err
}

func getTxRank(txHash string) int {
	hash := common.HexToHash(txHash)
	rank := new(big.Int)
	rank.SetBytes(hash[:])
	rankTrim := rank.Uint64() % 1000000
	return int(rankTrim)
}

func QueryTxRoute(db *norm.DB, address string) ([]*TxsRoute, error) {
	nql := fmt.Sprintf("MATCH p=(v:address)-[e:transaction*]->(v2:address{address:\"%s\"}) RETURN e AS p", address)
	result := make([]map[string]interface{}, 0)
	res, err := db.Debug().Execute(nql)
	if err != nil {
		return nil, err
	} else {
		err := UnmarshalResultSet(res, &result)
		if err != nil {
			return nil, err
		}
		paths := make([]*TxsRoute, 0, len(result))

		for _, vpath := range result {
			for _, v := range vpath {
				if path, ok := v.(*nebula.NList); ok {
					pathValue := path.GetValues()
					steps := ParseTxInfo(pathValue)
					tokenRoute := new(TxsRoute)
					tokenRoute.Steps = steps
					paths = append(paths, tokenRoute)
				}
			}
		}
		return paths, nil
	}
}

func ParseTxInfo(pathValue []*nebula.Value) []RouteTxStep {
	txs := make([]TransactionEdge, 0)
	steps := make([]RouteTxStep, 0)
	for _, value := range pathValue {
		if value.EVal != nil {
			tx := GetTxEdgeInfoFromProps(value.EVal)
			txs = append(txs, tx)
			routeStep := RouteTxStep{
				//Src: tx.FromAddress,
				//Dst: tx.ToAddress,
				Transaction: txs,
			}
			steps = append(steps, routeStep)
		}
	}
	return steps
}

const (
	TxHash      = "tx_hash"
	TxTime      = "tx_time"
	FromAddress = "from_address"
	ToAddress   = "to_address"
	Amount      = "amount"
)

func GetTxEdgeInfoFromProps(edge *nebula.Edge) TransactionEdge {
	tx := TransactionEdge{}
	if txHash, exist := edge.Props[TxHash]; exist {
		fmt.Println("txHash:", getValueofValue(txHash))
		tx.TxHash = getValueofValue(txHash)
	}
	if txTime, exist := edge.Props[TxTime]; exist {
		fmt.Println("txTime:", txTime)
		tx.TxTime = getValueofValue(txTime)
	}
	if fromAddress, exist := edge.Props[FromAddress]; exist {
		fmt.Println("fromAddress:", fromAddress)
		tx.FromAddress = getValueofValue(fromAddress)
	}
	if toAddress, exist := edge.Props[ToAddress]; exist {
		fmt.Println("toAddress:", toAddress)
		tx.ToAddress = getValueofValue(toAddress)
	}
	if amount, exist := edge.Props[Amount]; exist {
		fmt.Println("amount:", amount)
		tx.Amount = getValueofValue(amount)
	}
	return tx
}
