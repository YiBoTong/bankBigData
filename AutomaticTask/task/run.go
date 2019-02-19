package task

import (
	"bankBigData/AutomaticTask/dbConfig/tableTaskFile"
	"bankBigData/AutomaticTask/dbConfig/tableTaskTime"
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/_public/ftp"
	"bankBigData/_public/log"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"time"
)

// 当前服务的索引（0-10）
// 为0时表示当前服务处理所有的数据
// 为大于0的时候表示的是处理区间的数据
var ServerId = -1

// 当前服务处理的区间步长
var ServerSize = 100

func Run() error {
	log.Instance().Println("开始执行任务")
	lastTask, e := dbc_tableTaskTime.Last("asc", g.Slice{0, 1}, g.Map{})
	if lastTask.Id != 0 {
		// 任务文件已经下载并且解压
		if lastTask.State < 2 {
			e = runTask(lastTask)
		} else if lastTask.State == 2 {
			return nil
		}
		// 1分钟后重新进行任务
		time.Sleep(time.Duration(60) * time.Second)
		_ = Run()
	}
	return e
}

func runTask(task c_entity.TableTaskTime) error {
	_, e := dbc_tableTaskTime.Update(task.Date, g.Map{"start_time": util.GetLocalNowTimeStr()}, g.Map{})
	hasErr := true
	// 只有serverId为0或者1的服务才能进行文件下载
	if ServerId < 2 && task.State == 0 {
		log.Instance().Println("ftp下载文件")
		hasErr = loadFile(task)
	} else {
		hasErr = false
	}
	if !hasErr {
		if task.State == 1 {
			log.Instance().Println("开始分析")
			hasErr = analysis(task)
		} else {
			log.Instance().Println("重新运行分析任务：RUN")
			// 1分钟后重新进行任务
			time.Sleep(time.Duration(60) * time.Second)
			_ = Run()
		}
	}
	if !hasErr {
		_, e = dbc_tableTaskTime.Update(task.Date, g.Map{"end_time": util.GetLocalNowTimeStr(), "state": 2}, g.Map{})
	}
	return e
}

func loadFile(task c_entity.TableTaskTime) bool {
	buf := make(chan []string)
	flg := make(chan int)
	err := make(chan bool)
	hasErr := true
	go downloadAndGzTaskFile(task, buf)
	go gzTaskFileByDate(task.Date, buf, flg, err)
	hasErr, _ = <-err
	<-flg
	return hasErr
}

func downloadAndGzTaskFile(task c_entity.TableTaskTime, buf chan []string) {
	defer close(buf)
	taskFiles, e := dbc_tableTaskFile.Gets(task.Date, g.Slice{}, g.Map{"is_down": 0})
	if e == nil && len(taskFiles) > 0 {
		log.Instance().Println("开始下载 " + task.Date + " 的文件")
		for _, v := range taskFiles {
			item := c_entity.TableTaskFile{}
			ok := gconv.Struct(v, &item)
			filePath := "/" + task.Date + "/" + item.FileName
			if ok == nil {
				e = pub_ftp.DownloadFile(filePath)
			}
			if e != nil {
				break
			}
			_, _ = dbc_tableTaskFile.Update(task.Date, item.FileName, g.Map{"is_down": 1}, g.Map{"is_down": 0})
			if e == nil && item.IsGz {
				//e = gzTaskFileByDate(task.Date, item.FileName, filePath)
				buf <- []string{item.FileName, filePath}
			}
		}
		log.Instance().Println("" + task.Date + " 的文件下载完成")
		if e == nil {
			_, e = dbc_tableTaskTime.Update(task.Date, g.Map{"end_time": util.GetLocalNowTimeStr(), "state": 1}, g.Map{})
		}
	}
}

func gzTaskFileByDate(date string, buf chan []string, flg chan int, err chan bool) {
	e := error(nil)
	for {
		if v, ok := <-buf; ok && len(v) == 2 {
			hasFile := false
			theFilePath := pub_ftp.FtpFolder + v[1]
			hasFile, e = util.PathExists(theFilePath)
			if e == nil && hasFile {
				_, e = util.Gzip(theFilePath, "")
				if e == nil {
					_, _ = dbc_tableTaskFile.Update(date, v[0], g.Map{"gz_do": 1}, g.Map{"gz_do": 0})
				}
			}
		} else {
			break
		}
	}
	err <- (e != nil)
	flg <- 1
}
