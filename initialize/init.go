package initialize

import (
	"beego-admin/initialize/conf"
	"beego-admin/initialize/logs"
	"beego-admin/initialize/mysql"
	"beego-admin/initialize/session"
)

func init() {
	logs.Init()
	conf.Init()
	session.Init()
	mysql.Init()
}
