package task

import (
	"bankBigData/AutomaticTask/dbConfig/table"
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"sync"
	"time"
)

// 每次分析表数量
const analysisNum = 3

// 基础偏移量
var baseOffset = 0

// 同时分析表数量
var readTableNum = 3

func InitAnalysisConf() {
	if ServerId > 0 {
		baseOffset = (ServerId - 1) * ServerSize
	}
	if ServerId > 1 {
		baseOffset += 1
	}
	readTableNum = g.Config().GetInt("readTableNum")
}

var AnalysisRunNum = 0

func analysis(task c_entity.TableTaskTime) bool {
	offset := (ServerId-1)*analysisNum + baseOffset
	limitSize := 0
	if ServerId > 0 {
		limitSize = ServerSize
	}
	log.Instance().Println("正在进行索引：", ServerId)
	tables, e := dbc_table.Page(offset, limitSize)
	tableMap := []c_entity.Table{}
	wg := sync.WaitGroup{}
	if e == nil && len(tables) > 0 {
		for _, v := range tables {
			item := c_entity.Table{}
			if ok := gconv.Struct(v, &item); ok == nil {
				tableMap = append(tableMap, item)
			}
		}

		for {
			tbLen := len(tableMap)
			if AnalysisRunNum >= readTableNum {
				time.Sleep(time.Duration(1) * time.Second)
				continue
			}
			if tbLen == 0 {
				break
			}
			AnalysisRunNum++
			wg.Add(1)
			tableItem := tableMap[0]
			if tableItem.Id == 0 {
				AnalysisRunNum -= 1
				break
			}
			tableMap = tableMap[1:]
			log.Instance().Println("等待分析表数：", tbLen)
			go analysisTaskFileByTable(task, tableItem, &wg)
		}
	}
	tables = nil
	tableMap = nil
	wg.Wait()
	return e != nil
}

func analysisTaskFileByTable(task c_entity.TableTaskTime, table c_entity.Table, wg *sync.WaitGroup) {
	log.Instance().Println("开始分析表：", table.TableName)
	taskFiles, e := dbc_tableTaskFile.Gets(task.Date, g.Slice{"all", "add", "unSplit"}, g.Map{"table_name": table.TableName, "is_down": 1, "db_do": 0})
	entityFile := c_entity.TableTaskFile{}
	dataFile := c_entity.TableTaskFile{}
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
	wg.Done()
	AnalysisRunNum -= 1
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
