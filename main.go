package main

import (
	Router "github.com/MatheusMeloAntiquera/api-go/src/routes"
)

type User struct {
	ID   uint
	Name string
}

func main() {
	Router.Run()
}
