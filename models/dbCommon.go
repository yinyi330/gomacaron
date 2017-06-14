package models

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/yikeso/gomacaron/config"
	"github.com/yikeso/gomacaron/config"
	"github.com/jmoiron/sqlx"
)

const TIMESTAMP_FORMATE = "2006-01-02 03:04:05"

var resourceDb *sqlx.DB
var errorLogDb *sqlx.DB

func init(){
	initDb()
}

func initDb(){
	node := config.Read("common","runmodel")
	var err error
	resourceDb,err = sqlx.Connect(config.Read(node,"drivername1"), config.Read(node,"datasourcename1"))
	if err != nil {
		panic(err)
	}
	errorLogDb,err = sqlx.Connect(config.Read(node,"drivername2"), config.Read(node,"datasourcename2"))
	if err != nil {
		panic(err)
	}
}
//获取错误日志的事务
func GetErrorLogTx()(tx *sqlx.Tx,err error){
	return errorLogDb.Beginx()
}
//获取电子书资源的事务
func GetResourceTx()(tx *sqlx.Tx,err error){
	return resourceDb.Beginx()
}