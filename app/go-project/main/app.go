package main

import "github.com/astaxie/beego"
import (
	"com.cxria/app/go-project/main/configure"
	"runtime"
	"fmt"
	"com.cxria/app/go-project/main/routers"
	"com.cxria/app/go-project/main/filters"
)

func init() {
	configure.ConfigDataSource()
	routers.Config()
	filters.Config()
}

func main() {
	//项目配置在beego.BConfig中
	fmt.Println("CPU Num:", runtime.NumCPU())
	beego.BConfig.AppName = "GoProject"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.EnableDocs = true
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "/home/xinwuy/GoglandProjects/src/com.cxria/app/go-project/main/swagger"
	beego.Run()
}
