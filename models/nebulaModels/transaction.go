package nebulaModels

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/common"
	"github.com/server/txs-analysis/models/dbModels"
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
		TxHash:          tx.TxHash,
		TxTime:          tx.TxTime.String(),
		ContractAddress: tx.ContractAddress,
		FromAddress:     tx.From,
		ToAddress:       tx.To,
		Amount:          tx.Amount,
	}
	fmt.Println("tx.TxTime.String():",tx.TxTime.String())
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
