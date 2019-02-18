package dbInit

import (
	"bankBigData/AutomaticTask/dbConfig/table"
	"bankBigData/AutomaticTask/dbConfig/tableColumn"
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/log"
	"bankBigData/_public/table"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"regexp"
	"strings"
	"time"
)

// sql中的关键字
var numReg = regexp.MustCompile(`(\d+)`)

func InitDb() error {
	e := createdDB()
	tables := g.List{}
	// 5秒后进行表初始化
	time.Sleep(time.Duration(5) * time.Second)
	if e == nil {
		tables, e = dbc_table.All()
	}
	tabLen := len(tables)
	if e == nil {
		for i, v := range tables {
			item := c_entity.Table{}
			if ok := gconv.Struct(v, &item); ok == nil {
				e = initTable(item.TableName)
				log.Instance().Println("剩余初始化表数：", tabLen-i-1)
			}
			if e != nil {
				break
			}
		}
	}
	log.Instance().Println("初始化数据库完成")
	return e
}

func initTable(tbName string) error {
	cols, e := dbc_tableColumn.GetColByTableName(tbName)
	if e == nil && len(cols) > 0 {
		e = createdTable(tbName, cols)
	}
	return e
}

func createdDB() error {
	db := g.DB(table.CRDBName)
	_, e := db.Exec("CREATE DATABASE IF NOT EXISTS `" + table.DbDefaultName + "` default charset utf8 COLLATE utf8_general_ci")
	log.Instance().Println("创建数据库完成")
	return e
}

func createdTable(tableName string, col g.List) error {
	db := g.DB()
	keyItem := []string{}
	tableItem := []string{}
	for _, v := range col {
		item := c_entity.TableColumnConfig{}
		if ok := gconv.Struct(v, &item); ok == nil {
			def := "DEFAULT NULL"
			types := item.ColType
			result := numReg.FindAllStringSubmatch(types, -1)
			if len(result) > 0 && gconv.Int(result[0][1]) >= 1024 {
				types = "text(" + result[0][1] + ")"
			}
			if types == "timestamp" {
				types += " null"
			}
			if item.IsKey {
				keyItem = append(keyItem, "`"+item.Column+"`")
			}
			if item.IsKey || !item.IsNull {
				def = "NOT NULL"
			}
			strItem := []string{
				"`" + item.Column + "`",
				types,
				def,
				"COMMENT " + "'" + item.Text + "'",
			}
			tableItem = append(tableItem, strings.Join(strItem, " "))
		}
	}
	if len(keyItem) > 0 {
		tableItem = append(tableItem, "PRIMARY KEY ("+strings.Join(keyItem, ",")+")")
	}
	// MyISAM | InnoDB
	sql := "CREATE TABLE `" + tableName + "` (" + strings.Join(tableItem, ",") + ") ENGINE=InnoDB DEFAULT CHARSET=utf8"
	_, _ = db.Exec("DROP TABLE IF EXISTS `" + tableName + "`")
	_, e := db.Exec(sql)
	if e == nil {
		log.Instance().Println("初始化数据表：", tableName)
	} else {
		log.Instance().Println("初始化数据表失败：", tableName)
	}
	return e
}
