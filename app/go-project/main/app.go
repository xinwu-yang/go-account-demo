package main

import "github.com/astaxie/beego"
import (
	b "com.cxria/base"
	"com.cxria/app/go-project/main/configure"
	"runtime"
	"fmt"
	"encoding/json"
	"com.cxria/app/go-project/main/routers"
	"com.cxria/app/go-project/main/filters"
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
	j := b.Json{Ok: 1, Content: body}
	m.Ctx.WriteString(j.String())
}

func init() {
	configure.ConfigDataSource()
	routers.Config()
	filters.Config()
}

func main() {
	//项目配置在beego.BConfig中
	fmt.Println("CPU Num:", runtime.NumCPU())
	//beego.Router("/", &TestController{})
	beego.BConfig.AppName = "GoPro"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "/home/xinwuy/GoglandProjects/src/com.cxria/app/go-project/main/swagger"
	beego.Run()
}
