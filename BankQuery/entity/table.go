package entity

type TableItem struct {
	Id        int    `db:"id" json:"id" field:"id"`
	TableName string `db:"table_name" json:"tableName" field:"table_name"`
	Title     string `db:"title" json:"title" field:"title"`
	Type      string `db:"type" json:"type" field:"type"`
}
