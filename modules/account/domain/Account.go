package domain

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Account struct {
	AccountId int64 `orm:"auto"`
	Password  string
	State     int `orm:"default(0)"`
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
	LoginTime  time.Time `orm:"auto_now_add;type(datetime)"`
	LogoutTime time.Time `orm:"null"`
}

type User struct {
	AccountId  int64 `orm:"pk"`
	NickName   string
	Avatar     string `orm:"default(default_avatar.png)"`
	RealName   string
	Address    string
	Type       int `orm:"default(1)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

type Verification struct {
	VerificationId int64  `orm:"auto"`
	Code           string `orm:"size(6)"`
	Contact        string
	Expiry         time.Time
}

func init() {
	orm.RegisterModel(new(Account),
		new(Email),
		new(Mobile),
		new(Verification),
		new(Session),
		new(User))
}
