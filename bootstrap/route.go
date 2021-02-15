package bootstrap

import (
	"github.com/gin-gonic/gin"
	"go-oss/app/middlewares"
	"go-oss/routes"
)

// 引导安装路由
func SetupRoute() *gin.Engine {
	router := gin.Default()

	// 资源文件路由
	router.Static("/resources", "./resources")
	// Web 路由组注册
	routes.RegisterWebRouters(router)
	// Api 路由组注册
	routes.RegisterApiRoutes(router)

	// 全局中间件
	// 跨域处理中间件
	router.Use(middlewares.Cors())
	return router
}
