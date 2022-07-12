package initialize

import (
	"beego-admin/initialize/conf"
	"beego-admin/initialize/logs"
	"beego-admin/initialize/mysql"
	"beego-admin/initialize/session"
	"beego-admin/routers"
	_ "beego-admin/utils/template"
)

func init() {
	logs.Init()
	conf.Init()
	session.Init()
	mysql.Init()
	routers.Init()
}
