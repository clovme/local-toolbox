package article

import (
	"toolbox/internal/core"
	articleDTO "toolbox/internal/schema/dto/article"
	articleVO "toolbox/internal/schema/vo/article"
)

type ApiArticleService struct {
	Repo Repository
}

func (r *ApiArticleService) ServiceDeleteArticle(dto articleDTO.ApiArticleUpdateDTO) error {
	return r.Repo.DeleteArticle(dto)
}

func (r *ApiArticleService) ServiceUpdateArticle(dto articleDTO.ApiArticleDTO) error {
	return r.Repo.UpdateArticle(dto)
}

func (r *ApiArticleService) ServiceGetArticle(c *core.Context) (data *articleVO.ApiArticleVO, err error) {
	return r.Repo.GetArticle(c)
}

func (r *ApiArticleService) ServiceGetArticleList(c *core.Context) ([]articleVO.ApiArticleVO, error) {
	return r.Repo.GetArticleList(c)
}

func (r *ApiArticleService) ServiceAddArticle(dto articleDTO.ApiArticleDTO) (int64, error) {
	return r.Repo.AddArticle(dto)
}

func (r *ApiArticleService) ServiceUploadImages(c *core.Context) (*articleVO.ApiMdImagesVO, error) {
	return r.Repo.UploadImages(c)
}
