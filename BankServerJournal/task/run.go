package task

import (
	"bankBigData/BankServerJournal/db/file"
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/excel"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"fmt"
	"gitee.com/johng/gf/g"
	"github.com/360EntSecGroup-Skylar/excelize"
	"time"
)

var hasTaskRunning = false

var (
	excelInfoFileName  = "" // 农民工信息统计表
	excelLoanFileName  = "" // 不良贷款信息统计表
	excelEntreFileName = "" // 创业情况统计表
	excelTempPath      = "" // 模版文件夹
	excelFilePath      = "" // 存放文件夹
)

func initConfig() {
	excelInfoFileName = g.Config().GetString("excelInfoFileName")
	excelLoanFileName = g.Config().GetString("excelLoanFileName")
	excelEntreFileName = g.Config().GetString("excelEntreFileName")
	excelTempPath = g.Config().GetString("excelTempPath")
	excelFilePath = g.Config().GetString("excelFilePath")
}

func RunTask(restart ...bool) {
	if excelInfoFileName == "" {
		initConfig()
	}
	if len(restart) > 0 {
		hasTaskRunning = false
	}
	if hasTaskRunning {
		return
	}
	hasTaskRunning = true
	firstTask, err := db_file.GetFirstTask()
	if err == nil && firstTask.State == "doing" {
		run(firstTask)
		RunTask(true)
	}
	hasTaskRunning = false
}

func run(task entity.TaskFileItem) {
	_ = db_file.Update(task.Id, g.Map{"created_time": util.GetLocalNowTimeStr()})
	switch task.Type {
	case "info":
		runInfoTask(task)
	case "loan":
		runLoanTask(task)
	case "entre":
		runEntreTask(task)
	}
	_ = db_file.Update(task.Id, g.Map{"state": "done"})
}

func runInfoTask(task entity.TaskFileItem) {
	pageUsers := PageUser(1)
	if len(pageUsers) < 1 {
		return
	}
	iStartPage = 2
	infoColNum = 6
	xlsx, e := excelize.OpenFile(excelTempPath + excelInfoFileName)
	excel.ModifyExcelCellByAxis(xlsx, infoSheetName, excel.ChangIndexToAxis(2, 10), fmt.Sprintf("填表日期：%v", time.Now().Format("2006年01月02日")))
	done := writeInfoFile(xlsx, pageUsers)
	if e == nil && done {
		// 根据指定路径保存文件
		e = xlsx.SaveAs(excelFilePath + task.FilePath)
	}
	if e != nil {
		log.Instance().Error(e)
	}
	xlsx = nil
}

func runLoanTask(task entity.TaskFileItem) {
	pageUsers := PageUser(1)
	if len(pageUsers) < 1 {
		return
	}
	lStartPage = 2
	loanColNum = 6
	xlsx, e := excelize.OpenFile(excelTempPath + excelLoanFileName)
	excel.ModifyExcelCellByAxis(xlsx, loanSheetName, excel.ChangIndexToAxis(2, 16), fmt.Sprintf("填表日期：%v", time.Now().Format("2006年01月02日")))
	done := writeLoanFile(xlsx, pageUsers)
	if e == nil && done {
		// 根据指定路径保存文件
		e = xlsx.SaveAs(excelFilePath + task.FilePath)
	}
	if e != nil {
		log.Instance().Error(e)
	}
	xlsx = nil
}

func runEntreTask(task entity.TaskFileItem) {
	pageUsers := PageUser(1)
	if len(pageUsers) < 1 {
		return
	}
	eStartPage = 2
	entreColNum = 6
	xlsx, e := excelize.OpenFile(excelTempPath + excelLoanFileName)
	excel.ModifyExcelCellByAxis(xlsx, loanSheetName, excel.ChangIndexToAxis(2, 16), fmt.Sprintf("填表日期：%v", time.Now().Format("2006年01月02日")))
	done := writeEntreFile(xlsx, pageUsers)
	if e == nil && done {
		// 根据指定路径保存文件
		e = xlsx.SaveAs(excelFilePath + task.FilePath)
	}
	if e != nil {
		log.Instance().Error(e)
	}
	xlsx = nil
}
