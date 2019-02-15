package c_entity

type TableTaskFile struct {
	Id          int    `db:"id" json:"id" field:"id"`
	Date        string `db:"date" json:"date" field:"date"`
	TableName   string `db:"table_name" json:"tableName" field:"table_name"`
	FileName    string `db:"file_name" json:"fileName" field:"file_name"`
	ContentType int8   `db:"content_type" json:"contentType" field:"content_type"`
	Types       string `db:"types" json:"types" field:"types"`
	IsDown      bool   `db:"is_down" json:"isDown" field:"is_down"`
	IsGz        bool   `db:"is_gz" json:"isGz" field:"is_gz"`
	GzDo        bool   `db:"gz_do" json:"gzDo" field:"gz_do"`
	DbDo        bool   `db:"db_do" json:"dbDo" field:"db_do"`
	LineNum     int    `db:"line_num" json:"lineNum" field:"line_num"`
}
