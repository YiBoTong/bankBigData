package task

import (
	"bankBigData/AutomaticTask/db"
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"bufio"
	"gitee.com/johng/gf/g"
	"io"
	"os"
	"strconv"
	"strings"
)

type Structure struct {
	Index int    `json:"index"`
	Key   string `json:"key"`
	Name  string `json:"name"`
}

// @Title 根据表名和日期读取结构文件
func ReadStructureFile(filePath string) ([]Structure, error) {
	structureArr := []Structure{}
	log.Instance().Println("正在分析结构文件")
	fileData, err := os.Open(filePath)
	defer fileData.Close()
	if err != nil {
		return structureArr, err
	}
	buf := bufio.NewReader(fileData)
	for {
		// 读取每一行数据
		line, err := buf.ReadBytes('\n')
		if err != nil || err == io.EOF {
			break
		}
		// 转换编码之后格式化并追加到结构体数据数组中
		structureArr = append(structureArr, FormatStructureLineData(util.ConvertToString(string(line), "GBK", "UTF-8")))
	}
	return structureArr, nil
}

// 读取数据文件
func readTableDataFile(entityFile, dataFile string, readLineNum int, readNum chan int, listMap chan g.List) {
	e := error(nil)
	defer close(listMap)
	defer close(readNum)
	hasDataFile := false
	hasEntityFile := false
	hasEntityFile, _ = util.PathExists(entityFile)
	hasDataFile, _ = util.PathExists(entityFile)
	fileData := &os.File{}
	structure := []Structure{}
	if hasEntityFile && hasDataFile {
		structure, e = ReadStructureFile(entityFile)
		if len(structure) > 0 && e == nil {
			log.Instance().Println("正在分析数据文件")
			// 读取文件
			fileData, e = os.Open(dataFile)
			defer fileData.Close()
		}
		if len(structure) > 0 && e == nil {
			buf := bufio.NewReader(fileData)
			box := g.List{}
			num := 0
			for num = 0; true; num++ {
				// 读取每一行数据
				line, err := buf.ReadBytes('\n')
				if err != nil || err == io.EOF {
					break
				}
				if num < readLineNum {
					continue
				}
				if num != 0 && num%500 == 0 {
					readNum <- num
					listMap <- box
					box = g.List{}
				}
				box = append(box, FormatTableLineData(util.ConvertToString(string(line), "GBK", "UTF-8"), structure))
			}
			log.Instance().Println("本次分析数据数量：", num)
			if len(box) > 0 {
				readNum <- num
				listMap <- box
			}
		}
	} else {
		log.Instance().Println("文件不存在：", entityFile, dataFile)
	}
}

func WriteDataToTable(tbMap c_entity.Table, dataMap c_entity.TableTaskFile, readNum chan int, listMap chan g.List, flg chan int, err chan bool) {
	e := error(nil)
	// 转换编码之后格式化并追加到结构体数据数组中
	for {
		rv, rok := <-readNum
		if v, ok := <-listMap; ok && rok && len(v) > 0 {
			switch dataMap.Types {
			case "add":
				e = db.AddUpdate(tbMap.TableName, v)
			default:
				e = db.AllUpdate(tbMap.TableName, v)
			}
			if e != nil {
				break
			}
			_, _ = dbc_tableTaskFile.Update(dataMap.Date, dataMap.FileName, g.Map{"line_num": rv}, g.Map{})
		} else {
			break
		}
	}
	if e != nil {
		log.Instance().Error("写表错误：", tbMap.TableName, " 原因：", e)
	}
	err <- (e != nil)
	flg <- 1
}

// 格式化结构数据
func FormatStructureLineData(str string) Structure {
	structureArr := strings.Split(str, "|")
	index, err := strconv.Atoi(structureArr[0])
	if err != nil {
		index = 0
	}
	structure := Structure{index - 1, strings.ToLower(structureArr[1]), strings.Replace(structureArr[len(structureArr)-1], "\n", "", -1)}
	return structure
}

// 格式化表格数据
func FormatTableLineData(str string, structure []Structure) g.Map {
	tableData := g.Map{}
	tableArr := strings.Split(str, "|$|")
	tableArr[len(tableArr)-1] = strings.Replace(tableArr[len(tableArr)-1], "\n", "", -1)
	for index, item := range structure {
		tableData[string(item.Key)] = tableArr[index]
	}
	return tableData
}
