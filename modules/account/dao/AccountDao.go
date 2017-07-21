package dao

import (
	"com.cxria/modules/account/domain"
	"com.cxria/base"
	"github.com/astaxie/beego/orm"
	"github.com/sirupsen/logrus"
)

func GetEmailByAddress(address string) domain.Email {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("email").Where("address = ?")
	o := orm.NewOrm()
	var email domain.Email
	err := o.Raw(qb.String(), address).QueryRow(&email)
	if err != nil {
		logrus.Error(err)
		return nil
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
		logrus.Error(err)
		return nil
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
		logrus.Error(err)
		return nil
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
		logrus.Error(err)
		return nil
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
		logrus.Error(err)
		return nil, &o
	}
	return session, &o
}

func GetVerification(code string, number string) domain.Verification {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("verification").Where("code = ?").And("number = ?")
	o := orm.NewOrm()
	var verification domain.Verification
	err := o.Raw(qb.String(), code, number).QueryRow(&verification)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return verification
}

func GetUserByNickName(nickName string) domain.User {
	qb := base.GetQueryBuilder()
	qb.Select("*").From("user").Where("nickName = ?")
	o := orm.NewOrm()
	var user domain.User
	err := o.Raw(qb.String(), nickName).QueryRow(&user)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return user
}

func GetAccount(accountId int64) domain.Account {
	o := orm.NewOrm()
	account := domain.Account{AccountId: accountId}
	err := o.Read(&account)
	if &err == orm.ErrNoRows {
		return nil
	}
	return account
}
