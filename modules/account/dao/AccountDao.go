package dao

import (
	"com.cxria/modules/account/domain"
	"com.cxria/base"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/logs"
)

func GetEmailByAddress(address string) domain.Email {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("email").Where("address = ?")
	o := orm.NewOrm()
	var email domain.Email
	err := o.Raw(qb.String(), address).QueryRow(&email)
	if err != nil {
		logs.Error(err)
		return domain.Email{}
	}
	return email
}

func GetEmailByAccountId(accountId int64) domain.Email {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("email").Where("accountId = ?")
	o := orm.NewOrm()
	var email domain.Email
	err := o.Raw(qb.String(), accountId).QueryRow(&email)
	if err != nil {
		logs.Error(err)
		return domain.Email{}
	}
	return email
}

func GetMobileByNumber(number string) domain.Mobile {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("mobile").Where("number = ?")
	o := orm.NewOrm()
	var mobile domain.Mobile
	err := o.Raw(qb.String(), number).QueryRow(&mobile)
	if err != nil {
		logs.Error(err)
		return domain.Mobile{}
	}
	return mobile
}

func GetMobileByAccountId(accountId int64) domain.Mobile {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("mobile").Where("accountId = ?")
	o := orm.NewOrm()
	var mobile domain.Mobile
	err := o.Raw(qb.String(), accountId).QueryRow(&mobile)
	if err != nil {
		logs.Error(err)
		return domain.Mobile{}
	}
	return mobile
}

func GetSessionByToken(token string) (domain.Session, orm.Ormer) {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("session").Where("token = ?")
	o := orm.NewOrm()
	var session domain.Session
	err := o.Raw(qb.String(), token).QueryRow(&session)
	if err != nil {
		logs.Error(err)
		return domain.Session{}, o
	}
	return session, o
}

func GetVerification(code string, number string) domain.Verification {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("verification").Where("code = ?").And("contact = ?")
	o := orm.NewOrm()
	var verification domain.Verification
	err := o.Raw(qb.String(), code, number).QueryRow(&verification)
	if err != nil {
		logs.Error(err)
		return domain.Verification{}
	}
	return verification
}

func GetUserByNickName(nickName string) domain.User {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("user").Where("nick_name = ?")
	o := orm.NewOrm()
	var user domain.User
	err := o.Raw(qb.String(), nickName).QueryRow(&user)
	if err != nil {
		logs.Error(err)
		return domain.User{}
	}
	return user
}

func GetAccount(accountId int64) domain.Account {
	o := orm.NewOrm()
	account := domain.Account{AccountId: accountId}
	err := o.Read(&account)
	if err == orm.ErrNoRows {
		return domain.Account{}
	}
	return account
}

func GetUser(accountId int64) domain.User {
	o := orm.NewOrm()
	user := domain.User{AccountId: accountId}
	err := o.Read(&user)
	if err == orm.ErrNoRows {
		return domain.User{}
	}
	return user
}
