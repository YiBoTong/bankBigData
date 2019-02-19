package entity

import "gitee.com/johng/gf/g"

type LoanKeys struct {
	Index int    `db:"index" json:"index" field:"index"`
	Key   string `db:"key" json:"key" field:"key"`
	Text  string `db:"text" json:"text" field:"text"`
}

type LoanTableItem g.Map

type Loan struct {
	Keys  []LoanKeys      `db:"keys" json:"keys" field:"keys"`
	Table []LoanTableItem `db:"table" json:"table" field:"table"`
}

type LoanInfo struct {
	XdCol4  string `db:"xd_col4" json:"xdCol4" field:"xd_col4"`    // 贷款日期
	XdCol5  string `db:"xd_col5" json:"xdCol5" field:"xd_col5"`    // 到期日期
	XdCol6  string `db:"xd_col6" json:"xdCol6" field:"xd_col6"`    // 贷款金额
	XdCol7  string `db:"xd_col7" json:"xdCol7" field:"xd_col7"`    // 结欠金额
	XdCol18 string `db:"xd_col18" json:"xdCol18" field:"xd_col18"` // 贷款用途
}
