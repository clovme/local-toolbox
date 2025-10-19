package initdata

import (
	"toolbox/internal/models"
	"toolbox/pkg/enums/icon"
	"toolbox/pkg/enums/position"
	"toolbox/pkg/logger/log"
)

func (r *InitData) Sidebar() {
	m := r.Q.Home

	modelList := []models.Home{
		{Icon: "home", Name: "首页", Sort: 0, IconType: icon.Font, Position: position.Sidebar},
		{Icon: "dns", Name: "DNS管理", Sort: 0, IconType: icon.Image, Position: position.Body},
		{Icon: "article", Name: "文章列表", Sort: 1, IconType: icon.Image, Position: position.Body},
	}

	newModelList := insertIfNotExist[models.Home](modelList, func(model models.Home) (*models.Home, error) {
		return m.Where(m.Name.Eq(model.Name), m.Icon.Eq(model.Icon)).Take()
	})

	if len(newModelList) <= 0 {
		return
	}

	if err := m.CreateInBatches(newModelList, 100); err != nil {
		log.Error().Err(err).Msg("[首页]初始化失败！")
	} else {
		log.Info().Msgf("[首页]初始化成功，共%d条数据！", len(newModelList))
	}
}
