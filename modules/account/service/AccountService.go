package service

import (
	"com.cxria/base"
	"com.cxria/modules/account/dao"
	"com.cxria/modules/code"
	"time"
	"net/http"
	"com.cxria/modules/account/domain"
	"github.com/astaxie/beego/orm"
	"com.cxria/utils/key"
)

func GetAccount(accountId int64) base.Json {
	json := base.GetJson()
	json.Content = dao.GetAccount(accountId)
	json.Ok = 1
	return json
}

func VerifySession(token string) base.Json {
	json := base.GetJson()
	session, _ := dao.GetSessionByToken(token)
	if &session != nil {
		json.Ok = base.SUCCESS
		json.Content = session
	} else {
		json.ErrorCode = code.ErrorCode["NO_SESSION"]
		json.Message = code.Message["NO_SESSION"]
	}
	return json
}

func ArchiveSession(token string) base.Json {
	json := base.GetJson()
	session, o := dao.GetSessionByToken(token)
	if &session != nil {
		session.LogoutTime = time.Now()
		o.Update(&session)
	}
	json.Ok = base.SUCCESS
	return json
}

func AddSession(accountId int64, request http.Request, writer http.ResponseWriter) base.Json {
	json := base.GetJson()
	token := key.Generate(32)
	var session = domain.Session{AccountId: accountId, Token: token, UserAgent: request.UserAgent()}
	orm.NewOrm().Insert(session)
	cookie := http.Cookie{Name: "token", Value: token, Path: "/", MaxAge: 15 * 24 * 60 * 60}
	http.SetCookie(writer, &cookie)
	json.Ok = base.SUCCESS
	return json
}

func SendAuthEmail(email string, emailType int, accountId int64) base.Json{
	json := base.GetJson()

	return json
}
