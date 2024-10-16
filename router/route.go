package router

import (
	"dst-admin-go/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRoute() *gin.Engine {

	app := gin.Default()

	// 创建基于cookie的存储引擎，secret11111 参数是用于加密的密钥
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		// 设置过期时间为 30 分钟
		MaxAge:   int(30 * time.Minute.Seconds()),
		Path:     "/",
		HttpOnly: true, // 仅允许在 HTTP 请求中使用，增加安全性
	})
	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	app.Use(sessions.Sessions("mysession", store))

	app.Use(middleware.Recover)
	app.Use(middleware.Authentication())
	app.Use(middleware.Proxy)

	app.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello! Dont starve together 1.1.9.2 20230816")
	})

	// 映射静态文件目录
	app.Static("/dstfile", "./dstfile")

	router := app.Group("")
	initClusterRouter(router)
	initLoginRouter(router)
	initThirdPartyRouter(router)
	initUserRouter(router)

	initStaticFile(app)

	return app
}
