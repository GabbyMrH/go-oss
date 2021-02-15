package httpResponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 返回体封装
func Response(ctx *gin.Context, code string, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
