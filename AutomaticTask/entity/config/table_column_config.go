package c_entity

type TableColumnConfig struct {
	Id        int    `db:"id" json:"id" field:"id"`
	TableName string `db:"table_name" json:"tableName" field:"table_name"`
	Column    string `db:"column" json:"column" field:"column"`
	ColType   string `db:"col_type" json:"colType" field:"col_type"`
	IsKey     bool   `db:"is_key" json:"isKey" field:"is_key"`
	IsNull    bool   `db:"is_null" json:"isNull" field:"is_null"`
	Text      string `db:"text" json:"text" field:"text"`
	TableText string `db:"table_text" json:"tableText" field:"table_text"`
	SysText   string `db:"sys_text" json:"sysText" field:"sys_text"`
}
