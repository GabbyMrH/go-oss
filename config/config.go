package config

// Initialize 配置信息初始化
func Initialize() {
	// 触发加载本目录下其他文件中的 init 方法

	// 若生产环境则需取消下方注释
	// gin.SetMode(gin.ReleaseMode)
}
