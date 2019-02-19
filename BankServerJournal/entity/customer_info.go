package entity

// 用户基本信息
type Customer_info struct {
	Id              int    `db:"id" json:"id" field:"id"`                                            // 客户id
	CustomerId      int    `db:"customer_id" json:"customerId" field:"customer_id"`                  // 客户id
	ManagerId       int    `db:"manager_id" json:"managerId" field:"manager_id"`                     // 业务经理
	DepartmentId    int    `db:"department_id" json:"departmentId" field:"department_id"`            // 网点id
	PhoneNumber     string `db:"phone_number" json:"phoneNumber" field:"phone_number"`               // 手机号
	Name            string `db:"name" json:"name" field:"name"`                                      // 姓名
	Gender          string `db:"gender" json:"gender" field:"gender"`                                // 性别
	IdentityCard    string `db:"identity_card" json:"identityCard" field:"identity_card"`            // 身份证号
	Province        string `db:"province" json:"province" field:"province"`                          // 省
	City            string `db:"city" json:"city" field:"city"`                                      // 市
	County          string `db:"county" json:"county" field:"county"`                                // 区县
	Town            string `db:"town" json:"town" field:"town"`                                      // 镇
	Village         string `db:"village" json:"village" field:"village"`                             // 村
	AnnualIncome    int    `db:"annual_income" json:"annualIncome" field:"annual_income"`            // 年收入
	SourceOfIncome  string `db:"source_of_income" json:"sourceOfIncome" field:"source_of_income"`    // 收入来源
	IsPartyMember   int    `db:"is_party_member" json:"isPartyMember" field:"is_party_member"`       // 是否党员（-1：否，1：是）
	PartyMemberTime int    `db:"party_member_time" json:"partyMemberTime" field:"party_member_time"` // 入党时间
	PartyMemberAddr string `db:"party_member_addr" json:"partyMemberAddr" field:"party_member_addr"` // 组织关系所在地
}

// 务工创业信息
type Customer_info_job struct {
	CustomerId    int    `db:"customer_id" json:"customerId" field:"customer_id"`          // 客户id
	ManagerId     int    `db:"manager_id" json:"managerId" field:"manager_id"`             // 业务经理id
	Province      string `db:"province" json:"province" field:"province"`                  // 省
	City          string `db:"city" json:"city" field:"city"`                              // 市
	County        string `db:"county" json:"county" field:"county"`                        // 区县
	Town          string `db:"town" json:"town" field:"town"`                              // 镇
	Village       string `db:"village" json:"village" field:"village"`                     // 村
	Supplementary string `db:"supplementary" json:"supplementary" field:"supplementary"`   // 补充地址
	JobContent    string `db:"job_content" json:"jobContent" field:"job_content"`          // 工作内容
	Industry      string `db:"industry" json:"industry" field:"industry"`                  // 从事行业
	CorporateName string `db:"corporate_name" json:"corporateName" field:"corporate_name"` // 公司或者厂名
}

// 评级授信信息
type Pt_customer_info_filing struct {
	IsFiling     int    `db:"is_filing" json:"isFiling" field:"is_filing"`             // 是否建档（-1：否，1：是）
	IsRate       int    `db:"is_rate" json:"isRate" field:"is_rate"`                   // 是否评级（-1：否，1：是）
	CreditRating string `db:"credit_rating" json:"creditRating" field:"credit_rating"` // 授信等级
	CreditLimit  int    `db:"credit_limit" json:"creditLimit" field:"credit_limit"`    // 授信额度
}

// 网点信息表
type Org_department_info struct {
	Id             int    `db:"id" json:"id" field:"id"`                                       // 网点id
	DepartmentName string `db:"department_name" json:"departmentName" field:"department_name"` // 部门名称
	TelNumber      string `db:"tel_number" json:"telNumber" field:"tel_number"`                // 座机号
	ParentId       int    `db:"parent_id" json:"parentId" field:"parent_id"`
}
