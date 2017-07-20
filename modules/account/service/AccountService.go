package service

import (
	"com.cxria/base"
	"com.cxria/modules/account/dao"
	"com.cxria/modules/code"
)

func GetAccount(accountId int64) base.Json {
	json := base.GetJson()
	json.Content = GetAccount(accountId)
	json.Ok = 1
	return json
}

func verifySession(token string) base.Json {
	json := base.GetJson()
	session := dao.GetSessionByToken(token)
	if &session != nil {
		json.Ok = base.SUCCESS
		json.Content = session
	} else {
		json.ErrorCode = code.ErrorCode["NO_SESSION"]
		json.Message = code.Message["NO_SESSION"]
	}
	return json
}
