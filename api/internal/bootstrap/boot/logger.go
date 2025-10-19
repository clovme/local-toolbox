package boot

import (
	"fmt"
	"os"
	"toolbox/pkg/config"
	"toolbox/pkg/constants"
	"toolbox/pkg/logger"
	"toolbox/pkg/utils/file"
)

// InitializationLogger 初始化日志
func InitializationLogger() {
	path, err := file.GetFileAbsPath(constants.DataPath, "logs")
	if err != nil {
		fmt.Println("获取日志目录失败:", err)
		os.Exit(-1)
	}
	// 初始化一次
	cfg := config.GetConfig()
	logger.InitLogger(logger.LoggerConfig{
		Dir:        path,
		MaxSize:    cfg.Logger.MaxSize,
		MaxBackups: cfg.Logger.MaxBackups,
		MaxAge:     cfg.Logger.MaxAge,
		Compress:   true,
		Level:      cfg.Logger.Level,
		FormatJSON: false, // true=结构化；false=文本
	})
}
