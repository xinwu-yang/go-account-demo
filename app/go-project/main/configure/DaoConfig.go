package configure

import (
	"github.com/astaxie/beego/orm"
	"com.cxria/modules/account/domain"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	account  string = "grassroot"
	password string = "chengxun"
	host     string = "192.168.1.21"
	port     string = "3306"
	db       string = "go"
)

func ConfigDataSource() {
	orm.RegisterDataBase("default", "mysql", account+":"+password+"@("+host+":"+port+")/"+db+"?charset=utf8mb4,utf8&timeout=90s&collation=utf8mb4_unicode_ci&loc=Asia%2FShanghai")
	registerModel()
	orm.RunSyncdb("default", false, true)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.Local
	orm.Debug = true
}

func registerModel() {
	orm.RegisterModel(new(domain.Account), new(domain.Email), new(domain.Mobile), new(domain.User), new(domain.Verification), new(domain.Session))
}
