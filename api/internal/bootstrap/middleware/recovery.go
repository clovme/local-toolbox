package middleware

import (
	"toolbox/internal/core"
	"toolbox/pkg/enums/code"
	httpLog "toolbox/pkg/logger/http"
)

// RecoveryMiddleware panic 捕捉中间件
func RecoveryMiddleware() core.HandlerFunc {
	return func(c *core.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 错误，附带 stack trace
				httpLog.Panic(c.Context).Interface("panic", err).Msg("捕捉到请求异常")

				// 返回统一处理
				c.JsonDesc(code.ServerInternalError, nil)

				// 强制中断后续
				c.AbortWithStatus(500)
			}
		}()
		c.Next()
	}
}
