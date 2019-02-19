package entity

type TaskFileItem struct {
	Id          int    `db:"id" json:"id" field:"id"`
	FileName    string `db:"file_name" json:"fileName" field:"file_name"`
	FilePath    string `db:"file_path" json:"filePath" field:"file_path"`
	CreatedTime string `db:"created_time" json:"createdTime" field:"created_time"`
	StartTime   string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime     string `db:"end_time" json:"endTime" field:"end_time"`
	Type        string `db:"type" json:"type" field:"type"`
	State       string `db:"state" json:"state" field:"state"`
}
