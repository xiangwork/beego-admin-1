package main

import (
	_ "beego-admin/initialize"
	_ "beego-admin/routers"
	_ "beego-admin/utils/template"

	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	beego.Run() // 启动beego
}
