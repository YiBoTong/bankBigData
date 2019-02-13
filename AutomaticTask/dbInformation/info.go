package dbi_info

import (
	"bankBigData/_public/table"
	"gitee.com/johng/gf/g"
)

const tbName = "TABLES"

func DbNum(dbName string) (int, error) {
	db := g.DB(table.IDBName)
	sql := db.Table(tbName).Where(g.Map{"TABLE_SCHEMA": dbName})
	r, err := sql.Count()
	return r, err
}

func HasTable(dbName, tableName string) (bool, error) {
	db := g.DB(table.IDBName)
	sql := db.Table(tbName).Where(g.Map{"TABLE_SCHEMA": dbName, "TABLE_NAME": tableName})
	r, err := sql.Count()
	return r > 0, err
}
