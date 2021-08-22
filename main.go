package main

import (
	"log"

	Router "github.com/MatheusMeloAntiquera/api-go/src/routes"
	"github.com/MatheusMeloAntiquera/api-go/src/task"
	"github.com/robfig/cron/v3"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// 解析配置文件，就不做service了简单来
	t, err := task.NewTaskService(cron.New())
	if err != nil {
		log.Printf("task is fail %v", err)
	}
	t.Run()
	Router.Run()
}
