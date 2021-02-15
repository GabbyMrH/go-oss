package routes

import (
	"github.com/gin-gonic/gin"
	"go-oss/app/enum/response"
	"go-oss/pkg/http/httpResponse"
)

// Web 路由组
func RegisterWebRouters(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		httpResponse.Response(ctx, response.REQUEST_SUCCESS, "操作成功", "欢迎使用本系统")
	})
}
