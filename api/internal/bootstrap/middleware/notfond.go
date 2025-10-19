package middleware

import (
	"toolbox/internal/core"
	"toolbox/pkg/enums/code"
	httpLog "toolbox/pkg/logger/http"
)

// RegisterNoRoute 注册404处理
func RegisterNoRoute(engine *core.Engine) {
	engine.NoRoute(func(c *core.Context) {
		httpLog.Error(c.Context).Msg("请求地址错误")
		c.JsonDesc(code.RequestNotFound, nil)
		c.AbortWithStatus(404)
	})
}
