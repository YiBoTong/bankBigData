package msg

import "bankBigData/_public/app"

func WriteJson(data interface{}, msgType string, hasErr bool, msgStr ...string) interface{} {
	code := 200
	msgTip := ""
	if len(msgStr) > 0 {
		msgTip = msgStr[0]
	}
	if msgTip == "" {
		msgTip = GetTodoResMsg(msgType, hasErr)
	}
	if hasErr {
		code = 240
	}
	return app.Response{
		Data: data,
		Code: code,
		Msg:  msgTip,
	}
}
