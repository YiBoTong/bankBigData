package task

import (
	"bankBigData/AutomaticTask/dbConfig/table"
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"math"
)

// 每次分析表数量
const analysisNum = 3

// 基础偏移量
var baseOffset = 0

// 最后一页的limit数量
var offsetPage = 1

// 最后一次偏移数
var lastLimit = 1

func InitAnalysisConf() {
	baseOffset = (ServerId - 1) * ServerSize
	offsetPage = int(math.Floor(gconv.Float64(ServerSize / analysisNum)))
	lastLimit = ServerSize % analysisNum
	if ServerId > 1 {
		baseOffset += 1
		if lastLimit > 0 {
			lastLimit -= 1
		}
	}
}

func analysis(task c_entity.TableTaskTime, page int) bool {
	offset := (page-1)*analysisNum + baseOffset
	limitSize := analysisNum
	if ServerId > 0 {
		if page > offsetPage+1 {
			return true
		}
		if page-1 == offsetPage {
			limitSize = lastLimit
		}
		log.Instance().Println("正在进行", ServerId, "(", baseOffset, "-", baseOffset+ServerSize, ")", "索引：", page)
	} else {
		log.Instance().Println("正在进行索引：", page)
	}
	tables, e := dbc_table.Page(offset, limitSize)
	if e == nil && len(tables) > 0 {
		for _, v := range tables {
			flg := make(chan int)
			item := c_entity.Table{}
			if ok := gconv.Struct(v, &item); ok == nil {
				go analysisTaskFileByTable(task, item, flg)
			}
			<-flg
		}
		return analysis(task, page+1)
	}
	return e != nil
}

func analysisTaskFileByTable(task c_entity.TableTaskTime, table c_entity.Table, flg chan int) {
	taskFiles, e := dbc_tableTaskFile.Gets(task.Date, g.Slice{"all", "add", "unSplit"}, g.Map{"table_name": table.TableName, "is_down": 1, "db_do": 0})
	entityFile := c_entity.TableTaskFile{}
	dataFile := c_entity.TableTaskFile{}
	log.Instance().Println("开始分析表：", table.TableName)
	if e == nil && len(taskFiles) > 0 {
		for _, v := range taskFiles {
			item := c_entity.TableTaskFile{}
			if e = gconv.Struct(v, &item); e == nil {
				if item.ContentType == 1 {
					entityFile = item
				} else if item.ContentType == 2 {
					dataFile = item
				}
			}
		}
		if e == nil {
			e = analysisTaskFile(table, entityFile, dataFile)
		}
	}
	log.Instance().Println("表分析完成：", table.TableName)
	taskFiles = nil
	flg <- 1
}

func analysisTaskFile(table c_entity.Table, entityFile, dataFile c_entity.TableTaskFile) error {
	readNum := make(chan int)
	listMap := make(chan []g.List)
	writeError := make(chan bool)

	log.Instance().Println("开始表文件分析：", table.TableName)

	go readTableDataFile(entityFile, dataFile, readNum, listMap)
	go WriteData(table, dataFile, readNum, listMap, writeError)

	e, _ := <-writeError;
	listMap = nil
	if e == false {
		if entityFile.Date != "" {
			_, _ = dbc_tableTaskFile.Update(dataFile.Date, dataFile.FileName, g.Map{"db_do": 1}, g.Map{"db_do": 0})
			_, _ = dbc_tableTaskFile.Update(entityFile.Date, entityFile.FileName, g.Map{"db_do": 1}, g.Map{"db_do": 0})
		}
		log.Instance().Println("表文件分析完成：", table.TableName)
	} else {
		log.Instance().Println("表文件分析错误：", table.TableName)
	}
	_, _ = dbc_tableTaskFile.Update(dataFile.Date, dataFile.FileName, g.Map{"db_do_end_time": util.GetLocalNowTimeStr()}, g.Map{})
	return nil
}
