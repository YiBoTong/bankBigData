package main

import (
	"bankBigData/AutomaticTask/module"
	"bankBigData/_public/config"
	"bankBigData/_public/log"
	"gitee.com/johng/gf/g"
)

func main() {
	g.Config().SetFileName("config.json")
	log.Init(config.BankNameSpace)
	g.DB().SetDebug(true)
	//gcron.Add("20 * * * * *", func() {
	//	//module.AutoLoadData()
	//})
	module.AutoLoadData()
	select { }
}
