package core

import (
	"fmt"
	"net/http"
	"toolbox/pkg/constants"
	"toolbox/pkg/enums/code"
	"toolbox/pkg/logger/log"

	"github.com/gin-gonic/gin"
)

type page struct {
	PageSize    int64 `json:"pageSize"`
	CurrentPage int64 `json:"currentPage"`
	Total       int64 `json:"total"`
}

// response 响应结构体
type response struct {
	Code    code.ResponseCode `json:"code"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
	Page    interface{}       `json:"page,omitempty"`
}

// Limit 设置分页信息
// 参数：
//   - pageSize: 每页数量
//   - currentPage: 当前页码
//   - total: 总记录数
//
// 返回值：
//   - *Context: 自定义gin.Context对象
func (r *Context) Limit(pageSize, currentPage, total int64) *Context {
	r.Context.Set(constants.LimitPage, &page{
		PageSize:    pageSize,
		CurrentPage: currentPage,
		Total:       total,
	})
	return r
}

// Json 响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - message: 响应消息
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context) Json(httpCode code.ResponseCode, message string, data interface{}) {
	r.Context.JSON(http.StatusOK, response{
		Code:    httpCode,
		Message: fmt.Sprintf("[%d] %s", httpCode.Int(), message),
		Data:    data,
		Page:    r.Get(constants.LimitPage),
	})
}

// JsonDesc Json响应
// 参数：
//   - c: gin.Context
//   - code: 响应码
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context) JsonDesc(httpCode code.ResponseCode, data interface{}) {
	r.Json(httpCode, httpCode.Desc(), data)
}

// JsonSuccess 响应成功
// 参数：
//   - data: 响应数据
//
// 返回值：
//   - 无
func (r *Context) JsonSuccess(data interface{}) {
	r.Json(code.Success, code.Success.Desc(), data)
}

// JsonFailDesc 响应失败
// 参数：
//   - httpCode: 响应码
//   - err: 错误信息
//
// 返回值：
//   - 无
func (r *Context) JsonFailDesc(httpCode code.ResponseCode, err error) {
	log.Error().Err(err).Msgf(httpCode.Desc())
	r.Json(httpCode, httpCode.Desc(), nil)
}

// JsonFail 响应失败
// 参数：
//   - httpCode: 响应码
//   - msg: 错误消息
//   - err: 错误信息
//
// 返回值：
//   - 无
func (r *Context) JsonFail(httpCode code.ResponseCode, msg string, err error) {
	log.Error().Err(err).Msgf(msg)
	r.Json(httpCode, msg, nil)
}

// JsonFailData 响应失败
// 参数：
//   - httpCode: 响应码
//   - msg: 错误消息
//   - err: 错误信息
//   - data: 数据
//
// 返回值：
//   - 无
func (r *Context) JsonFailData(httpCode code.ResponseCode, msg string, err error, data interface{}) {
	log.Error().Err(err).Interface("错误信息", data).Msgf(msg)
	r.Json(httpCode, msg, data)
}

// JsonDnsStatus 响应DNS状态
// 参数：
//   - c: gin.Context
//   - running: DNS运行状态
//
// 返回值：
//   - 无
func (r *Context) JsonDnsStatus(running string) {
	r.JsonSuccess(gin.H{
		"running": running,
	})
}
