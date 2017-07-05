package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblog/lib/runner"
	_ "myblog/routers"
	"runtime"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/myblog.log","maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.BeeLogger.DelLogger("console")
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)

	runner.Init()

	//设置并发数
	cpus := runtime.NumCPU()
	n := runtime.GOMAXPROCS(cpus)
	beego.Info(fmt.Sprintf("cpu bum :%d,default: %d", cpus, n))

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.Listen.EnableAdmin = true
		beego.BConfig.Listen.AdminAddr = ""
		beego.BConfig.Listen.AdminPort = 8089
	}

	beego.Run()
}
