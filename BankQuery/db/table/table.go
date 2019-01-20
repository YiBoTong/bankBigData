package db_table

import (
	"bankBigData/BankQuery/entity"
	"bankBigData/_public/table"
	"database/sql"
	"gitee.com/johng/gf/g"
)

func AddTableInfo(tables g.List) (int, error) {
	db := g.DB()
	tx, err := db.Begin()
	var r sql.Result
	if err == nil {
		_, _ = tx.Table(table.TableConfig).Delete()
		r, err = tx.BatchInsert(table.TableConfig, tables, 10)
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}
	id, _ := r.LastInsertId()
	return int(id), err
}

func GetInfoByTableName(tbName string) (entity.TableItem, error) {
	db := g.DB()
	item := entity.TableItem{}
	sql := db.Table(table.TableConfig).Where("table_name=?", tbName)
	r, err := sql.One()
	_ = r.ToStruct(&item)
	return item, err
}
