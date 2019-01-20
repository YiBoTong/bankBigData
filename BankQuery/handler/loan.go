package handler

import (
	"bankBigData/_public/app"
	"gitee.com/johng/gf/g/frame/gmvc"
)

type Loan struct {
	gmvc.Controller
}

func (r *Loan) Loan() {
	//util.CreateClassNameByExcel()
	_ = r.Response.WriteJson(app.Response{
		Data: 0,
		Status: app.Status{
			Code:  0,
			Error: false,
			Msg:   "获取成功",
		},
	})
}
