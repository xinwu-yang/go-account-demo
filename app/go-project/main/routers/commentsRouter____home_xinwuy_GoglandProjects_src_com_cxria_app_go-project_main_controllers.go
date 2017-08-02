package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["com.cxria/app/go-project/main/controllers:AccountController"] = append(beego.GlobalControllerRouter["com.cxria/app/go-project/main/controllers:AccountController"],
		beego.ControllerComments{
			Method: "SendAuthEmail",
			Router: `/sendAuthEmail`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
