package middleware

import (
	"io/fs"
	"net/http"
	"strings"
	"toolbox/internal/core"
	"toolbox/pkg/constants"
	"toolbox/pkg/enums/code"
	"toolbox/public"

	"github.com/gin-gonic/gin"
)

// FaviconMiddleware 加载 favicon.ico
func FaviconMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/favicon.ico" {
			c.Data(200, "image/x-icon", public.Favicon)
			c.Set(constants.HttpLogKey, "favicon")
			c.Abort()
			return
		}
		c.Next()
	}
}

// ResourceDirInterception 静态路由拦截
func ResourceDirInterception(engine *core.Engine) {
	// 静态资源路由拦截中间件处理，中间件部分
	engine.Use(func(c *core.Context) {
		if strings.HasSuffix(c.FullPath(), "*filepath") {
			c.Set(constants.HttpLogKey, "静态资源")
		}
		if strings.EqualFold(c.Request.URL.Path, "/assets/") {
			c.JsonDesc(code.RequestNotFound, nil)
			c.AbortWithStatus(404)
		}
	})

	// 读取嵌入二进制的静态资源目录
	staticFS, _ := fs.Sub(public.WebFS, "web/assets")
	engine.Engine.StaticFS("/assets", http.FS(staticFS))
	engine.Engine.Static("/uploads", constants.UploadPath)
}
