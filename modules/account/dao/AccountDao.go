package dao

import "com.cxria/modules/account/domain"

type IAccountDao interface {
	GetAccount(id int64) domain.Account

	GetEmailByAddress(address string) domain.Email

	GetMobileByNumber(number string) domain.Mobile

	GetMobileByAccountId(accountId int64) domain.Mobile

	GetSessionByToken(token string) domain.Session

	GetVerification(code string, mobile string)

	GetUser(accountId int64) domain.User

	GetUserByNickName(nickName string) domain.User
}

type AccountDao struct {

}
