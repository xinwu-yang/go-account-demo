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
	if session.SessionId != 0 {
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

func addSession(accountId int64, userAgent string, writer http.ResponseWriter) base.Json {
	json := base.GetJson()
	token := key.Generate(32)
	var session = domain.Session{AccountId: accountId, Token: token, UserAgent: userAgent}
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
	if redis.Exists(mobile) {
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

func verifyCode(code string, contact string) base.Json {
	json := base.GetJson()
	verification := dao.GetVerification(code, contact)
	if verification.VerificationId == 0 || verification.Expiry.Before(time.Now()) {
		json.SetError("EXPIRED_CODE")
		return json
	}
	json.Ok = base.SUCCESS
	return json
}

func createSession(accountId int64, userAgent string, writer http.ResponseWriter) base.Json {
	json := addSession(accountId, userAgent, writer)
	if json.Ok == base.SUCCESS {
		user := dao.GetUser(accountId)
		var returnData = make(map[string]interface{})
		returnData["user"] = user
		json.Content = user
	}
	return json
}

func checkPassword(json *base.Json, account domain.Account, password, contact string, request *http.Request, writer http.ResponseWriter) base.Json {
	if account.State == 1 {
		json.SetError("ABNORMAL_ACCOUNT")
		return *json
	}
	contactSha256 := crypto.SHA256Hex(contact)
	k := []byte(contactSha256[0:32])
	i := []byte(contactSha256[32:64])
	hex.Decode(k, k)
	hex.Decode(i, i)
	pwd, err := crypto.AesDecrypt([]byte(password), k[:16], i[:16])
	if err != nil {
		json.SetError("PASSWORD_NOT_AES")
		return *json
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), pwd)
	if err != nil {
		json.SetError("WRONG_PASSWORD")
		return *json
	}
	return createSession(account.AccountId, request.UserAgent(), writer)
}

func createAccount(account, password string) domain.Account {
	var a domain.Account
	var pwd []byte
	if str.IsEmpty(password) {
		pwd, _ = bcrypt.GenerateFromPassword([]byte(account), bcrypt.DefaultCost)
		a = domain.Account{Password: string(pwd), State: 0}
	} else {
		shaAccount := crypto.SHA256Hex(account)
		k := []byte(shaAccount[0:32])
		i := []byte(shaAccount[32:64])
		hex.Decode(k, k)
		hex.Decode(i, i)
		aesPwd, _ := hex.DecodeString(password)
		plainText, err := crypto.AesDecrypt(aesPwd, k[:16], i[:16])
		if err != nil {
			return a
		}
		pwd, _ = bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	}
	a = domain.Account{Password: string(pwd), State: 0}
	user := domain.User{NickName: account, Avatar: "default_avatar.png", IsVip: 0, VipLevel: 1, Type: 1}
	o := orm.NewOrm()
	o.Begin()
	accountId, createAccountErr := o.Insert(&a)
	user.AccountId = accountId
	_, createUserErr := o.Insert(&user)
	if createAccountErr != nil || createUserErr != nil {
		o.Rollback()
	} else {
		o.Commit()
	}
	return a
}

func LoginWithPassword(contact, password, areaCode string, request *http.Request, writer http.ResponseWriter) base.Json {
	json := base.GetJson()
	var accountId int64 = 0
	var isAuth int = 0
	var mobile domain.Mobile
	if str.IsEmpty(areaCode) {
		mobile = dao.GetMobileByNumber(contact)
	} else {
		mobile = dao.GetMobileByNumber(areaCode + contact)
	}
	if mobile.AccountId != 0 {
		accountId = mobile.AccountId
		isAuth = mobile.Auth
	} else {
		email := dao.GetEmailByAddress(contact)
		if email.AccountId != 0 {
			accountId = email.AccountId
			isAuth = email.Auth
		} else {
			user := dao.GetUserByNickName(contact)
			if user.AccountId != 0 {
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

func LoginWithCode(mobile, code string, request *http.Request, writer http.ResponseWriter) base.Json {
	json := verifyCode(code, mobile)
	if json.Ok == base.SUCCESS {
		var account domain.Account
		mobileAccount := dao.GetMobileByNumber(mobile)
		if mobileAccount.AccountId == 0 {
			account = createAccount(mobile, "")
			m := domain.Mobile{AccountId: account.AccountId, Number: mobile, Auth: 1}
			orm.NewOrm().Insert(&m)
		} else if mobileAccount.Auth == 0 {
			json.SetError("CONTACT_NOT_AUTH")
			return json
		} else {
			account = dao.GetAccount(mobileAccount.AccountId)
			if account.AccountId == 0 || account.State == 1 {
				json.SetError("ABNORMAL_ACCOUNT")
				return json
			}
		}
		json = createSession(account.AccountId, request.UserAgent(), writer)
	}
	return json
}
