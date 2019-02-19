package main

import (
	"bankBigData/AutomaticTask/module"
	"bankBigData/_public/config"
	"bankBigData/_public/log"
	"bankBigData/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/os/os/gcron"
)

func main() {
	g.Config().SetFileName("config.json")
	debug := g.Config().GetBool("debug")
	table.DbDefaultName = g.Config().GetString("dbDefaultName")
	log.Init(config.TaskNameSpace)
	g.DB().SetDebug(debug)
	g.DB(table.CDbName).SetDebug(debug)
	g.DB(table.IDBName).SetDebug(debug)
	g.DB(table.CRDBName).SetDebug(debug)
	module.InitConf()
	e := module.InitTable()
	if e != nil {
		return
	}
	// 自动定时执行任务
	_ = gcron.Add(g.Config().GetString("taskRunTime"), func() {
		module.AutoLoadData()
	})
	select {}
}
