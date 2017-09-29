package configure

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	//account  string = "gaiamount"
	//password string = "chengxun"
	//host     string = "192.168.1.21"

	account  string = "go"
	password string = "Chengxun@1806"
	host     string = "rm-wz9eynjsfx9iqm8tpo.mysql.rds.aliyuncs.com"

	port string = "3306"
	db   string = "go"
)

func ConfigDataSource() {
	orm.RegisterDataBase("default", "mysql", account+":"+password+"@("+host+":"+port+")/"+db+"?charset=utf8mb4,utf8&timeout=90s&collation=utf8mb4_unicode_ci&loc=Asia%2FShanghai")
	orm.RunSyncdb("default", false, true)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.Local
	orm.Debug = true
}
