package main

import (
	Router "github.com/MatheusMeloAntiquera/api-go/src/routes"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	// 解析配置文件，就不做service了简单来
	Router.Run()
}

