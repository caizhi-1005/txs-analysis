package types

import (
	"encoding/json"
	"fmt"
)

type RoutePairInfo struct {
	Pair string `json:"pair"`
	Fee  string `json:"fee"`
	Dex  string `json:"dex"`
}

func TextAddress(addr string) string {
	return fmt.Sprintf("\"%s\"", addr)
}

type RouteStep struct {
	Pairs []RoutePairInfo `json:"pair"`
	Src   string          `json:"from"`
	Dst   string          `json:"to"`
}

type RouteTxStep struct {
	TbTransaction []ResTransaction `json:"transaction"`
	Src           string                   `json:"from"`
	Dst           string                   `json:"to"`
}

type ResTransaction struct {
	TxHash          string `norm:"tx_hash"`
	TxTime          string `norm:"tx_time"`
	ContractAddress string `norm:"contract_address"`
	FromAddress     string `norm:"from_address"`
	ToAddress       string `norm:"to_address"`
	Amount          string `norm:"amount"`
}


type TokenRoute struct {
	Steps []RouteStep `json:"steps"`
}

type TxsRoute struct {
	Steps []RouteTxStep `json:"tx_steps"`
}

func (r TxsRoute) String() string {
	d, _ := json.Marshal(r)
	return string(d)
}

func (r TokenRoute) String() string {
	d, _ := json.Marshal(r)
	return string(d)
}

type SortTokenRoutes []*TokenRoute

func (s SortTokenRoutes) Len() int           { return len(s) }
func (s SortTokenRoutes) Less(i, j int) bool { return len(s[i].Steps) < len(s[j].Steps) }
func (s SortTokenRoutes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

