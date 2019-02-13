package main

import (
	"bankBigData/AutomaticTask/module"
	"bankBigData/_public/config"
	"bankBigData/_public/log"
	"bankBigData/_public/table"
	"gitee.com/johng/gf/g"
)

func main() {
	g.Config().SetFileName("config.json")
	table.DbDefaultName = g.Config().GetString("dbDefaultName")
	log.Init(config.TaskNameSpace)
	g.DB().SetDebug(true)
	g.DB(table.CDbName).SetDebug(true)
	g.DB(table.IDBName).SetDebug(true)
	g.DB(table.CRDBName).SetDebug(true)
	e := module.InitTable()
	if e != nil {
		return
	}
	//gcron.Add("5 * * * * *", func() {
	//	module.AutoLoadData()
	//})
	module.AutoLoadData()
	//pub_ftp.Test()
	//select {}
}
