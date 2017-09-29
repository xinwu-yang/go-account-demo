package controllers

import (
	"github.com/astaxie/beego"
	"com.cxria/modules/account/service"
	"com.cxria/base"
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"com.cxria/utils/mail"
)

type AccountController struct {
	beego.Controller
}

type SendEmailParam struct {
	To      string `valid:"Email; MaxSize(100)"`
	Html    string `valid:"MinSize(1);"`
	Subject string `valid:"MinSize(1);"`
}

type SendAuthEmailParam struct {
	Email     string `valid:"Email; MaxSize(100)"`
	EmailType int    `valid:"Min(1); Max(2)"`
	AccountId int    `valid:"Min(1); Max(999999999)"`
}

type SendSmsCodeParam struct {
	Mobile string `valid:"Mobile; MaxSize(100)"`
}

type LoginParam struct {
	Account  string `valid:"MaxSize(100)"`
	Password string `valid:"Length(32)"`
	AreaCode string
}

type LoginWithCodeParam struct {
	SendSmsCodeParam
	Code string `valid:"Length(6)"`
}

func (a *AccountController) URLMapping() {
	a.Mapping("sendAuthEmail", a.SendAuthEmail)
	a.Mapping("sendEmail", a.SendEmail)
	a.Mapping("sendSmsCode", a.SendSmsCode)
	a.Mapping("login", a.Login)
	a.Mapping("loginWithCode", a.LoginWithCode)
}

// @Title 发送邮件
// @Description 发送Html邮件
// @Param   body body     SendEmailParam true "收件人<br>收件内容<br>标题"
// @Success 200  {string} {"b" : 1}
// @router /sendEmail [post]
func (a *AccountController) SendEmail() {
	var param SendEmailParam
	json.Unmarshal(a.Ctx.Input.RequestBody, &param)
	valid := validation.Validation{}
	b, _ := valid.Valid(&param)
	if !b {
		j := base.GetJson()
		j.ErrorArray = valid.Errors
		a.Ctx.WriteString(j.String())
		return
	}
	mail.Send(param.To, param.Html, param.Subject)
	j := base.Json{Ok: 1}
	a.Ctx.WriteString(j.String())
}

// @Title 发送验证邮件
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

// @Title 登录
// @Description 手机/邮箱/昵称登录
// @Param   body body     LoginParam true "账号<br>密码<br>区号"
// @router /login [post]
func (a *AccountController) Login() {
	var param LoginParam
	json.Unmarshal(a.Ctx.Input.RequestBody, &param)
	valid := validation.Validation{}
	b, _ := valid.Valid(&param)
	if !b {
		j := base.GetJson()
		j.ErrorArray = valid.Errors
		a.Ctx.WriteString(j.String())
		return
	}
	j := service.LoginWithPassword(param.Account, param.Password, param.AreaCode, a.Ctx.Request, a.Ctx.ResponseWriter.ResponseWriter)
	a.Ctx.WriteString(j.String())
}

// @Title 验证码登录
// @Description 手机验证码登录
// @Param   body body     LoginWithCodeParam true "手机账号<br>验证码"
// @router /loginWithCode [post]
func (a *AccountController) LoginWithCode() {
	var param LoginWithCodeParam
	json.Unmarshal(a.Ctx.Input.RequestBody, &param)
	valid := validation.Validation{}
	b, _ := valid.Valid(&param)
	if !b {
		j := base.GetJson()
		j.ErrorArray = valid.Errors
		a.Ctx.WriteString(j.String())
		return
	}
	j := service.LoginWithCode(param.Mobile, param.Code, a.Ctx.Request, a.Ctx.ResponseWriter.ResponseWriter)
	a.Ctx.WriteString(j.String())
}