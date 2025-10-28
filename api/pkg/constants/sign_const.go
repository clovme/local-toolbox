// 常量定义，常用标志定义, 用于上下文传递，全局使用, 不建议修改
// 用于项目中常使用的一些通用键(签名)，用于设置值和取值

package constants

import "time"

const (
	ProjectName = "toolbox" // 项目名称
	WebTitle    = "本地系统工具箱"

	HttpLogKey = "HTTP_LOG_KEY"
	LimitPage  = "LIMIT_PAGE"

	DNSStop    = "stop"
	DNSRunning = "running"
)

var (
	DataPath   = "."
	UploadPath = "upload"
	RunTime    = time.Now()
)
