package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-oss/pkg/logger"
)

// 记录异常(未启用)
func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// r.(string) 将 interface{} 转型为具体类型
			logger.LogError(errors.New(r.(string)))
			// 打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			//debug.PrintStack()
			// 是否终止后续接口调用，不加则recover到异常后，还会继续执行接口里后续代码
			//ctx.Abort()
		}
	}()
	// 加载完 defer recover，继续后续接口调用
	ctx.Next()
}
