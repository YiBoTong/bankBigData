package db_query_loan

import (
	"bankBigData/BankServerJournal/entity"
	table2 "bankBigData/BankServerJournal/table"
	"bankBigData/_public/table"
	"gitee.com/johng/gf/g"
)

func GetKeys(keyStr g.Slice) (g.List, error) {
	db := g.DB(table.CDbName)
	sql := db.Table(table.CTableColumnConfig).Fields("`id`,`column`,`text`")
	sql.Where(g.Map{"table_name": table2.Loan})
	if len(keyStr) > 0 {
		sql.And("`column` IN (?)", keyStr)
	}
	sql.OrderBy("id asc")
	res, err := sql.All()
	return res.ToList(), err
}

func QueryByIdCard(idCard, queryKeys string) (g.List, error) {
	db := g.DB()
	sql := db.Table(table2.Loan).Where(g.Map{"xd_col16": idCard})
	sql.Fields(queryKeys)
	sql.OrderBy("xd_col28 desc")
	r, err := sql.All()
	return r.ToList(), err
}

func QueryByIdCardLast(idCard, queryKeys string, orderStr ...string) (entity.LoanInfo, error) {
	db := g.DB()
	order := "xd_col5"
	if len(orderStr) > 0 {
		order = orderStr[0]
	}
	sql := db.Table(table2.Loan).Where(g.Map{"xd_col16": idCard})
	sql.Fields(queryKeys)
	// 贷款日期倒序
	sql.OrderBy(order + " desc")
	data := entity.LoanInfo{}
	r, err := sql.One()
	_ = r.ToStruct(&data)
	return data, err
}
