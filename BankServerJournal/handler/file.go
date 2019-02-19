package handler

import (
	"bankBigData/BankServerJournal/db/file"
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/msg"
	"bankBigData/BankServerJournal/task"
	"bankBigData/_public/util"
	"fmt"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
	"time"
)

type File struct {
	gmvc.Controller
}

func (r *File) Create() {
	startTime := r.Request.GetQueryString("startTime")
	endTime := r.Request.GetQueryString("endTime")
	nowTime := time.Now()
	nowTimeStr := nowTime.Format("20060102_150405")
	files := g.List{
		g.Map{"text": "农民工信息统计表", "type": "info"},
		g.Map{"text": "不良贷款信息统计表", "type": "loan"},
		g.Map{"text": "创业情况统计表", "type": "entre"},
	}
	taskArr := g.List{}

	if !util.DayReg.MatchString(startTime) {
		startTime = "-"
	}
	if !util.DayReg.MatchString(endTime) {
		endTime = "-"
	}
	for i, v := range files {
		taskArr = append(taskArr, g.Map{
			"file_name":  fmt.Sprintf("%v_%v", nowTimeStr, v["text"]),
			"file_path":  fmt.Sprintf("%v_%v_%v.xlsx", gconv.String(i+1), gconv.String(nowTime.Unix()), v["text"]),
			"start_time": startTime,
			"end_time":   endTime,
			"type":       v["type"],
		})
	}
	err := db_file.Add(taskArr)
	if err == nil {
		go task.RunTask()
	}
	_ = r.Response.WriteJson(msg.WriteJson(nil, msg.CreateStr, err != nil))
}

func (r *File) List() {
	res := []entity.TaskFileItem{}
	list, err := db_file.List()
	if err == nil {
		for _, v := range list {
			item := entity.TaskFileItem{}
			if ok := gconv.Struct(v, &item); ok == nil {
				res = append(res, item)
			}
		}
	}
	_ = r.Response.WriteJson(msg.WriteJson(res, msg.ListStr, err != nil))
}
