package logs

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// Init 初始化
func Init() {
	filename := time.Now().Format("20060102") + ".log"
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"storage/logs/`+filename+`","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":0}`)
}
