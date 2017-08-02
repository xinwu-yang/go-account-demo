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
	"strconv"
	"com.cxria/utils/mail"
	"com.cxria/api/sms"
	"golang.org/x/crypto/bcrypt"
	"com.cxria/utils/str"
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
	if session.SessionId != 0 {
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
	orm.NewOrm().Insert(&session)
	cookie := http.Cookie{Name: "token", Value: token, Path: "/", MaxAge: 15 * 24 * 60 * 60}
	http.SetCookie(writer, &cookie)
	json.Ok = base.SUCCESS
	return json
}

func SendAuthEmail(email string, emailType int, accountId int64) base.Json {
	json := base.GetJson()
	if redis.Exists(email) {
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
		mail.Send(email, params, "邮箱验证码")
		mm, _ := time.ParseDuration(sms.OVERDUE + "m")
		verification := domain.Verification{Code: code, Contact: email, Expiry: time.Now().Add(mm)}
		orm.NewOrm().Insert(&verification)
	}()
	redis.SetEx(email, 60, "")
	json.Ok = base.SUCCESS
	return json
}

func SendMobileCode(mobile string) base.Json {
	json := base.GetJson()
	if !redis.Exists(mobile) {
		json.SetError("SMS_TIME_NOT_EX")
		return json
	}
	code := key.Generate(6)
	go func() {
		if sms.SendByYzx(mobile, code) {
			mm, _ := time.ParseDuration(sms.OVERDUE + "m")
			verification := domain.Verification{Code: code, Contact: mobile, Expiry: time.Now().Add(mm)}
			orm.NewOrm().Insert(&verification)
		}
	}()
	redis.SetEx(mobile, 60, "")
	json.Ok = base.SUCCESS
	return json
}

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

func createSession(accountId int64, request http.Request, writer http.ResponseWriter) base.Json {
	json := AddSession(accountId, request, writer)
	if json.Ok == base.SUCCESS {
		user := dao.GetUser(accountId)
		var returnData = make(map[string]interface{})
		returnData["user"] = user
		json.Content = user
	}
	return json
}

func checkPassword(json *base.Json, account domain.Account, password, contact string, request http.Request, writer http.ResponseWriter) base.Json {
	if account.State == 1 {
		json.SetError("ABNORMAL_ACCOUNT")
		return *json
	}
	contactSha256 := crypto.SHA256Hex(contact)
	k := []byte(contactSha256[0:32])
	hex.Decode(k, k)
	pwd, err := crypto.AesDecrypt([]byte(password), k[:16])
	if err != nil {
		json.SetError("PASSWORD_NOT_AES")
		return *json
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), pwd)
	if err != nil {
		json.SetError("WRONG_PASSWORD")
		return *json
	}
	return createSession(account.AccountId, request, writer)
}

func LoginWithPassword(contact, password, areaCode string, request http.Request, writer http.ResponseWriter) base.Json {
	json := base.GetJson()
	var accountId int64 = 0
	var isAuth int = 0
	var mobile domain.Mobile
	if str.IsEmpty(areaCode) {
		mobile = dao.GetMobileByNumber(contact)
	} else {
		mobile = dao.GetMobileByNumber(areaCode + contact)
	}
	if &mobile != nil {
		accountId = mobile.AccountId
		isAuth = mobile.Auth
	} else {
		email := dao.GetEmailByAddress(contact)
		if &email != nil {
			accountId = email.AccountId
			isAuth = email.Auth
		} else {
			user := dao.GetUserByNickName(contact)
			if &user != nil {
				accountId = user.AccountId
				isAuth = 1
			}
		}
	}
	if accountId == 0 || isAuth == 0 {
		json.SetError("NO_ACCOUNT")
		return json
	}
	account := dao.GetAccount(accountId)
	return checkPassword(&json, account, password, contact, request, writer)
}

func GetUser(accountId int64) base.Json {
	json := base.GetJson()
	json.Content = dao.GetUser(accountId)
	json.Ok = base.SUCCESS
	return json
}
