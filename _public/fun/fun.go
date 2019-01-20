package pubFun

import (
	"bankBigData/BankQuery/db/table"
	"gitee.com/johng/gf/g"
)

func GetTableName(tableStr string) (string, error) {
	tbName := ""
	tbInfo, err := db_table.GetInfoByTableName(tableStr)
	if err == nil && tbInfo.Id != 0 {
		tbName = tbInfo.TableName + "_" + g.Config().GetString("agencyCode")
	}
	return tbName, err
}
