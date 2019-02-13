package table

const (
	CDbName  = "config"      // 数据库配置分组别名
	IDBName  = "information" // 数据库信息分组别名
	CRDBName = "create"      // 数据库创建分组别名

	CTable             = "tables"               // 配置表
	CTableColumnConfig = "tables_column_config" // 表列表
	CTaskTime          = "task_time"           // 任务表
	CTaskFile          = "task_file"           // 任务文件表
)

var DbDefaultName string = ""
