package config

import (
	"strings"
	"sync"
	"toolbox/pkg/constants"
	"toolbox/pkg/utils"
	"toolbox/pkg/utils/file"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/ini.v1"
)

var (
	config *configs
	path   string
	ones   sync.Once
)

func GetConfig() *configs {
	ones.Do(func() {
		ifaces, err := utils.GetNetworkInterfaces()
		if err != nil {
			ifaces[0].Name = "以太网"
		}
		config = &configs{
			Server: server{
				Port:       6500,
				DNSRunning: constants.DNSStop,
				Iface:      ifaces[0].Name,
				UploadSize: 32 << 20,
			},
			Logger: logger{
				Level:      zerolog.InfoLevel.String(),
				MaxSize:    50,
				MaxAge:     7,
				MaxBackups: 5,
			},
		}

		// ini 覆盖
		path, _ = file.GetFileAbsPath(constants.DataPath, "config.ini")

		if _file, err := ini.Load(path); err == nil {
			_ = _file.MapTo(config)
			config.Server.UploadSize = config.Server.UploadSize << 20
		}

		if config.Logger.Level == "no" {
			config.Logger.Level = ""
		}
		config.Logger.Level = strings.ToLower(config.Logger.Level)
	})
	return config
}

// SaveToIni 保存配置到 ini 文件
func SaveToIni() {
	_file := ini.Empty()
	config = GetConfig()
	config.Server.UploadSize = config.Server.UploadSize >> 20
	err := _file.ReflectFrom(config)
	if err != nil {
		log.Fatal().Err(err).Msg("配置保存，序列化成ini失败")
	}

	if _file.SaveTo(path) != nil {
		log.Fatal().Err(err).Msg("配置文件保存失败")
	}
}
