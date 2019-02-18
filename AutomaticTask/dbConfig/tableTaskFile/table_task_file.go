package dbc_tableTaskFile

import (
	"bankBigData/AutomaticTask/dbConfig/tableTaskTime"
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/database/gdb"
	"strings"
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

func Gets(date string, types g.Slice, where ...g.Map) (g.List, error) {
	db := g.DB(table.CDbName)
	sql := db.Table(table.CTaskFile).Where(g.Map{"date": date})
	if len(where) > 0 {
		sql.And(where[0])
	}
	typeLen := len(types)
	if typeLen > 0 {
		t := []string{}
		for i := 1; i < typeLen-1; i++ {
			t = append(t, "?")
		}
		sql.And("`types` IN("+strings.Join(t, ",")+")", types)
	}
	r, err := sql.OrderBy("id asc").All()
	return r.ToList(), err
}

func Update(date, fileName string, data g.Map, where g.Map, tx ...gdb.TX) (int, error) {
	e := error(nil)
	where["date"] = date
	where["file_name"] = fileName
	if len(tx) > 0 {
		_, e = tx[0].Table(table.CTaskFile).Where(where).Data(data).Update()
	} else {
		db := g.DB(table.CDbName)
		_, e = db.Table(table.CTaskFile).Where(where).Data(data).Update()
	}
	return 1, e
}
