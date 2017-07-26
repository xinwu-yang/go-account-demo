package service

import (
	"com.cxria/base"
	"com.cxria/modules/account/dao"
	"time"
	"net/http"
	"com.cxria/modules/account/domain"
	"github.com/astaxie/beego/orm"
	"com.cxria/utils/key"
	"com.cxria/api/redis"
	"com.cxria/utils/crypto"
	"encoding/hex"
	"fmt"
	"strconv"
	"com.cxria/utils/mail"
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
		json.SetError("NO_SESSION")
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

func SendAuthEmail(email string, emailType int, accountId int64) base.Json {
	json := base.GetJson()
	if !redis.Exists(email) {
		json.SetError("SMS_TIME_NOT_EX")
		return json
	}
	code := key.Generate(6)
	emailSha := crypto.SHA256Hex(email)
	k := []byte(emailSha[0:32])
	i := []byte(emailSha[32:64])
	hex.Decode(k, k)
	hex.Decode(i, i)
	aesCode, _ := crypto.AesEncrypt([]byte(code), k[:16], i[:16])
	aesCodeStr := hex.EncodeToString(aesCode)
	params := "code=" + aesCodeStr + "&email=" + email + "&type=" + strconv.Itoa(emailType) + "&a=" + strconv.FormatInt(accountId, 10)
	var contents map[string]interface{} = make(map[string]interface{})
	contents["params"] = params
	go func() {
		fmt.Println(params)
		mail.Send(email, params, "邮箱验证码")
		mm, _ := time.ParseDuration("1m")
		verification := domain.Verification{Code: code, Contact: email, Expiry: time.Now().Add(15 * mm)}
		orm.NewOrm().Insert(verification)
	}()
	redis.SetEx(email, 60, "")
	json.Ok = base.SUCCESS
	return json
}

//还未实现发送手机短信

func VerifyCode(code string, contact string) base.Json {
	json := base.GetJson()
	verification := dao.GetVerification(code, contact)
	if &verification == nil {
		json.SetError("EXPIRED_CODE")
		return json
	}
	json.Ok = base.SUCCESS
	return json
}
