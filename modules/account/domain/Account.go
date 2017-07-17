package domain

import (
	"time"
)

type Account struct {
	AccountId int64
	Password  string
	State     int
}

type Email struct {
	AccountId int64
	Address   string
	Auth      int
}

type Mobile struct {
	AccountId int64
	Number    string
	Auth      int
}

type Session struct {
	SessionId int64
	AccountId int64
	UserAgent string
	Token string
	LoginTime time.Time
}

type User struct {
	AccountId int64
	NickName string
	Avatar string
	RealName string
	Address string
	IsVip int
	VipLevel int
	Type int
	CreateTime time.Time
}

type Verification struct {
	VerificationId int64
	Code string `orm:"size(6)"`
	Mobile string
	Expiry time.Time
}
