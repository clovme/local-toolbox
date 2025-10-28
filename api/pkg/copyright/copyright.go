package copyright

import (
	"fmt"
	"os"
	"runtime"
	"time"
	"toolbox/pkg/constants"
	"toolbox/version"
)

// _copyright 版权信息
type _copyright struct {
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	GoVersion string    `json:"goVersion"`
	Platform  string    `json:"platform"`
	Pid       int       `json:"pid"`
	BuildTime string    `json:"buildTime"`
	RunTime   time.Time `json:"runTime"`
	NowTime   time.Time `json:"nowTime"`
}

func NewCopyright() _copyright {
	time.Local = time.UTC
	return _copyright{
		Name:      constants.WebTitle,
		Version:   version.Version,
		GoVersion: runtime.Version(),
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
		Pid:       os.Getpid(),
		BuildTime: version.BuildTime,
		RunTime:   constants.RunTime,
		NowTime:   time.Now(),
	}
}
