package nebulaModels

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/server/txs-analysis/models/nebulaModels/types"
	"github.com/vesoft-inc/nebula-go/v3/nebula"
	"github.com/zhihu/norm"
	"github.com/zhihu/norm/dialectors"
	"sync"
	"time"
)

type Config struct {
	DbHost     string `json:"db_host"`
	DbSpace    string `json:"db_space"`
	DbUser     string `json:"db_username"`
	DbPassword string `json:"db_password"`
}

var nebulaConfig *Config

func init() {
	nebulaConfig = &Config{
		DbHost:     beego.AppConfig.String("nebula::dbhost"),
		DbSpace:    beego.AppConfig.String("nebula::dbspace"),
		DbUser:     beego.AppConfig.String("nebula::dbuser"),
		DbPassword: beego.AppConfig.String("nebula::dbpassword"),
	}
	beego.Info("nebula config init success")
}

func newDB() *norm.DB {
	dialector := dialectors.MustNewNebulaDialector(dialectors.DialectorConfig{
		Addresses: []string{nebulaConfig.DbHost},
		Timeout:   time.Second * 5,
		Space:     nebulaConfig.DbSpace,
		Username:  nebulaConfig.DbUser,
		Password:  nebulaConfig.DbPassword,
	})
	db := norm.MustOpen(dialector, norm.Config{})
	return db
}

func Init() *norm.DB {
	var nebulaDB *norm.DB
	var once sync.Once
	once.Do(func() {
		nebulaDB = newDB()
	})
	return nebulaDB
}

func getValueofValue(value *nebula.Value) string {
	if value.NVal != nil {
		return fmt.Sprintf("value.NVal=%v", value.NVal)
	}
	if value.BVal != nil {
		return fmt.Sprintf("value.BVal=%v", value.BVal)
	}
	if value.IVal != nil {
		return fmt.Sprintf("value.IVal=%v", value.IVal)
	}
	if value.FVal != nil {
		return fmt.Sprintf("value.FVal=%v", value.FVal)
	}
	if value.SVal != nil {
		return fmt.Sprintf("%v", string(value.SVal))
	}
	if value.DVal != nil {
		return fmt.Sprintf("value.DVal=%v", value.DVal)
	}
	if value.TVal != nil {
		return fmt.Sprintf("value.TVal=%v", value.TVal)
	}
	if value.DtVal != nil {
		return fmt.Sprintf("value.DtVal=%v", value.DtVal)
	}
	if value.VVal != nil {
		return fmt.Sprintf("value.VVal=%v", value.VVal)
	}
	if value.EVal != nil {
		return fmt.Sprintf("value.EVal=%v", value.EVal)
	}
	if value.PVal != nil {
		return fmt.Sprintf("value.PVal=%v", value.PVal)
	}
	if value.LVal != nil {
		return fmt.Sprintf("value.LVal=%v", value.LVal)
	}
	if value.MVal != nil {
		return fmt.Sprintf("value.MVal=%v", value.MVal)
	}
	if value.UVal != nil {
		return fmt.Sprintf("value.UVal=%v", value.UVal)
	}
	if value.GVal != nil {
		return fmt.Sprintf("value.GVal=%v", value.GVal)
	}
	if value.GgVal != nil {
		return fmt.Sprintf("value.GgVal=%v", value.GgVal)
	}
	if value.DuVal != nil {
		return fmt.Sprintf("value.DuVal=%v", value.DuVal)
	}
	return ""
}

func getValueofTag(tag *nebula.Tag) string {
	v := ""
	s := fmt.Sprintf("name:%s", string(tag.Name))
	v += s
	for k, p := range tag.Props {
		s = fmt.Sprintf("[%s]=%s\n", k, getValueofValue(p))
		v += s
	}
	return v
}

func getValueofTags(tags []*nebula.Tag) string {
	v := ""
	for i, tag := range tags {
		s := fmt.Sprintf("t[%d]=%s\n", i, getValueofTag(tag))
		v += s
	}
	return v
}

func printVertex(v *nebula.Vertex) {
	fmt.Println("vertex.VID = %s", getValueofValue(v.Vid))
	if len(v.Tags) > 0 {
		fmt.Println("vertex.Tags = %s", getValueofTags(v.Tags))
	}
}

func printStep(index int, s *nebula.Step) {
	//Dst *Vertex `thrift:"dst,1" db:"dst" json:"dst"`
	//Type EdgeType `thrift:"type,2" db:"type" json:"type"`
	//Name []byte `thrift:"name,3" db:"name" json:"name"`
	//Ranking EdgeRanking `thrift:"ranking,4" db:"ranking" json:"ranking"`
	//Props map[string]*Value `thrift:"props,5" db:"props" json:"props"`
	//log.Infof("step[%d].dst = %s", index, getValueofValue(s.Dst.Vid))
	fmt.Println("step[%d].dst = %s", index, getValueofValue(s.Dst.Vid))
	//printVertex(s.Dst)
	//log.Infof("step[%d].edgetype = %d", index, s.Type)
	//log.Infof("step[%d].rank = %d", index, s.Ranking)
	//log.Infof("step[%d].name = %s", index, string(s.Name))
	for k, v := range s.Props {
		//log.Infof("step[%d].prop[%s]=%v", index, k, getValueofValue(v))
		fmt.Println("step[%d].prop[%s]=%v", index, k, getValueofValue(v))
	}
	//log.Infof("step[%d].props = %v", index, s.Props)

}

func GetDstFromStep(step *nebula.Step) string {
	return getValueofValue(step.Dst.GetVid())
}

func GetPairInfoFromStep(step *nebula.Step, routeStep *types.RouteStep) {
	Pairs := make([]types.RoutePairInfo, 1)

	if dex, exist := step.Props[PairProp_dex]; exist {
		Pairs[0].Dex = getValueofValue(dex)
	}
	if pairAddr, exist := step.Props[PairProp_paircontract]; exist {
		Pairs[0].Pair = getValueofValue(pairAddr)
	}
	if fee, exist := step.Props[PairProp_fee]; exist {
		Pairs[0].Fee = getValueofValue(fee)
	}
	routeStep.Pairs = Pairs
}

const (
	TxHash          = "tx_hash"
	TxTime          = "tx_time"
	ContractAddress = "contract_address"
	FromAddress     = "from_address"
	ToAddress       = "to_address"
	Amount          = "amount"
)

func GetTxInfoFromStep(step *nebula.Step, routeStep *types.RouteTxStep) {
	Txs := make([]types.ResTransaction, 1)

	if dex, exist := step.Props[TxHash]; exist {
		Txs[0].TxHash = getValueofValue(dex)
	}
	if pairAddr, exist := step.Props[TxTime]; exist {
		Txs[0].TxTime = getValueofValue(pairAddr)
	}
	if fee, exist := step.Props[ContractAddress]; exist {
		Txs[0].ContractAddress = getValueofValue(fee)
	}
	if fee, exist := step.Props[FromAddress]; exist {
		Txs[0].FromAddress = getValueofValue(fee)
	}
	if fee, exist := step.Props[ToAddress]; exist {
		Txs[0].ToAddress = getValueofValue(fee)
	}
	if fee, exist := step.Props[Amount]; exist {
		Txs[0].Amount = getValueofValue(fee)
	}
	routeStep.TbTransaction = Txs
}

func ParsePathInfo(path *nebula.Path) []types.RouteStep {
	src := path.GetSrc()
	steps := path.GetSteps()
	routePath := make([]types.RouteStep, len(steps))
	srcToken := getValueofValue(src.Vid)
	for i, step := range steps {
		routeStep := types.RouteStep{
			Src: srcToken,
		}
		routeStep.Dst = GetDstFromStep(step)
		GetPairInfoFromStep(step, &routeStep)
		srcToken = routeStep.Dst
		routePath[i] = routeStep
	}
	return routePath
}

func ParseTxPathInfo(path *nebula.Path) []types.RouteTxStep {
	src := path.GetSrc()
	steps := path.GetSteps()
	routePath := make([]types.RouteTxStep, len(steps))
	srcToken := getValueofValue(src.Vid)
	for i, step := range steps {
		routeStep := types.RouteTxStep{
			Src: srcToken,
		}
		routeStep.Dst = GetDstFromStep(step)
		GetTxInfoFromStep(step, &routeStep)
		srcToken = routeStep.Dst
		routePath[i] = routeStep
	}
	return routePath
}

func QueryRoute(db *norm.DB, token0, token1 string) []*types.TxsRoute {
	nql := fmt.Sprintf("FIND NOLOOP PATH WITH PROP FROM \"%s\" TO \"%s\" OVER * YIELD path AS p", token0, token1)
	result := make([]map[string]interface{}, 0)
	res, err := db.Debug().Execute(nql)
	if err != nil {
		//log.WithField("err", err).Error("query route failed")
		return []*types.TxsRoute{}
	} else {
		//log.WithField("rows", len(res.GetRows())).Info("query route")
		err := UnmarshalResultSet(res, &result)
		if err != nil {
			//log.WithField("err", err).Error("parse route failed")
			return []*types.TxsRoute{}
		}
		paths := make([]*types.TxsRoute, 0, len(result))

		for _, vpath := range result {
			// vpath only have one key (AS p)
			for _, v := range vpath {
				if path, ok := v.(*nebula.Path); ok {
					steps := ParseTxPathInfo(path)
					tokenRoute := new(types.TxsRoute)
					tokenRoute.Steps = steps
					paths = append(paths, tokenRoute)
				}
			}
		}
		return paths
	}
}

func QueryRouteWithMaxJump(db *norm.DB, token0, token1 string, op int) []*types.TokenRoute {
	nql := fmt.Sprintf("FIND NOLOOP PATH WITH PROP FROM \"%s\" TO \"%s\" OVER * UPTO %d STEPS YIELD path AS p", token0, token1, op)
	result := make([]map[string]interface{}, 0)
	res, err := db.Execute(nql)
	//res, err := db.Debug().Execute(nql)
	if err != nil {
		//log.WithField("err", err).Error("query route failed")
		return []*types.TokenRoute{}
	} else {
		//log.WithField("rows", len(res.GetRows())).Info("query route")
		err := UnmarshalResultSet(res, &result)
		if err != nil {
			//log.WithField("err", err).Error("parse route failed")
			return []*types.TokenRoute{}
		}
		paths := make([]*types.TokenRoute, 0, len(result))

		for _, vpath := range result {
			// vpath only have one key (AS p)
			for _, v := range vpath {
				if path, ok := v.(*nebula.Path); ok {
					steps := ParsePathInfo(path)
					tokenRoute := new(types.TokenRoute)
					tokenRoute.Steps = steps
					paths = append(paths, tokenRoute)
				}
			}
		}
		return paths
	}
}
