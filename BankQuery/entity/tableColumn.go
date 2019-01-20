package entity

type TableColumnItem struct {
	Id        int    `db:"id" json:"id" field:"id"`
	TableName string `db:"table_name" json:"tableName" field:"table_name"`
	Column    string `db:"column" json:"column" field:"column"`
}
