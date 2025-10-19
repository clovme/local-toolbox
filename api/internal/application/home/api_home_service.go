package home

import (
	homeVO "toolbox/internal/schema/vo/home"
)

type ApiHomeService struct {
	Repo Repository
}

func (r *ApiHomeService) ServiceFindHomeIconData() ([]*homeVO.ApiHomeVO, error) {
	return r.Repo.GetHomeIconData()
}
