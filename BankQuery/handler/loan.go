package handler

import (
	"bankBigData/_public/app"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g/frame/gmvc"
)

type Loan struct {
	gmvc.Controller
}

func (r *Loan) Loan() {
	util.CreateClassNameByExcel()
	res := []string{}
	_ = r.Response.WriteJson(app.Response{
		Data: res,
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   "获取成功",
		},
	})
}
