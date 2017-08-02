// @APIVersion 1.0.0
// @Title Go Pro API
// @Description 这是一个基于Golang开发的Web服务器
// @Contact summng@qq.com
// @TermsOfServiceUrl NO terms of service
// @License The Apache License, Version 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"com.cxria/app/go-project/main/controllers"
	"fmt"
)

func Config() {
	fmt.Println("Test")
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/account", beego.NSInclude(
			&controllers.AccountController{},
		)))
	beego.AddNamespace(ns)
}
