package config

import "go-oss/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称，暂时没有使用到
		"name": config.Env("APP_NAME", "go-oss"),

		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "production"),

		// 是否进入调试模式
		"debug": config.Env("APP_DEBUG", false),

		// 应用服务端口
		"url": config.Env("APP_URL", "http://localhost:8080"),

		// 应用服务端口
		"port": config.Env("APP_PORT", ":8080"),

		// APP Key,用于操作时鉴权(临时)
		"key": config.Env("APP_KEY", ""),
	})
}
