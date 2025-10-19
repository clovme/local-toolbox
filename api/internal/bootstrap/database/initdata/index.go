package initdata

import (
	"errors"
	"gorm.io/gorm"
	"toolbox/internal/core"
	"toolbox/internal/infrastructure/query"
)

type InitData struct {
	Router []core.RoutesInfo
	Q      *query.Query
}

// insertIfNotExist 插入数据
func insertIfNotExist[T any](modelList []T, exists func(model T) (*T, error)) []*T {
	newModelList := make([]*T, 0)

	for _, model := range modelList {
		if _, err := exists(model); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				newModelList = append(newModelList, &model)
				continue
			}
		}
	}

	return newModelList
}
