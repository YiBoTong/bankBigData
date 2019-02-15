package dbc_tableTaskTime

import (
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
)

func Add(data g.List) (int, error) {
	db := g.DB(table.CDbName)
	tx, err := db.Begin()
	var r sql.Result
	if err == nil {
		r, err = tx.BatchInsert(table.CTaskTime, data, 10)
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	id, _ := r.LastInsertId()
	return int(id), err
}

func Update(date string, data g.Map, where g.Map, tx ...gdb.TX) (int, error) {
	res := sql.Result(nil)
	e := error(nil)
	where["date"] = date
	if len(tx) > 0 {
		res, e = tx[0].Table(table.CTaskTime).Where(where).Data(data).Update()
	} else {
		db := g.DB(table.CDbName)
		res, e = db.Table(table.CTaskTime).Where(where).Data(data).Update()
	}
	row, _ := res.RowsAffected()
	return int(row), e
}

func Count(where ...g.Map) (int, error) {
	db := g.DB(table.CDbName)
	sql := db.Table(table.CTaskTime)
	if len(where) > 0 {
		sql.And(where[0])
	}
	r, err := sql.Count()
	return r, err
}

func Last(idOrder string, state g.Slice, where ...g.Map) (c_entity.TableTaskTime, error) {
	db := g.DB(table.CDbName)
	data := c_entity.TableTaskTime{}
	sql := db.Table(table.CTaskTime)
	if len(where) > 0 {
		sql.Where(where[0])
	}
	if len(state) > 0 {
		t := []string{}
		for i := 0; i < len(state)-1; i++ {
			t = append(t, "?")
		}
		if len(where) > 0 {
			sql.And("state IN ("+strings.Join(t, ",")+")", state)
		} else {
			sql.Where("state IN ("+strings.Join(t, ",")+")", state)
		}
	}
	if idOrder == "" {
		idOrder = "desc"
	}
	sql.OrderBy("id " + idOrder)
	sql.Limit(0, 1)
	r, err := sql.One()
	_ = r.ToStruct(&data)
	return data, err
}
