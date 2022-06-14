package main

import (
	_ "beego-admin/initialize/conf"
	_ "beego-admin/initialize/mysql"
	_ "beego-admin/initialize/session"
	_ "beego-admin/routers"
	_ "beego-admin/utils/template"

	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	beego.Run() // 启动beego
}
