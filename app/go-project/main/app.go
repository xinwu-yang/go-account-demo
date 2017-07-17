package main

import "github.com/astaxie/beego"
import b "com.cxria/base"

type MainController struct {
	beego.Controller
}

func (m *MainController) Get() {
	j := b.Base{Ok: 1, ErrorCode: 2}
	m.Ctx.WriteString(j.String())
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
