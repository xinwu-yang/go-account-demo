package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"com.cxria/utils/str"
	"com.cxria/base"
	"com.cxria/modules/account/domain"
	"com.cxria/modules/account/service"
	"time"
)

//API过滤器重写Content-Type
var apiFilter = func(ctx *context.Context) {
	ctx.Output.Header("Content-Type", "application/json;charset=UTF-8")
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
}

//用户相关过滤器
var userFilter = func(ctx *context.Context) {
	json := base.GetJson()
	token := ctx.GetCookie("t")
	if str.IsEmpty(token) {
		json.SetError("NO_TOKEN")
		ctx.WriteString(json.String())
	}
	session := service.VerifySession(token).Content.(domain.Session)
	dd, _ := time.ParseDuration("360h")
	if session.SessionId == 0 || session.LoginTime.Add(dd).Before(time.Now()) || session.LogoutTime.Year() == 1 {
		json.SetError("NO_TOKEN")
		ctx.WriteString(json.String())
	}
	account := service.GetAccount(session.AccountId).Content.(domain.Account)
	if account.AccountId == 0 || account.State == 1 {
		json.SetError("ABNORMAL_ACCOUNT")
		ctx.WriteString(json.String())
	}
	user := service.GetUser(session.AccountId).Content.(domain.User)
	ctx.Input.SetData("user", user)
	ctx.Input.SetData("token", token)
}

func Config() {
	beego.InsertFilter("/v1/*", beego.BeforeExec, apiFilter)
	beego.InsertFilter("/v1/account/user/*", beego.BeforeExec, userFilter)
}
