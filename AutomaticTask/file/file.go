package file

import (
	"gitee.com/johng/gf/g"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

var rootFile = g.Config().GetString("agencyCode")

// 读取指定文件夹下的文件夹并以字符串数组的形式返回
func GetDateByFolder() ([]string, error) {
	var dateFolder []string
	files, err := ioutil.ReadDir(rootFile)
	if err != nil {
		logs.Error("error", err)
		return dateFolder, err
	}
	for _, file := range files {
		if file.IsDir() {
			dateFolder = append(dateFolder, file.Name())
		}
	}
	return dateFolder, nil
}
