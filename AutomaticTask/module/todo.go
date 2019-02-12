package module

import (
	"bankBigData/AutomaticTask/db/tableTaskTime"
	"bankBigData/AutomaticTask/file"
	"fmt"
	"gitee.com/johng/gf/g"
	"time"
)

func AutoLoadData() {
	//SaveDayFolder()
	fmt.Println(time.Now(), "--> 123")
}

func SaveDayFolder() {
	dayFolderArr := g.List{}
	dayNum, err := db_tableTaskTime.Count()
	if dayNum == 0 && err == nil {
		// 获取日期文件夹
		dayFolders, err := file.GetDateByFolder()
		if len(dayFolders) > 0 && err == nil {
			for _, v := range dayFolders {
				dayFolderArr = append(dayFolderArr, g.Map{
					"date": v,
				})
			}
		}
	}
	dayFolderArr = append(dayFolderArr, g.Map{
		"date": time.Now().Format("20060102"),
	})
	fmt.Println(dayFolderArr)
}
