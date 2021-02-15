package main

import (
	"go-oss/bootstrap"
	"go-oss/config"
	c "go-oss/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	bootstrap.SetupRoute().Run(c.GetString("app.port"))
}
