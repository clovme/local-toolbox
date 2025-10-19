package article

import (
	"toolbox/internal/core"
	articleDTO "toolbox/internal/schema/dto/article"
	"toolbox/internal/schema/vo/article"
	articleVO "toolbox/internal/schema/vo/article"
)

type Repository interface {
	UploadImages(c *core.Context) (*article.ApiMdImagesVO, error)
	GetArticleList(c *core.Context) (data []articleVO.ApiArticleVO, err error)
	GetArticle(c *core.Context) (data *articleVO.ApiArticleVO, err error)
	AddArticle(dto articleDTO.ApiArticleDTO) (int64, error)
	UpdateArticle(dto articleDTO.ApiArticleDTO) error
	DeleteArticle(dto articleDTO.ApiArticleUpdateDTO) error
}
