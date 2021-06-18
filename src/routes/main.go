package routes

import (
	"github.com/MatheusMeloAntiquera/api-go/src/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	Router = gin.Default()
)

func Run() {
	addUserRoutes()
	addBlockRoutes()
	Router.Use(middlewares.Cors())
	Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
