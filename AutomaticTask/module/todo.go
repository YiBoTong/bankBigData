package module

import (
	"bankBigData/AutomaticTask/dbInit"
	"bankBigData/AutomaticTask/initTable"
	"bankBigData/AutomaticTask/task"
	"bankBigData/_public/ftp"
	"bankBigData/_public/log"
	"fmt"
	"gitee.com/johng/gf/g"
)

func InitConf() {
	if pub_ftp.FtpFolder == "" {
		pub_ftp.FtpFolder = g.Config().GetString("ftpFolder")
	}
	if task.ServerId == -1 {
		task.ServerId = g.Config().GetInt("serverId")
	}
	if task.ServerId < 1 {
		task.ServerSize = 0
	} else {
		task.ServerSize = g.Config().GetInt("serverSize")
	}
	task.InitAnalysisConf()
}

func InitTable() error {
	e := error(nil)
	hasDb := initTable.HasDb()
	if !hasDb {
		fmt.Println("不存在数据库，即将初始化")
		e = dbInit.InitDb()
	} else {
		fmt.Println("存在数据库")
	}
	if e != nil {
		log.Instance().Println("数据库初始化错误：", e)
	}
	return e
}

func AutoLoadData() {
	e := error(nil)
	// 只有serverId为0或者1的服务才能进行创建任务，其他的都是执行任务
	if task.ServerId < 2 {
		e = task.Create()
	}
	if e == nil && !task.TaskRuning {
		task.TaskRuning = true
		e = task.Run()
	}
	task.TaskRuning = false
	if e != nil {
		log.Instance().Error("任务执行失败：", e)
	}
}
