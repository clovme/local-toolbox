package initdata

import (
	"toolbox/internal/models"
	"toolbox/pkg/logger/log"
)

func (r *InitData) Category() {
	m := r.Q.Category

	modelList := []models.Category{
		{Pid: 0, Title: "全部文章", Name: "all", Description: "全部文章", Sort: 0},
		{Pid: 0, Title: "默认分类", Name: "default", Description: "默认分类", Sort: 1},
	}

	newModelList := insertIfNotExist[models.Category](modelList, func(model models.Category) (*models.Category, error) {
		return m.Where(m.Name.Eq(model.Name)).Take()
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
