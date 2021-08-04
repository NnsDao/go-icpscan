package routes

import (
	"time"

	_ "github.com/MatheusMeloAntiquera/api-go/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	Router = gin.Default()
)

// @host 139.224.134.111:8080

func Run() {

	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"*"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{"Origin", "Authorization", "Content-Type"},
		//显示的请求表头
		ExposeHeaders: []string{"Content-Type"},
		//凭证共享,确定共享
		AllowCredentials: true,
		//容许跨域的原点网站,可以直接return true就万事大吉了
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		//超时时间设定
		MaxAge: 24 * time.Hour,
	})
	Router.Use(mwCORS)
	addUserRoutes()
	addBlockRoutes()

	// 开启swag
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
