package session

import (
	"beego-admin/models"
	"encoding/gob"
)

// Init 初始化
func Init() {
	//注册session 解析结构
	gob.Register(models.AdminUser{})
}
