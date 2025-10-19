package category

import (
	"toolbox/internal/schema/dto/category"
	categoryVO "toolbox/internal/schema/vo/category"
)

type Repository interface {
	GetCategoryList() ([]*categoryVO.ApiCategoryVO, error)
	DeleteCategory(id int64) (int64, int64, error)
	AddCategory(dto category.ApiCategoryAddDTO) error
	UpdateCategory(dto category.ApiCategoryAddDTO) error
}
