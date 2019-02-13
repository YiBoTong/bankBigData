package dbc_tableColumn

import (
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func Add(tables g.List, tbName string) (int, error) {
	db := g.DB(table.CDbName)
	tx, err := db.Begin()
	var r sql.Result
	if err == nil {
		//_, _ = tx.Table(table.CTable).Where("table_name=?", tbName).Delete()
		r, err = tx.BatchInsert(table.CTableColumnConfig, tables, 10)
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	id, _ := r.LastInsertId()
	return int(id), err
}

func GetColByTableName(tbName string) (g.List, error) {
	db := g.DB(table.CDbName)
	model := db.Table(table.CTableColumnConfig).Where("table_name=?", tbName)
	r, err := model.OrderBy("id asc").All()
	return r.ToList(), err
}
