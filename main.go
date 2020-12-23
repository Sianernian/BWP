package main

import (
	_ "BWP/routers"
	"github.com/astaxie/beego"
)

func main() {
	//加载静态资源

	beego.SetStaticPath("/views","./views")
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/img","./static/img")
	beego.SetStaticPath("/css","./static/css")

	beego.Run()
}

