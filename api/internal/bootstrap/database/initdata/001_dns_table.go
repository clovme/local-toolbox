package initdata

import (
	"strconv"
	"toolbox/internal/models"
	"toolbox/pkg/config"
	"toolbox/pkg/enums/status"
	"toolbox/pkg/logger/log"
)

// DNSTable 初始化配置
func (r *InitData) DNSTable() {
	m := r.Q.DNSTable
	cfg := config.GetConfig()

	modelList := []models.DNSTable{
		{Protocol: "http", Domain: "localhost", IP: "127.0.0.1", Port: strconv.Itoa(cfg.Server.Port), Status: status.Enable},
		{Protocol: "http", Domain: "www.localhost.com", IP: "127.0.0.1", Port: strconv.Itoa(cfg.Server.Port), Status: status.Enable},
	}

	newModelList := insertIfNotExist[models.DNSTable](modelList, func(model models.DNSTable) (*models.DNSTable, error) {
		return m.Where(m.Domain.Eq(model.Domain)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := m.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[DNS表]初始化失败！")
	} else {
		log.Info().Msgf("[DNS表]初始化成功，共%d条数据！", len(newModelList))
	}
}
