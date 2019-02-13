package dbc_tableTaskFile

import (
	"bankBigData/AutomaticTask/dbConfig/tableTaskTime"
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func Add(date string, data g.List) (int, error) {
	db := g.DB(table.CDbName)
	tx, err := db.Begin()
	var id int64 = 0
	r := sql.Result(nil)
	if err == nil {
		if len(data) > 0 {
			r, err = tx.BatchInsert(table.CTaskFile, data, 10)
			id, _ = r.LastInsertId()
		}
		if err == nil {
			_, err = dbc_tableTaskTime.Update(date, g.Map{"is_scan": 1}, g.Map{}, *tx)
		}
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	return int(id), err
}

func Count(where ...g.Map) (int, error) {
	db := g.DB(table.CDbName)
	sql := db.Table(table.CTaskFile)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Count()
	return r, err
}
