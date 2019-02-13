package util

import (
	"bankBigData/AutomaticTask/dbConfig/tableColumn"
	"bankBigData/_public/log"
	"gitee.com/johng/gf/g"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/logs"
	"strings"
)

// 根据excel创建分类名称
//func CreateClassNameByExcel() {
//	tb := "S_ECIF_ECIF_CERT_INFO"
//	key := ""
//	log.Instance().Println("根据Excel解析生成表配置")
//	xis, err := excelize.OpenFile("D:/go/src/bankBigData/_static/excel/附件4_下发数据表结构 (自动保存的).xlsx")
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//	log.Instance().Println("Excle文件读取完成")
//	//tables := g.List{}
//	// 获取工作表中sheet1中的值
//	rows := xis.GetRows("表字典")
//	tableItem := []string{}
//	for _, row := range rows[1:] {
//		fmt.Println(row)
//		types := strings.ToLower(row[3])
//		if types == "timestamp" {
//			types += " null"
//		}
//		def := "NULL"
//		if row[1] == "COD_CUST_ID" {
//			def = "'-'"
//			key = strings.ToLower(row[1])
//		}
//		strItem := []string{
//			"`" + strings.ToLower(row[1]) + "`",
//			types,
//			"DEFAULT " + def,
//			"COMMENT " + "'" + row[2] + "'",
//		}
//		if row[6] == tb {
//			tableItem = append(tableItem, strings.Join(strItem, " "))
//		}
//		// 表明（英文）转换为小写
//		//tables = append(tables, g.Map{
//		//	"table_name": strings.ToLower(row[0]),
//		//	"title":      gconv.String(row[1]),
//		//	"type":       strings.ToLower(row[2]),
//		//})
//	}
//	tableItem = append(tableItem, "PRIMARY KEY (`"+key+"`)")
//	tbName, err := pubFun.GetTableName(strings.ToLower(tb))
//	if err == nil {
//		sql := "CREATE TABLE `" + tbName + "` (" + strings.Join(tableItem, ",") + ") ENGINE=InnoDB DEFAULT CHARSET=utf8";
//		fmt.Println(sql)
//	}
//	//_, _ = db_table.InitTableInfo(tables)
//	logs.Info("初始化完成")
//}
//func CreateClassNameByExcel() {
//	tb := "S_OFCR_CH_ACCT_MAST"
//	key := ""
//	log.Instance().Println("根据Excel解析生成表配置")
//	xis, err := excelize.OpenFile("D:/go/src/bankBigData/_static/excel/附件4_下发数据表结构 (自动保存的).xlsx")
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//	log.Instance().Println("Excle文件读取完成")
//	//tables := g.List{}
//	// 获取工作表中sheet1中的值
//	rows := xis.GetRows("表字典")
//	tableItem := []string{}
//	for _, row := range rows[1:] {
//		fmt.Println(row)
//		types := strings.ToLower(row[3])
//		if types == "timestamp" {
//			types += " null"
//		}
//		def := "NULL"
//		if row[1] == "COD_ACCT_NO" {
//			def = "'-'"
//			key = strings.ToLower(row[1])
//		}
//		strItem := []string{
//			"`" + strings.ToLower(row[1]) + "`",
//			types,
//			"DEFAULT " + def,
//			"COMMENT " + "'" + row[2] + "'",
//		}
//		if row[6] == tb {
//			tableItem = append(tableItem, strings.Join(strItem, " "))
//		}
//		// 表明（英文）转换为小写
//		//tables = append(tables, g.Map{
//		//	"table_name": strings.ToLower(row[0]),
//		//	"title":      gconv.String(row[1]),
//		//	"type":       strings.ToLower(row[2]),
//		//})
//	}
//	tableItem = append(tableItem, "PRIMARY KEY (`"+key+"`)")
//	tbName, err := pubFun.GetTableName(strings.ToLower(tb))
//	if err == nil {
//		sql := "CREATE TABLE `" + tbName + "` (" + strings.Join(tableItem, ",") + ") ENGINE=InnoDB DEFAULT CHARSET=utf8";
//		fmt.Println(sql)
//	}
//	//_, _ = db_table.InitTableInfo(tables)
//	logs.Info("初始化完成")
//}
//func CreateClassNameByExcel() {
//	tb := "S_LOAN_DK"
//	key := ""
//	log.Instance().Println("根据Excel解析生成表配置")
//	xis, err := excelize.OpenFile("D:/go/src/bankBigData/_static/excel/附件4_下发数据表结构 (自动保存的).xlsx")
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//	log.Instance().Println("Excle文件读取完成")
//	//tables := g.List{}
//	// 获取工作表中sheet1中的值
//	rows := xis.GetRows("表字典")
//	tableItem := []string{}
//	for _, row := range rows[1:] {
//		fmt.Println(row)
//		types := strings.ToLower(row[3])
//		if types == "timestamp" {
//			types += " null"
//		}
//		def := "NULL"
//		if row[1] == "XD_COL167" {
//			def = "'-'"
//			key = strings.ToLower(row[1])
//		}
//		strItem := []string{
//			"`" + strings.ToLower(row[1]) + "`",
//			types,
//			"DEFAULT " + def,
//			"COMMENT " + "'" + row[2] + "'",
//		}
//		if row[6] == tb {
//			tableItem = append(tableItem, strings.Join(strItem, " "))
//		} else {
//			break
//		}
//		// 表明（英文）转换为小写
//		//tables = append(tables, g.Map{
//		//	"table_name": strings.ToLower(row[0]),
//		//	"title":      gconv.String(row[1]),
//		//	"type":       strings.ToLower(row[2]),
//		//})
//	}
//	tableItem = append(tableItem, "PRIMARY KEY (`"+key+"`)")
//	tbName, err := pubFun.GetTableName(strings.ToLower(tb))
//	if err == nil {
//		sql := "CREATE TABLE `" + tbName + "` (" + strings.Join(tableItem, ",") + ") ENGINE=InnoDB DEFAULT CHARSET=utf8";
//		fmt.Println(sql)
//	}
//	//_, _ = db_table.InitTableInfo(tables)
//	logs.Info("初始化完成")
//}

// 写入列
func CreateClassNameByExcel() {
	//tb := "S_LOAN_DK"
	//tb := "S_ECIF_ECIF_CERT_INFO"
	//tb := "S_OFCR_CH_ACCT_MAST"
	log.Instance().Println("根据Excel解析生成表配置")
	xis, err := excelize.OpenFile("D:/go/src/bankBigData/_static/excel/下发数据表结构.xlsx")
	if err != nil {
		logs.Error(err)
		return
	}
	log.Instance().Println("Excle文件读取完成")
	tables := g.List{}
	// 获取工作表中sheet1中的值
	rows := xis.GetRows("表字典")
	for _, row := range rows[1:] {
		text := strings.ToLower(row[2])
		if text == "(null)" {
			text = ""
		}
		tables = append(tables, g.Map{
			"table_name": strings.ToLower(row[6]),
			"column":     strings.ToLower(row[1]),
			"col_type":       strings.ToLower(row[3]),
			"is_key":       strings.Split(row[4],"-")[0],
			"is_null":       strings.Split(row[5],"-")[0],
			"text":       text,
			"table_text": strings.ToLower(row[7]),
			"sys_text":   strings.ToLower(row[8]),
		})
		// 表明（英文）转换为小写
		//tables = append(tables, g.Map{
		//	"table_name": strings.ToLower(row[0]),
		//	"title":      gconv.String(row[1]),
		//	"type":       strings.ToLower(row[2]),
		//})
	}
	//fmt.Println(tables)
	if err == nil {
		_, _ = dbc_tableColumn.Add(tables,"")
	}
	//fmt.Println(tables)
	//_, _ = db_table.InitTableInfo(tables)
	logs.Info("初始化完成")
}
