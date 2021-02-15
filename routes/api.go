package routes

import (
	"github.com/gin-gonic/gin"
	"go-oss/app/http/controllers"
	"go-oss/app/middlewares"
)

// Api 路由组
func RegisterApiRoutes(r *gin.Engine) {
	// 文件相关路由组
	fileRouter := r.Group("/file").Use(middlewares.CheckKey())
	{
		fc := new(controllers.FileController)
		fileRouter.POST("/upload", fc.FileUpload)
		fileRouter.POST("/delete", fc.FileDrop)
	}

	// 助手相关路由组
	handlerRouter := r.Group("/handler")
	{
		hc := new(controllers.HandlerController)
		handlerRouter.GET("/app-key", hc.GenerateAppKey)
	}
}
