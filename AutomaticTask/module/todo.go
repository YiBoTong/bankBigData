package module

import (
	"bankBigData/AutomaticTask/dbInit"
	"bankBigData/AutomaticTask/initTable"
	"bankBigData/AutomaticTask/task"
	"bankBigData/_public/log"
	"fmt"
)

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
	e := task.Create()
	if e == nil {
		e = task.Run()
	}
	if e != nil {
		log.Instance().Error("任务执行失败：", e)
	}
}
