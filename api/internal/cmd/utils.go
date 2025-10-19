package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"toolbox/internal/bootstrap/boot"
	"toolbox/internal/bootstrap/database"
	"toolbox/internal/infrastructure/query"
	"toolbox/pkg/config"
	"toolbox/pkg/constants"
	"toolbox/pkg/utils/file"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

func getString(cmd *cobra.Command, name string) string {
	value, _ := cmd.Flags().GetString(name)
	return value
}

func preRunInitialization(cmd *cobra.Command, args []string) {
	dataPath := getString(cmd, "data")
	if dataPath == "." {
		return
	}
	if !strings.HasSuffix(dataPath, constants.ProjectName) {
		dataPath = filepath.Join(dataPath, constants.ProjectName)
	}
	constants.DataPath = filepath.ToSlash(dataPath)
	if err := file.CreateDir(constants.DataPath); err != nil {
		log.Printf("创建数据目录失败 %s, err: %v", constants.DataPath, err)
	}

	constants.UploadPath = filepath.ToSlash(filepath.Join(dataPath, "uploads"))
	if err := file.CreateDir(constants.UploadPath); err != nil {
		log.Printf("创建上传目录失败 %s, err: %v", constants.UploadPath, err)
	}
}

func flagsData(cmd *cobra.Command, usage string) {
	cmd.Flags().StringP("data", "d", ".", usage)
}

func initializationData() {
	if configPath := filepath.Join(constants.DataPath, "config.ini"); !file.IsFileExist(configPath) {
		config.SaveToIni()
		// 初始化系统日志
		boot.InitializationLogger()
		// 连接数据库
		db := boot.InitializationDB()
		if err := database.AutoMigrate(db, query.Q, nil); err != nil {
			fmt.Println("数据库迁移失败", err)
		}
		fmt.Println("数据初始化完成")
		fmt.Printf("配置文件已生成，路径：%s\n", filepath.ToSlash(configPath))
	}
}

func Execute() error {
	svcConfig = &service.Config{
		Name:        "toolbox",
		DisplayName: constants.WebTitle,
		Description: "本地常用工具集合，提供本地服务管理、DNS解析、Markdown文档编辑等功能，默认管理页面：http://localhost:6500",
	}

	prg = &program{}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Printf("service.New error %s, err: %v", svcConfig.Name, err)
		os.Exit(1)
	}
	serv = s

	return toolboxCmd.Execute()
}
