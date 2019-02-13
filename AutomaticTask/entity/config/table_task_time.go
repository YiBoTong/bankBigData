package c_entity

type TableTaskTime struct {
	Id        int    `db:"id" json:"id" field:"id"`
	Date      string `db:"date" json:"date" field:"date"`
	StartTime string `db:"start_time" json:"startTime" field:"start_time"`
	EndTime   string `db:"end_time" json:"endTime" field:"end_time"`
	IsScan    int8   `db:"is_scan" json:"isScan" field:"is_scan"`
	State     int8   `db:"state" json:"state" field:"state"`
}
