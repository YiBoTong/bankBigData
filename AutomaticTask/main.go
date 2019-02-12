package main

import (
	"bankBigData/_public/config"
	"bankBigData/_public/ftp"
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
	//module.AutoLoadData()
	pub_ftp.Test()
	select { }
}
