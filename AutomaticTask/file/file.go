package file

import (
	"bankBigData/_public/log"
	"gitee.com/johng/gf/g"
	"io/ioutil"
)

var rootFile string

func init() {
	rootFile = g.Config().GetString("filePath")
}

// 读取指定文件夹下的文件夹并以字符串数组的形式返回
func GetDateByFolder() ([]string, error) {
	var dateFolder []string
	log.Instance().Println(rootFile)
	files, err := ioutil.ReadDir(rootFile)
	if err != nil {
		log.Instance().Error(err)
		return dateFolder, err
	}
	for _, file := range files {
		if file.IsDir() {
			dateFolder = append(dateFolder, file.Name())
		}
	}
	return dateFolder, nil
}
