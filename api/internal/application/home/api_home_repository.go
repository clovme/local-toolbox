package home

import (
	homeVO "toolbox/internal/schema/vo/home"
)

type Repository interface {
	GetHomeIconData() ([]*homeVO.ApiHomeVO, error)
}
