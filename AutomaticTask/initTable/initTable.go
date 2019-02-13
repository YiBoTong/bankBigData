package initTable

import (
	"bankBigData/AutomaticTask/dbInformation"
	"bankBigData/_public/table"
)

func HasDb() bool {
	num, _ := dbi_info.DbNum(table.DbDefaultName)
	return num > 0
}

