package main

import (
	"io"
	"log"
	"os"
	"time"
	"toolbox/internal/cmd"
	"toolbox/pkg/constants"
	"toolbox/pkg/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	time.Local = time.UTC
	// 禁用 Gin 框架的日志输出
	gin.DefaultWriter = io.Discard
	// 初始化雪花算法节点
	utils.InitSnowflake(1)
	constants.RunTime = time.Now()
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Printf("执行命令失败, err: %v\n", err)
		os.Exit(1)
	}
}
