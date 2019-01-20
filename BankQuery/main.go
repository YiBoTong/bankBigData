package main

import (
	"bankBigData/BankQuery/handler"
	"bankBigData/_public/config"
	"bankBigData/_public/log"
	"bankBigData/_public/router"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/net/ghttp"
)

const (
	apiPath = "/api/" + config.BankNameSpace
)

func main()  {
	g.Config().SetFileName("config.json")
	log.Init(config.BankNameSpace)
	g.DB().SetDebug(true)
	s := g.Server(config.BankNameSpace)
	s.SetSessionIdName(config.CookieIdName)
	// 贷款
	_ = s.BindController(apiPath+"/public/query/loan", new(handler.Loan))
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		500: router.Status_500,
	})
	s.SetLogPath(config.LogPath)
	s.SetAccessLogEnabled(true)
	s.SetErrorLogEnabled(true)
	s.SetPort(g.Config().GetInt("appPort"))
	_ = s.Run()
}
