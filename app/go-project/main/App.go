package main

import (
	"github.com/astaxie/beego"
	"com.cxria/app/go-project/main/configure"
	"runtime"
	"com.cxria/app/go-project/main/filters"
	"github.com/astaxie/beego/logs"
	_ "com.cxria/app/go-project/main/routers"
)

func init() {
	configure.ConfigDataSource()
	filters.Config()
}

func main() {
	logs.EnableFuncCallDepth(true)
	logs.Async()

	//level 6 = info
	logs.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logs.SetLogger(logs.AdapterFile, `{"filename":"/var/ftp/logs/go.log","level":7,"daily":true}`)

	//项目配置在beego.BConfig中
	logs.Informational("CPU Num:", runtime.NumCPU())
	beego.BConfig.AppName = "GoProject"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "/home/xinwuy/GoglandProjects/src/com.cxria/app/go-project/main/swagger"
	beego.Run()
}
