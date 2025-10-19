package cmd

import (
	"toolbox/internal/bootstrap/boot"
	"toolbox/internal/bootstrap/routers"
	"toolbox/pkg/config"

	"github.com/kardianos/service"
)

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Start the work in a goroutine.
	p.exit = make(chan struct{})
	go p.run()
	return nil
}

func (p *program) run() {
	// 初始化系统日志
	boot.InitializationLogger()

	// 初始化表单验证器
	boot.InitializationFormValidate()

	// 连接数据库
	db := boot.InitializationDB()

	cfg := config.GetConfig()

	// 初始化路由
	engine := routers.Initialization(db, cfg.Server.UploadSize)
	engine.Run("localhost", cfg.Server.Port)
}

func (p *program) Stop(s service.Service) error {
	// Stop 应尽快返回，通知 goroutine 退出
	close(p.exit)
	return nil
}
