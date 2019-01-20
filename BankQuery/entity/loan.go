package entity

type QueryLoanListItem struct {
	// 贷款日期
	XdCol4 int `db:"xd_col4" json:"xd_col4" field:"xd_col4"`
	// 到期日期
	XdCol5 int `db:"xd_col5" json:"xd_col5" field:"xd_col5"`
	// 贷款金额
	XdCol6 int `db:"xd_col6" json:"xd_col6" field:"xd_col6"`
	// 结欠金额
	XdCol7 int `db:"xd_col7" json:"xd_col7" field:"xd_col7"`
}

