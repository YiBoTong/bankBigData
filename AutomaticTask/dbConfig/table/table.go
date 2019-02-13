package dbc_table

import (
	"bankBigData/_public/fun"
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func Add(data g.List, tbStr string) (int, error) {
	db := g.DB(table.CDbName)
	tx, err := db.Begin()
	var r sql.Result
	tbName := ""
	if err == nil {
		tbName, err = pubFun.GetTableName(tbStr)
	}
	if err == nil && tbName != "" {
		r, err = tx.BatchInsert(tbName, data, 10)
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	id, _ := r.LastInsertId()
	return int(id), err
}

func All() (g.List, error) {
	db := g.DB(table.CDbName)
	model := db.Table(table.CTable).OrderBy("id asc")
	r, err := model.All()
	return r.ToList(), err
}
