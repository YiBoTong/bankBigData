package db_query_loan

// 证件信息
import (
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/table"
	"gitee.com/johng/gf/g"
)

func GetUserInfoByIdCard(idCard string) (entity.S_ecif_ecif_cert_info, error) {
	db := g.DB()
	sql := db.Table(table.SEcifEcifCertInfo).Where(g.Map{"cert_num": idCard})
	data := entity.S_ecif_ecif_cert_info{}
	r, err := sql.One()
	_ = r.ToStruct(&data)
	return data, err
}
