package nebulaModels

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zhihu/norm"
	"github.com/zhihu/norm/constants"
	"math/big"
)

func PrepareTxs(db *norm.DB) error {
	createSchema :=
		"CREATE EDGE IF NOT EXISTS transaction(tx_hash string, tx_time DATETIME, contract_address string, from_address string, to_address string, amount double);" +
			"CREATE EDGE INDEX txs_index on transaction();"
	_, err := db.Execute(createSchema)
	return err
}

func InsertTxn(db *norm.DB, txs *NebulaTransaction) error {
	rank := getTxRank(txs.TxHash)
	tx := NebulaTransaction{
		EModel: norm.EModel{
			Src:       txs.FromAddress,
			SrcPolicy: constants.PolicyNothing,
			Dst:       txs.ToAddress,
			DstPolicy: constants.PolicyNothing,
			Rank:      rank,
		},
		TxHash:          txs.TxHash,
		TxTime:          txs.TxTime,
		ContractAddress: txs.ContractAddress,
		FromAddress:     txs.FromAddress,
		ToAddress:       txs.ToAddress,
		Amount:          txs.Amount,
	}
	err := db.InsertEdge(&tx)
	if err != nil {
		beego.Error("InsertTxn err: ", err)
	} else {
		beego.Info("InsertTxn success")
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
