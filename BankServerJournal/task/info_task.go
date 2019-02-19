package task

import (
	"bankBigData/BankServerJournal/db/query"
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/excel"
	"bankBigData/_public/util"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/util/gconv"
	"github.com/360EntSecGroup-Skylar/excelize"
)

var iStartPage = 2

const infoSheetName = "Sheet1" //
const infoRowNum = 35          // 一行要写入的字段数
var infoColNum = 6             // 起始列

func PageUser(page int) g.List {
	limit := 100
	offset := (page - 1) * limit
	if page > 1 {
		offset += 1
	}
	allUser, _ := db_query_loan.QueryPageUserInfo_Pt(offset, limit)
	return allUser
}

func writeInfoFile(xlsx *excelize.File, pageUser g.List) bool {
	if len(pageUser) < 1 {
		return true
	}
	for _, v := range pageUser {
		item := entity.Customer_info{}
		if ok := gconv.Struct(v, &item); ok == nil {
			iWriteInfoFileLine(item, xlsx)
		}
	}
	iStartPage += 1
	pageUser = PageUser(iStartPage)
	return writeInfoFile(xlsx, pageUser)
}

func iWriteInfoFileLine(userInfo entity.Customer_info, xlsx *excelize.File) {
	lineData := [infoRowNum]interface{}{}
	tempUserInfo, _ := db_query_loan.GetUserInfo_buss(userInfo.Id)
	// 业务经理录入的信息覆盖用户自己的信息
	if tempUserInfo.CustomerId != 0 {
		userInfo = tempUserInfo
		userInfo.Id = userInfo.CustomerId
	}
	// 序号
	lineData[0] = (infoColNum - 6) + 1
	iLoadUserInfoToLineArr(&lineData, userInfo)
	iLoadUserInfoJobToLineArr(&lineData, userInfo)
	iLoadUserLoanToLineArr(&lineData, userInfo)
	iLoadUserBankToLineArr(&lineData, userInfo)
	iLoadUserFilingToLineArr(&lineData, userInfo)

	for y := 0; y < infoRowNum; y++ {
		excel.ModifyExcelCellByAxis(xlsx, infoSheetName, excel.ChangIndexToAxis(infoColNum, y), lineData[y])
	}
	infoColNum += 1
}

func iLoadUserInfoToLineArr(lineData *[infoRowNum]interface{}, userInfo entity.Customer_info) {
	// 基本信息
	lineData[1] = userInfo.Name
	lineData[2] = userInfo.Gender
	lineData[4] = userInfo.PhoneNumber
	lineData[5] = userInfo.IdentityCard
	lineData[6] = userInfo.City
	lineData[7] = userInfo.County
	lineData[8] = userInfo.Town
	lineData[9] = userInfo.Village
	// 政治面貌
	IsPartyMember := ""
	if userInfo.IsPartyMember == -1 {
		IsPartyMember = "否"
	} else if userInfo.IsPartyMember == 1 {
		IsPartyMember = "是"
	}
	PartyMemberTime := ""
	if userInfo.PartyMemberTime != 0 {
		// todo 格式化日期
	}
	lineData[18] = IsPartyMember
	lineData[19] = PartyMemberTime
	lineData[20] = userInfo.PartyMemberAddr
}

// 务工
func iLoadUserInfoJobToLineArr(lineData *[infoRowNum]interface{}, userInfo entity.Customer_info) {
	newData, _ := db_query_loan.GetUserInfo_buss_job(userInfo.Id)
	if newData.CustomerId == 0 {
		newData, _ = db_query_loan.GetUserInfoJob(userInfo.Id)
	}
	lineData[10] = newData.Province
	lineData[11] = newData.City
	lineData[12] = newData.County
	lineData[13] = newData.Town
	lineData[14] = newData.Village
	lineData[15] = ""
	lineData[16] = newData.Industry
	lineData[17] = newData.CorporateName
}

// 辖内农信社信息
func iLoadUserBankToLineArr(lineData *[infoRowNum]interface{}, userInfo entity.Customer_info) {
	data, _ := db_query_loan.GetOrgDepartmentInfo(userInfo.Id)
	sonBank := data.DepartmentName
	parentBank := data.DepartmentName
	if data.ParentId == 0 {
		sonBank = ""
	} else {
		parentBank = ""
	}
	lineData[30] = parentBank
	lineData[31] = sonBank
	lineData[32] = data.TelNumber
}

// 建档、评级信息
func iLoadUserFilingToLineArr(lineData *[infoRowNum]interface{}, userInfo entity.Customer_info) {
	data, _ := db_query_loan.GetPtCustomerInfoFiling(userInfo.Id)
	is_filing := ""
	is_rate := ""
	if data.IsFiling == -1 {
		is_filing = "否"
	} else if data.IsFiling == 1 {
		is_filing = "是"
	}
	if data.IsRate == -1 {
		is_rate = "否"
	} else if data.IsRate == 1 {
		is_rate = "是"
	}
	lineData[21] = is_filing
	lineData[22] = is_rate
	lineData[23] = data.CreditRating
	lineData[24] = data.CreditLimit
}

// 贷款
func iLoadUserLoanToLineArr(lineData *[infoRowNum]interface{}, userInfo entity.Customer_info) {
	data := entity.LoanInfo{}
	if util.DayReg.MatchString(userInfo.IdentityCard) {
		data, _ = db_query_loan.QueryByIdCardLast(userInfo.IdentityCard, "xd_col4,xd_col5,xd_col6,xd_col7,xd_col18")
	}
	lineData[25] = data.XdCol4
	lineData[26] = data.XdCol5
	lineData[27] = data.XdCol6
	lineData[28] = data.XdCol7
	lineData[29] = data.XdCol18
}
