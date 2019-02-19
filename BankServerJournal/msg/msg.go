package msg

const (
	ListStr       = "获取列表"
	DelStr        = "删除"
	GetStr        = "获取"
	CreateStr     = "创建"
	SuccessStr    = "成功"
	ErrorStr      = "失败"
	StateStr      = "状态"
	Had           = "已存在"
	NoHad         = "不存在"
	MastHasOneStr = "至少要有一个"
	IdCardNumErr  = "身份证不合法"
)

// 获取操作提示
func GetTodoResMsg(str string, error bool) string {
	if error {
		return str + ErrorStr
	}
	return str + SuccessStr
}
