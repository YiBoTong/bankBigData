package task

import (
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/dbConfig/tableTaskTime"
	"bankBigData/_public/ftp"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"regexp"
	"strings"
	"time"
)

// 是否正在执行扫描
var scanning = false

// 日期格式
var dateReg = regexp.MustCompile(`^\d{8}$`)
// 压缩文件
var gzReg = regexp.MustCompile(`\.gz$`)

// 创建任务
func Create() error {
	log.Instance().Println("开始创建任务")
	yesterday := util.GetYesterday()
	lastTask, e := dbc_tableTaskTime.Last("desc")
	if lastTask.Id == 0 {
		e = ScanAllDate()
	} else if lastTask.Date != yesterday {
		log.Instance().Println("添加指定日期任务：", yesterday)
		e = Add([]string{yesterday})
	}
	if e == nil && scanning == false {
		e = Scan()
	}
	log.Instance().Println("任务创建完毕")
	return e
}

// 扫描所有日期
func ScanAllDate() error {
	allFileList := pub_ftp.GetNameList("/")
	allDate := []string{}
	for _, v := range allFileList {
		v = strings.Replace(v, "/", "", 1)
		if dateReg.MatchString(v) {
			allDate = append(allDate, v)
		}
	}
	log.Instance().Println("扫描所有日期")
	return Add(allDate)
}

// 扫描
func Scan() error {
	scanning = true
	log.Instance().Println("开始扫描")
	// 还未扫描过的任务
	lastTask, e := dbc_tableTaskTime.Last("asc", g.Map{
		"is_scan": 0,
	})
	if lastTask.Id != 0 {
		e = ScanDate(lastTask.Date)
		log.Instance().Println("扫描结束", lastTask.Date)
		if e == nil {
			// 5秒后进行下一轮扫描
			time.Sleep(time.Duration(60) * time.Second)
			// 递归扫描任务
			e = Scan()
		} else {
			log.Instance().Error("扫描 "+lastTask.Date+" 出现错误：", e)
		}
	} else {
		scanning = false
		log.Instance().Println("扫描结束")
	}
	return e
}

// 扫描指定日期
func ScanDate(date string) error {
	log.Instance().Println("扫描指定日期：", date)
	e := error(nil)
	taskFiles := g.List{}
	allFileList := pub_ftp.GetNameList("/" + date)
	if len(allFileList) > 0 {
		for _, v := range allFileList {
			fileName := strings.Split(v, "/")[2]
			taskFiles = append(taskFiles, g.Map{
				"date":      date,
				"file_name": fileName,
				"is_gz":     gzReg.MatchString(fileName),
			})
		}
	}
	_, e = dbc_tableTaskFile.Add(date, taskFiles)
	return e
}

// 添加任务
func Add(days []string) error {
	e := error(nil)
	taskDays := g.List{}
	if len(days) > 0 {
		for _, v := range days {
			taskDays = append(taskDays, g.Map{
				"date": v,
			})
		}
		_, e = dbc_tableTaskTime.Add(taskDays)
	}
	return e
}
