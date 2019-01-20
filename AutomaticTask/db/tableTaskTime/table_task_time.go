package db_tableTaskTime

import (
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func Add(data g.List) (int, error) {
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

func Count(where ...g.Map) (int, error) {
	db := g.DB()
	sql := db.Table(table.TableTaskTime)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Count()
	return r, err
}
