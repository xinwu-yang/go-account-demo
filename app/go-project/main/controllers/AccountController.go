package controllers

import (
	"github.com/astaxie/beego"
	"com.cxria/modules/account/service"
	"com.cxria/base"
	"encoding/json"
	"github.com/astaxie/beego/validation"
)

type AccountController struct {
	beego.Controller
}

type SendAuthEmailParam struct {
	Email     string `valid:"Email; MaxSize(100)"`
	EmailType int `valid:"Min(1); Max(2)"`
	AccountId int `valid:"Min(1); Max(999999999)"`
}

type SendSmsCodeParam struct {
	Mobile string `valid:"Mobile; MaxSize(100)"`
}

func (a *AccountController) URLMapping() {
	a.Mapping("sendAuthEmail", a.SendAuthEmail)
	a.Mapping("sendSmsCode", a.SendSmsCode)
}

// @Title 发送邮件
// @Description 发送认证邮件和忘记密码
// @Param   body body     SendAuthEmailParam true "账号Id<br>邮件地址<br>邮件类型：1=邮箱认证，2=忘记密码"
// @Success 200  {string} {"b" : 1}
// @router /sendAuthEmail [post]
func (a *AccountController) SendAuthEmail() {
	var param SendAuthEmailParam
	json.Unmarshal(a.Ctx.Input.RequestBody, &param)
	valid := validation.Validation{}
	b, _ := valid.Valid(&param)
	if !b {
		j := base.GetJson()
		j.ErrorArray = valid.Errors
		a.Ctx.WriteString(j.String())
		return
	}
	j := service.SendAuthEmail(param.Email, param.EmailType, int64(param.AccountId))
	a.Ctx.WriteString(j.String())
}

// @Title 发送短信验证码
// @Description 发送短信验证码
// @Param   body body     SendSmsCodeParam true "手机号码"
// @Success 200  {string} {"b" : 1}
// @router /sendSmsCode [post]
func (a *AccountController) SendSmsCode() {
	var param SendSmsCodeParam
	json.Unmarshal(a.Ctx.Input.RequestBody, &param)
	valid := validation.Validation{}
	b, _ := valid.Valid(&param)
	if !b {
		j := base.GetJson()
		j.ErrorArray = valid.Errors
		a.Ctx.WriteString(j.String())
		return
	}
	j := service.SendMobileCode(param.Mobile)
	a.Ctx.WriteString(j.String())
}
