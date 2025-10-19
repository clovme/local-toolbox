package persistence

import (
	"gorm.io/gorm"
	"toolbox/internal/infrastructure/query"
	homeVO "toolbox/internal/schema/vo/home"
)

type ApiHomeRepository struct {
	DB *gorm.DB
	Q  *query.Query
}

func (r *ApiHomeRepository) GetHomeIconData() ([]*homeVO.ApiHomeVO, error) {
	var results []*homeVO.ApiHomeVO
	err := r.Q.Home.Order(r.Q.Home.Position.Desc(), r.Q.Home.Sort.Asc()).Scan(&results)
	if err != nil {
		return nil, err
	}
	return results, err
}
