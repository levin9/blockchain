package main

import (
	_ "sdrms/routers"
	_ "sdrms/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.AutoRender = false
	beego.Run()
}
