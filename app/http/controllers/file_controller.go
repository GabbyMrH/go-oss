package controllers

import (
	"github.com/gin-gonic/gin"
	"go-oss/app/enum/response"
	c "go-oss/pkg/config"
	"go-oss/pkg/handlers"
	"go-oss/pkg/http/httpResponse"
	"go-oss/pkg/logger"
	"os"
	"time"
)

type FileController struct {
	BaseController
}

// 单文件上传
func (fc *FileController) FileUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	// 验证文件
	if err != nil {
		logger.LogError(err)
		httpResponse.Response(ctx, response.REQUEST_FAILS, err.Error(), "文件上传失败")
		return
	}
	// 获取文件后缀
	extString := handlers.Ext(file.Filename)
	// 存储位置
	savePath := "resources/" + time.Now().Format("2006-01-02") + "/"
	// 存储路径不存在则创建
	if _, err := os.Stat(savePath); err != nil {
		if !os.IsExist(err) {
			os.MkdirAll(savePath, os.ModePerm)
		}
	}
	// 文件名
	fileName := handlers.GenerateRandomString(32) + extString
	_ = ctx.SaveUploadedFile(file, savePath+fileName)
	// 返回数据
	result := make(map[string]string)
	result["url"] = c.GetString("app.url")
	result["file_name"] = fileName
	result["save_path"] = savePath + fileName
	httpResponse.Response(ctx, response.REQUEST_SUCCESS, "操作成功", result)

}

// 文件删除
func (fc *FileController) FileDrop(ctx *gin.Context) {

	path := ctx.PostForm("path") //源文件路径
	//删除文件
	err := os.Remove(path)

	if err != nil {
		logger.LogError(err)
		httpResponse.Response(ctx, response.REQUEST_FAILS, "操作失败", err.Error())
		return
	} else {
		httpResponse.Response(ctx, response.REQUEST_SUCCESS, "操作成功", path)
	}
}
