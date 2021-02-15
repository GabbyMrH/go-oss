package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-oss/app/enum/response"
	"go-oss/pkg/handlers"
	"go-oss/pkg/http/httpResponse"
)

type HandlerController struct {
	BaseController
}

// 生成app key
func (hc *HandlerController) GenerateAppKey(ctx *gin.Context) {
	keyData := handlers.GenerateRandomString(32)
	fmt.Println(keyData)
	httpResponse.Response(ctx, response.REQUEST_SUCCESS, "成功", keyData)
}
