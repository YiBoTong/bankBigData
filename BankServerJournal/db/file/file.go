package db_file

import (
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/table"
	"gitee.com/johng/gf/g"
)

func List() (g.List, error) {
	db := g.DB(table.PSDBName)
	res, err := db.Table(table.ComTaskFile).OrderBy("id desc").Limit(0, 30).All()
	return res.ToList(), err
}

func Add(data g.List) error {
	db := g.DB(table.PSDBName)
	_, err := db.BatchInsert(table.ComTaskFile, data, 3)
	return err
}

func GetFirstTask() (entity.TaskFileItem, error) {
	db := g.DB(table.PSDBName)
	data := entity.TaskFileItem{}
	res, err := db.Table(table.ComTaskFile).Where(g.Map{"state": "doing"}).OrderBy("id asc").One()
	_ = res.ToStruct(&data)
	return data, err
}

func Update(id int, data g.Map) error {
	db := g.DB(table.PSDBName)
	_, err := db.Table(table.ComTaskFile).Where(g.Map{"id": id}).Data(data).Update()
	return err
}
