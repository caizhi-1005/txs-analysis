package nebulaModels

import (
	"github.com/astaxie/beego"
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

	nebulaDB := Init()

	//准备tag
	err := PrepareAddress(nebulaDB)
	if err != nil {
		beego.Error("nebulaModels.PrepareAddress error: ", err)
	}
	//准备edge
	err2 := PrepareTxs(nebulaDB)
	if err2 != nil {
		beego.Error("nebulaModels.PrepareTxs error: ", err2)
	}
	nebulaDB.Close()
	beego.Info("nebula prepare success")
}


func Init() *norm.DB {
	var nebulaDB *norm.DB
	var once sync.Once
	once.Do(func() {
		nebulaDB = newDB()
	})
	return nebulaDB
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
