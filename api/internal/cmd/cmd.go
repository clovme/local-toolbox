package cmd

import (
	"fmt"
	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"toolbox/pkg/config"
	"toolbox/pkg/constants"
)

var (
	serv      service.Service
	prg       *program
	svcConfig *service.Config
)

const toolboxHelpText = `
说明:
   - 安装/卸载/启动/停止 服务需要管理员权限
   - server 模式会在控制台运行并输出日志，适合开发和调试使用
`

var toolboxCmd = &cobra.Command{
	Use:  constants.ProjectName,
	Long: toolboxHelpText,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "安装服务",
	Run: func(cmd *cobra.Command, args []string) {
		svcConfig.Arguments = []string{"run", fmt.Sprintf("--data=%s", constants.DataPath)}

		if err := serv.Install(); err != nil {
			fmt.Println("安装服务失败：", err)
		} else {
			initializationData()
			fmt.Println("服务安装成功")
			fmt.Printf("使用 %s start 启动服务\n", constants.ProjectName)
		}
	},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "卸载服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serv.Uninstall(); err != nil {
			fmt.Println("卸载服务失败：", err)
		} else {
			fmt.Println("服务卸载成功")
		}
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serv.Start(); err != nil {
			fmt.Println("启动服务失败：", err)
		} else {
			cfg := config.GetConfig()
			fmt.Println("服务启动成功")
			fmt.Printf("访问地址：http://localhost:%d\n", cfg.Server.Port)
		}
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serv.Stop(); err != nil {
			fmt.Println("停止服务失败：", err)
		} else {
			fmt.Println("服务已停止")
		}
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "以控制台模式运行",
	Run: func(cmd *cobra.Command, args []string) {
		initializationData()
		if err := serv.Run(); err != nil {
			fmt.Println("运行服务失败：", err)
		}
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "重启服务",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serv.Restart(); err != nil {
			fmt.Println("重启服务失败：", err)
		} else {
			cfg := config.GetConfig()
			fmt.Println("服务重启成功")
			fmt.Printf("访问地址：http://localhost:%d\n", cfg.Server.Port)
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "查看服务状态",
	Run: func(cmd *cobra.Command, args []string) {
		status, err := serv.Status()
		if err != nil {
			fmt.Println("获取服务状态失败：", err)
		} else {
			switch status {
			case 1:
				fmt.Println("服务状态：运行中")
			case 2:
				fmt.Println("服务状态：已停止")
			default:
				fmt.Println("服务状态：未知")
			}
		}
	},
}

func init() {
	flagsData(runCmd, fmt.Sprintf("%s 的数据路径，包括上传的文件、配置文件、数据库文件等", constants.WebTitle))
	flagsData(installCmd, fmt.Sprintf("%s 的数据路径，包括上传的文件、配置文件、数据库文件等", constants.WebTitle))

	runCmd.PreRun = preRunInitialization
	installCmd.PreRun = preRunInitialization

	toolboxCmd.AddCommand(installCmd)
	toolboxCmd.AddCommand(uninstallCmd)
	toolboxCmd.AddCommand(startCmd)
	toolboxCmd.AddCommand(stopCmd)
	toolboxCmd.AddCommand(runCmd)
	toolboxCmd.AddCommand(restartCmd)
	toolboxCmd.AddCommand(statusCmd)

	toolboxCmd.PersistentFlags().BoolP("help", "h", false, "查看命令帮助信息")
}
