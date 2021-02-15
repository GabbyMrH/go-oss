package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-oss/app/enum/response"
	"go-oss/pkg/config"
	"go-oss/pkg/http/httpResponse"
)

// AppKey验证中间件
func CheckKey() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appKey := ctx.Request.Header.Get("AppKey")
		if appKey == "" || appKey != config.GetString("app.key") {
			httpResponse.Response(ctx, response.REQUEST_FAILS, "操作失败", "Header处AppKey有误")
			// 终止向下执行
			ctx.Abort()
		}
		ctx.Next()
	}
}
