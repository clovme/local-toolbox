package core

import (
	"strconv"
	"toolbox/pkg/logger/log"

	"github.com/gin-gonic/gin"
)

// Context 自定义gin.Context
type Context struct {
	*gin.Context
	Router routesMap
}

// Get 获取自定义gin.Context中的值
func (r *Context) Get(key string) any {
	if value, exists := r.Context.Get(key); exists {
		return value
	}
	return nil
}

// GetParamInt 获取路由参数中的int值
func (r *Context) GetParamInt(key string) int {
	paramValue := r.Context.Param(key)
	value, err := strconv.Atoi(paramValue)
	if err != nil {
		log.Error().Err(err).Msgf("core.ParamInt(%T(\"%s\"))转int失败", paramValue, paramValue)
		return 0
	}
	return value
}

// QueryInt64 获取路由参数中的int值
func (r *Context) QueryInt64(key string) int64 {
	paramValue := r.Context.Query(key)
	value, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		log.Error().Err(err).Msgf("core.QueryInt64(%T(\"%s\"))转int64失败", paramValue, paramValue)
		return 0
	}
	return value
}

// NewContext 创建自定义gin.Context
//
// 参数:
//   - ctx: gin.Context对象
//
// 返回值:
//   - *Context: 自定义gin.Context对象
//
// 说明:
//   - 创建自定义gin.Context对象，用于自定义路由和中间件。
func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context: ctx,
		Router:  initRoutesMap(),
	}
}
