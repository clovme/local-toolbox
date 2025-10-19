package category

import (
	"toolbox/internal/schema/dto/category"
	categoryVO "toolbox/internal/schema/vo/category"
)

type ApiCategoryService struct {
	Repo Repository
}

func (r *ApiCategoryService) ServiceFindCategoryAll() ([]*categoryVO.ApiCategoryVO, error) {
	return r.Repo.GetCategoryList()
}

func (r *ApiCategoryService) ServiceDeleteCategory(id int64) (int64, int64, error) {
	return r.Repo.DeleteCategory(id)
}

func (r *ApiCategoryService) ServiceAddCategory(dto category.ApiCategoryAddDTO) error {
	return r.Repo.AddCategory(dto)
}

func (r *ApiCategoryService) ServiceUpdateCategory(dto category.ApiCategoryAddDTO) error {
	return r.Repo.UpdateCategory(dto)
}
