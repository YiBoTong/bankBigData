package task

import (
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/dbConfig/tableTaskTime"
	"bankBigData/_public/ftp"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"regexp"
	"strings"
	"time"
)

// 是否正在执行扫描
var scanning = false

// 任务是否运行中
var TaskRuning = false

// 日期格式
var dateReg = regexp.MustCompile(`^\d{8}$`)

// 压缩文件
var gzReg = regexp.MustCompile(`\.gz$`)

// 数据文件
var dataFileReg = regexp.MustCompile(`\.(ddl|del|gz|ok)$`)

// 数据校验文件
var dataOkFileReg = regexp.MustCompile(`\.ok$`)

// 增量文件
var addFileReg = regexp.MustCompile(`_add_`)

// 全量文件
var allFileReg = regexp.MustCompile(`_all_`)

// 非拆分文件
var unSplitFileReg = regexp.MustCompile(`_0000000\.del\.gz`)

// 结构文件
var entityFileSuffixReg = regexp.MustCompile(`\.ddl$`)

// 数据文件
var dataFileSuffixReg = regexp.MustCompile(`\.del\.gz$`)

// 跳过指定的日期之前的日期
var AfterDay = 0

// 创建任务
func Create() error {
	log.Instance().Println("开始创建任务")
	yesterday := util.GetYesterday()
	lastTask, e := dbc_tableTaskTime.Last("desc", g.Slice{})
	if e == nil {
		if lastTask.Id == 0 {
			e = ScanAllDate()
		}
		if lastTask.Date != yesterday {
			log.Instance().Println("添加指定日期任务：", yesterday)
			e = Add([]string{yesterday})
		}
		if scanning == false {
			e = Scan()
		}
		log.Instance().Println("任务创建完毕")
	} else {
		log.Instance().Println("任务创建失败：", e)
	}
	return e
}

// 扫描所有日期
func ScanAllDate() error {
	allFileList := pub_ftp.GetNameList("/")
	allDate := []string{}
	for _, v := range allFileList {
		v = strings.Replace(v, "/", "", 1)
		if dateReg.MatchString(v) && gconv.Int(v) >= AfterDay {
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
	lastTask, e := dbc_tableTaskTime.Last("asc", g.Slice{}, g.Map{
		"is_scan": 0,
	})
	if lastTask.Id != 0 {
		e = ScanDate(lastTask.Date)
		log.Instance().Println("扫描结束", lastTask.Date)
		if e == nil {
			// 5秒后进行下一轮扫描
			time.Sleep(time.Duration(5) * time.Second)
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
			if !dataFileReg.MatchString(fileName) {
				continue
			}
			types := "other"
			if allFileReg.MatchString(fileName) {
				types = "all"
			} else if addFileReg.MatchString(fileName) {
				types = "add"
			}
			tableName := strings.Split(fileName, "_"+types+"_")[0]
			if unSplitFileReg.MatchString(fileName) {
				types = "unSplit"
			} else if dataOkFileReg.MatchString(fileName) {
				types = "ok"
			}
			contentType := 0
			if entityFileSuffixReg.MatchString(fileName) {
				contentType = 1
			} else if dataFileSuffixReg.MatchString(fileName) {
				contentType = 2
			}

			taskFiles = append(taskFiles, g.Map{
				"date":         date,
				"file_name":    fileName,
				"table_name":   tableName,
				"types":        types,
				"content_type": contentType,
				"is_gz":        gzReg.MatchString(fileName),
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
