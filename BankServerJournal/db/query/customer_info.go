package db_query_loan

import (
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/table"
	"gitee.com/johng/gf/g"
)

// 系统里的所有用户
func QueryPageUserInfo_Pt(start, limit int) (g.List, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.PtCustomerInfo).Limit(start, limit)
	sql.OrderBy("id asc")
	r, err := sql.All()
	return r.ToList(), err
}

// 获取用户信息
func GetUserInfo(userId int) (entity.Customer_info, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.PtCustomerInfo).Where("id=?", userId)
	r, err := sql.One()
	data := entity.Customer_info{}
	_ = r.ToStruct(&data)
	return data, err
}

// 获取客户信息（业务经理填写）
func GetUserInfo_buss(userId int) (entity.Customer_info, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.BussCustomerInfo).Where("customer_id=?", userId)
	r, err := sql.One()
	data := entity.Customer_info{}
	_ = r.ToStruct(&data)
	return data, err
}

// 获取用户创业信息
func GetUserInfoJob(userId int) (entity.Customer_info_job, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.PtCustomerInfoJob).Where("customer_id=?", userId)
	r, err := sql.One()
	data := entity.Customer_info_job{}
	_ = r.ToStruct(&data)
	return data, err
}

// 获取用户的创业信息（业务经理填写）
func GetUserInfo_buss_job(userId int) (entity.Customer_info_job, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.BussCustomerInfoJob).Where("customer_id=?", userId)
	r, err := sql.One()
	data := entity.Customer_info_job{}
	_ = r.ToStruct(&data)
	return data, err
}

// 用户的评级授信
func GetPtCustomerInfoFiling(userId int) (entity.Pt_customer_info_filing, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.PtCustomerInfoFiling).Where("customer_id=?", userId)
	r, err := sql.One()
	data := entity.Pt_customer_info_filing{}
	_ = r.ToStruct(&data)
	return data, err
}

// 网点信息
func GetOrgDepartmentInfo(id int) (entity.Org_department_info, error) {
	db := g.DB(table.PSDBName)
	sql := db.Table(table.OrgDepartmentInfo).Where("id=?", id)
	r, err := sql.One()
	data := entity.Org_department_info{}
	_ = r.ToStruct(&data)
	return data, err
}
