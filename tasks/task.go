package tasks

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/task"
)

// Init 初始化 定时任务
func Init() {
	tk1 := task.NewTask("tk1", "* * * * * *",
		func(ctx context.Context) error {
			logs.Info("tk1 is running")
			return nil
		})

	task.AddTask("tk1", tk1)
	task.StartTask()
	defer task.StopTask()
}
