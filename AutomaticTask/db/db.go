package db

import (
	"gitee.com/johng/gf/g"
)

// 增量更新
func AddUpdate(tbName string, data g.List) error {
	db := g.DB()
	_, e := db.BatchSave(tbName, data, 20)
	return e
}

// 全量更新
func AllUpdate(tbName string, data g.List) error {
	db := g.DB()
	_, e := db.BatchReplace(tbName, data, 20)
	return e
}
