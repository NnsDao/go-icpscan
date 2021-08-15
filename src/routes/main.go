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

func Run() {
	// 增加COSR头-------------------------------
	const (
		accessControlAllowHeaders = "Origin, X-Requested-With, Content-Type, Accept, X-HTTP-Method-Override, Cookie, X-Token, X-Device, X-Plt, X-Ver, X-Page-Id, X-Chl, X-QToken, X-Feature, X-Tk, X-Tag, X-Epid, X-Encrypt, L-Token, X-App, X-Os"
	)
	mwCORS := cors.New(cors.Config{
		//准许跨域请求网站,多个使用,分开,限制使用*
		AllowOrigins: []string{"https://icpscan.co", "https://api.baqiye.com"},
		//准许使用的请求方式
		AllowMethods: []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		//准许使用的请求表头
		AllowHeaders: []string{accessControlAllowHeaders},
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
