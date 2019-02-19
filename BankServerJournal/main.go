package main

import (
	"bankBigData/BankServerJournal/handler"
	"bankBigData/BankServerJournal/redis"
	"bankBigData/BankServerJournal/table"
	"bankBigData/_public/config"
	"bankBigData/_public/log"
	"git.91ybt.com/lib/gf/g/os/gcron"
	"gitee.com/johng/gf/g"
)

var apiPath = "/api/qnb_go/journal"

func main() {
	g.Config().SetFileName("config.json")
	debug := g.Config().GetBool("debug")
	log.Init(config.BankServerJournalNameSpace)
	g.DB().SetDebug(debug)
	g.DB(table.PSDBName).SetDebug(debug)
	g.DB(table.BCDDBName).SetDebug(debug)
	s := g.Server(config.BankNameSpace)
	// 贷款
	_ = s.BindController(apiPath+"/query", new(handler.Query))
	// 报表
	_ = s.BindController(apiPath+"/file", new(handler.File))
	s.SetLogPath(g.Config().GetString("logPath"))
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(g.Config().GetInt("appPort"))
	// 自动定时执行任务
	_ = gcron.Add(g.Config().GetString("redisClearTime"), func() {
		log.Instance().Println("清除Redis缓存")
		_ = redis.Clear()
	})
	_ = s.Run()
}
