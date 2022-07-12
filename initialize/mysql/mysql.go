package mysql

import (
	"beego-admin/global"
	"fmt"
	"github.com/beego/beego/v2/core/logs"

	"github.com/beego/beego/v2/client/orm"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Init  注册mysql
func Init() {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		logs.Error("mysql register driver error:", err)
	}

	//dataSource := "root:root@tcp(127.0.0.1:3306)/test"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		global.BA_CONFIG.Mysql.Username,
		global.BA_CONFIG.Mysql.Password,
		global.BA_CONFIG.Mysql.Host,
		global.BA_CONFIG.Mysql.Port,
		global.BA_CONFIG.Mysql.Database,
	)

	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		logs.Error("mysql register database error:", err)
	}
}
