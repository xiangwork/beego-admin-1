package logs

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
)

func init() {
	filename := time.Now().Format("20060102") + ".log"
	logs.SetLogger(logs.AdapterFile, `{"filename":"logs/`+filename+`","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":3,"color":true}`)
	logs.EnableFuncCallDepth(true)
}
