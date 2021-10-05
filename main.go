package main

import (
	Router "icpscan/src/routes"
	"icpscan/src/task"
	"log"

	"github.com/robfig/cron/v3"
)

type User struct {
	ID   uint
	Name string
}

// @host 139.224.134.111:8080

func main() {
	// 解析配置文件，就不做service了简单来
	t, err := task.NewTaskService(cron.New())
	if err != nil {
		log.Printf("task is fail %v", err)
	}
	t.Run()
	Router.Run()
}
