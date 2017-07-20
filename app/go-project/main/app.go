package main

import "github.com/astaxie/beego"
import (
	b "com.cxria/base"
	"com.cxria/app/go-project/main/configure"
	"runtime"
	"fmt"
	"encoding/json"
)

type TestController struct {
	beego.Controller
}

type Param struct {
	Name string
}

func (m *TestController) Post() {
	var body Param
	json.Unmarshal(m.Ctx.Input.RequestBody, &body)
	j := b.Base{Ok: 1, Content: body}
	m.Ctx.WriteString(j.String())
}

func init() {
	configure.ConfigDataSource()
}

func main() {
	//项目配置在beego.BConfig中
	beego.BConfig.CopyRequestBody = true
	fmt.Println("CPU Num:", runtime.NumCPU())
	beego.Router("/", &TestController{})
	beego.Run()
}
