package task

import (
	"bankBigData/AutomaticTask/db"
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/ftp"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"bufio"
	"fmt"
	"gitee.com/johng/gf/g"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Structure struct {
	Index int    `json:"index"`
	Key   string `json:"key"`
	Name  string `json:"name"`
}

// @Title 根据表名和日期读取结构文件
func ReadStructureFile(entityFile c_entity.TableTaskFile, filePath string) ([]Structure, error) {
	structureArr := []Structure{}
	log.Instance().Println("正在分析结构文件")
	fileData, err := os.Open(filePath)
	defer fileData.Close()
	if err != nil {
		return structureArr, err
	}
	_, _ = dbc_tableTaskFile.Update(entityFile.Date, entityFile.FileName, g.Map{"db_do_start_time": util.GetLocalNowTimeStr()}, g.Map{})

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
	_, _ = dbc_tableTaskFile.Update(entityFile.Date, entityFile.FileName, g.Map{"db_do_end_time": util.GetLocalNowTimeStr()}, g.Map{})
	return structureArr, nil
}

// 读取数据文件
func readTableDataFile(entityFile, dataFile c_entity.TableTaskFile, readNum chan int, listMap chan []g.List) {
	e := error(nil)
	defer close(listMap)
	defer close(readNum)
	hasDataFile := false
	hasEntityFile := false

	dataFileStr := dataFile.FileName
	if dataFile.IsGz {
		dataFileStr = dataFileStr[:len(dataFileStr)-3]
	}

	basePath := strings.Join([]string{pub_ftp.FtpFolder, entityFile.Date}, "/")

	dataFilePath := strings.Join([]string{basePath, dataFileStr}, "/")
	entityFilePath := strings.Join([]string{basePath, entityFile.FileName}, "/")

	hasEntityFile, _ = util.PathExists(entityFilePath)
	hasDataFile, _ = util.PathExists(dataFilePath)
	structure := []Structure{}
	fileData := &os.File{}
	temp := g.List{}
	box := []g.List{}
	if hasEntityFile && hasDataFile {
		structure, e = ReadStructureFile(entityFile, entityFilePath)
		if len(structure) > 0 && e == nil {
			log.Instance().Println("正在分析数据文件")
			// 读取文件
			fileData, e = os.Open(dataFilePath)
			_, _ = dbc_tableTaskFile.Update(dataFile.Date, dataFile.FileName, g.Map{"db_do_start_time": util.GetLocalNowTimeStr()}, g.Map{})
			defer fileData.Close()
		}
		if len(structure) > 0 && e == nil {
			buf := bufio.NewReader(fileData)

			num := 0
			tempNum := 0
			for num = 0; true; num++ {
				// 读取每一行数据
				line, err := buf.ReadBytes('\n')
				if err != nil || err == io.EOF {
					break
				}
				if num < dataFile.LineNum {
					continue
				}
				if num != 0 && num%500 == 0 {
					box = append(box, temp)
					temp = g.List{}
					tempNum += 1
				}
				if tempNum == 10 {
					tempNum = 0
					readNum <- num
					listMap <- box
					box = []g.List{}
				}
				temp = append(temp, FormatTableLineData(util.ConvertToString(string(line), "GBK", "UTF-8"), structure))
			}
			log.Instance().Println("本次分析数据数量：", num)
			if len(temp) > 0 {
				box = append(box, temp)
			}
			if len(box) > 0 {
				readNum <- num
				listMap <- box
			}
		}
	} else {
		log.Instance().Println("文件不存在：", entityFilePath, dataFilePath)
	}
	structure = nil
	fileData = nil
	temp = nil
	box = nil
}

func WriteData(tbMap c_entity.Table, dataMap c_entity.TableTaskFile, readNum chan int, listMap chan []g.List, err chan bool) {
	e := error(nil)
	wg := sync.WaitGroup{}
	// 转换编码之后格式化并追加到结构体数据数组中
	for {
		rv, rok := <-readNum
		if list, ok := <-listMap; ok && rok && len(list) > 0 {
			swg := sync.WaitGroup{}
			for _, v := range list {
				if len(v) == 0 {
					continue
				}
				wg.Add(1)
				swg.Add(1)
				go func() {
					hasErr := error(nil)
					switch tbMap.Type {
					case "add":
						hasErr = db.AddUpdate(tbMap.TableName, v)
					default:
						hasErr = db.AllUpdate(tbMap.TableName, v)
					}
					if hasErr != nil {
						log.Instance().Error("写表错误：", tbMap.TableName, " 原因：", hasErr)
						_, _ = dbc_tableTaskFile.Update(dataMap.Date, dataMap.FileName, g.Map{"has_err": 1, "err_msg": fmt.Sprintf("%v", hasErr)}, g.Map{})
					}
					wg.Done()
					swg.Done()
				}()
			}
			swg.Wait()
			if e != nil {
				break
			}
			_, _ = dbc_tableTaskFile.Update(dataMap.Date, dataMap.FileName, g.Map{"line_num": rv}, g.Map{})
		} else {
			break
		}
	}
	wg.Wait()
	err <- (e != nil)
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
