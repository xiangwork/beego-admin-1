package initialize

import (
	"beego-admin/initialize/conf"
	"beego-admin/initialize/logs"
	"beego-admin/initialize/mysql"
	"beego-admin/initialize/session"
	"beego-admin/routers"
	_ "beego-admin/utils/template"
	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func init() {
	logs.Init()
	conf.Init()
	session.Init()
	mysql.Init()
	routers.Init()
}
