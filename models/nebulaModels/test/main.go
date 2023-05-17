package main

import (
	"fmt"
	"github.com/server/txs-analysis/models/nebulaModels"
)

func main() {
	fmt.Println("got query")
	//from:="0xa3a3c435c6f5a7abad47b10473ccf9e3de7b3bd4"
	//from:="0xb234ea5dc59a761038a7ea6278812aaafbd2afc4"
	//to:="0xbfb6e0d4896c32f1f61d9c6bdd84884e4957d328"
	//to:="0x9b1b9cb9da5bc96bc505ad22c2c605012c49a340"
	address := "0xbfb6e0d4896c32f1f61d9c6bdd84884e4957d328"

	db := nebulaModels.Init()

	result0 := nebulaModels.QueryTxRoute(db, address)
	fmt.Println("result0:", result0)

	//result := nebulaModels.QueryRoute(db,from,to)
	//fmt.Println("result:", result)
	//
	//result1 := nebulaModels.QueryRouteWithMaxJump(db,from,to,3)
	//fmt.Println("result1:", result1)
}
