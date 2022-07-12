package main

import (
	_ "beego-admin/initialize"
	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	beego.Run() // 启动beego
}
