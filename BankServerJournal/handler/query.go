package handler

import (
	"bankBigData/AutomaticTask/entity/config"
	"bankBigData/BankServerJournal/db/query"
	"bankBigData/BankServerJournal/entity"
	"bankBigData/BankServerJournal/msg"
	"bankBigData/BankServerJournal/redis"
	"bankBigData/BankServerJournal/table"
	"encoding/json"
	"gitee.com/johng/gf/g"
	"gitee.com/johng/gf/g/frame/gmvc"
	"gitee.com/johng/gf/g/util/gconv"
	"regexp"
	"strings"
)

type Query struct {
	gmvc.Controller
}

var cacheLoanKeys []entity.LoanKeys
var cacheQueryKeys []string

// 身份证正则
var idCardReg = regexp.MustCompile(`(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}[0-9Xx]$)`)

func getLoanData(idCard string) ([]entity.LoanKeys, []entity.LoanTableItem, error) {
	queryKeys := []string{}
	loanKeys := []entity.LoanKeys{}
	loanTable := []entity.LoanTableItem{}
	loanTableList := g.List{}
	err := error(nil)

	if len(cacheLoanKeys) == 0 {
		loanKeys, queryKeys, err = getLoanKeys()
		cacheLoanKeys = loanKeys
		cacheQueryKeys = queryKeys
	} else {
		loanKeys = cacheLoanKeys
		queryKeys = cacheQueryKeys
	}

	if err == nil {
		loanTableList, err = db_query_loan.QueryByIdCard(idCard, strings.Join(queryKeys, ","))
	}
	if err == nil && len(loanTableList) > 0 {
		for _, v := range loanTableList {
			loanTable = append(loanTable, v)
		}
	}
	loanTableList = nil
	return loanKeys, loanTable, err
}

func getLoanKeys() ([]entity.LoanKeys, []string, error) {
	loanKeys := []entity.LoanKeys{}
	queryKeys := []string{
		"xd_col4", "xd_col5", "xd_col6", "xd_col7", "xd_col8", "xd_col9",
		/*"xd_col17", "xd_col18", "xd_col19", "xd_col21","xd_col22", */
		"xd_col26", "xd_col121",
	}
	queryKeysSlice := g.Slice{}
	for _, v := range queryKeys {
		queryKeysSlice = append(queryKeysSlice, v)
	}
	loanKeyList, err := db_query_loan.GetKeys(queryKeysSlice)
	for i, v := range loanKeyList {
		item := c_entity.TableColumnConfig{}
		if ok := gconv.Struct(v, &item); ok == nil {
			loanKeys = append(loanKeys, entity.LoanKeys{
				Index: i + 1,
				Key:   item.Column,
				Text:  item.Text,
			})
		}
	}
	return loanKeys, queryKeys, err
}

func (r *Query) Loan() {
	msgStr := ""
	err := error(nil)
	rspData := entity.Loan{}
	loanKeys := []entity.LoanKeys{}
	loanTable := []entity.LoanTableItem{}
	idCard := r.Request.GetQueryString("idCard")
	hasCache := true
	cacheKey := ""
	var cacheData interface{}
	if idCardReg.MatchString(idCard) {
		cacheKey = table.Loan + idCard
		cacheData, _ = redis.Get(cacheKey)
		if cacheData != nil {
			_ = json.Unmarshal(gconv.Bytes(cacheData), &rspData)
		} else {
			hasCache = false
			loanKeys, loanTable, err = getLoanData(idCard)
		}
	} else {
		msgStr = msg.IdCardNumErr
	}
	if !hasCache {
		rspData = entity.Loan{
			Keys:  loanKeys,
			Table: loanTable,
		}
		cacheVal, _ := json.Marshal(rspData)
		_ = redis.Set(cacheKey, cacheVal)
	}
	_ = r.Response.WriteJson(msg.WriteJson(rspData, msg.ListStr, err != nil, msgStr))
}
