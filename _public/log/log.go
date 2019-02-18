package log

import (
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/os/glog"
)

var log *glog.Logger

func Init(nameSpace string) {
	log = glog.New()
	log.SetPath(g.Config().GetString("logPath") + "/" + nameSpace)
}

func Instance() *glog.Logger {
	return log
}
