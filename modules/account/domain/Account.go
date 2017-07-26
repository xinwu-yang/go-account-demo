package domain

import (
	"time"
	_ "github.com/astaxie/beego/orm"
)

type Account struct {
	AccountId int64 `orm:"auto"`
	Password  string
	State     int
}

type Email struct {
	AccountId int64 `orm:"pk"`
	Address   string
	Auth      int
}

type Mobile struct {
	AccountId int64 `orm:"pk"`
	Number    string
	Auth      int
}

type Session struct {
	SessionId  int64 `orm:"auto"`
	AccountId  int64
	UserAgent  string
	Token      string
	LoginTime  time.Time
	LogoutTime time.Time
}

type User struct {
	AccountId  int64 `orm:"pk"`
	NickName   string
	Avatar     string
	RealName   string
	Address    string
	IsVip      int
	VipLevel   int
	Type       int
	CreateTime time.Time
}

type Verification struct {
	VerificationId int64  `orm:"auto"`
	Code           string `orm:"size(6)"`
	Contact        string
	Expiry         time.Time
}
