package db_tableTaskTime

import (
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func Add(data g.List, tbStr string) (int, error) {
	db := g.DB()
	tx, err := db.Begin()
	var r sql.Result
	if err == nil {
		r, err = tx.BatchInsert(table.TableTaskTime, data, 10)
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	id, _ := r.LastInsertId()
	return int(id), err
}
