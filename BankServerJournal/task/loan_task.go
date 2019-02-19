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

var lStartPage = 2

const loanSheetName = "Sheet1" //
const loanRowNum = 28          // 一行要写入的字段数
var loanColNum = 6             // 起始列

func writeLoanFile(xlsx *excelize.File, pageUser g.List) bool {
	if len(pageUser) < 1 {
		return true
	}
	for _, v := range pageUser {
		item := entity.Customer_info{}
		if ok := gconv.Struct(v, &item); ok == nil {
			lWriteloanFileLine(item, xlsx)
		}
	}
	lStartPage += 1
	pageUser = PageUser(lStartPage)
	return writeLoanFile(xlsx, pageUser)
}

func lWriteloanFileLine(userInfo entity.Customer_info, xlsx *excelize.File) {
	lineData := [loanRowNum]interface{}{}
	tempUserInfo, _ := db_query_loan.GetUserInfo_buss(userInfo.Id)
	// 业务经理录入的信息覆盖用户自己的信息
	if tempUserInfo.CustomerId != 0 {
		userInfo = tempUserInfo
		userInfo.Id = userInfo.CustomerId
	}
	// 序号
	lineData[0] = (loanColNum - 6) + 1
	lLoadUserloanLineArr(&lineData, userInfo)
	lLoadUserLoanJobToLineArr(&lineData, userInfo)
	lLoadUserLoansToLineArr(&lineData, userInfo)
	lLoadUserLoanBankToLineArr(&lineData, userInfo)

	for y := 0; y < loanRowNum; y++ {
		excel.ModifyExcelCellByAxis(xlsx, infoSheetName, excel.ChangIndexToAxis(loanColNum, y), lineData[y])
	}
	loanColNum += 1
}

func lLoadUserloanLineArr(lineData *[loanRowNum]interface{}, userInfo entity.Customer_info) {
	// 基本信息
	lineData[1] = userInfo.Name
	lineData[2] = userInfo.Gender
	lineData[4] = userInfo.PhoneNumber
	lineData[5] = userInfo.IdentityCard
	lineData[6] = userInfo.City
	lineData[7] = userInfo.County
	lineData[8] = userInfo.Town
	lineData[9] = userInfo.Village
}

// 务工
func lLoadUserLoanJobToLineArr(lineData *[loanRowNum]interface{}, userInfo entity.Customer_info) {
	newData, _ := db_query_loan.GetUserInfo_buss_job(userInfo.Id)
	if newData.CustomerId == 0 {
		newData, _ = db_query_loan.GetUserInfoJob(userInfo.Id)
	}
	lineData[16] = newData.Province
	lineData[17] = newData.City
	lineData[18] = newData.County
	lineData[19] = newData.Town
	lineData[20] = newData.Village
	lineData[21] = ""
	lineData[22] = newData.Industry
	lineData[23] = newData.CorporateName
}

// 辖内农信社信息
func lLoadUserLoanBankToLineArr(lineData *[loanRowNum]interface{}, userInfo entity.Customer_info) {
	data, _ := db_query_loan.GetOrgDepartmentInfo(userInfo.Id)
	sonBank := data.DepartmentName
	parentBank := data.DepartmentName
	if data.ParentId == 0 {
		sonBank = ""
	} else {
		parentBank = ""
	}
	lineData[24] = parentBank
	lineData[25] = sonBank
	lineData[26] = data.TelNumber
}

// 贷款
func lLoadUserLoansToLineArr(lineData *[loanRowNum]interface{}, userInfo entity.Customer_info) {
	data := entity.LoanInfo{}
	if util.DayReg.MatchString(userInfo.IdentityCard) {
		data, _ = db_query_loan.QueryByIdCardLast(userInfo.IdentityCard, "xd_col4,xd_col5,xd_col6,xd_col7,xd_col18")
	}
	lineData[10] = data.XdCol4
	lineData[11] = data.XdCol5
	lineData[12] = data.XdCol6
	lineData[13] = data.XdCol7
}
