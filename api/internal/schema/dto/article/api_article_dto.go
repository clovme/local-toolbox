package article

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiArticleDTO struct {
	ID           *int64 `json:"id,string" binding:"required"`
	Title        string `json:"title" binding:"required"`
	CategoryName string `json:"categoryName"`
	CategoryID   *int64 `json:"categoryID,string" binding:"required"`
	Tags         string `json:"tags" binding:"required"`
	Summary      string `json:"summary" binding:"required"`
	Content      string `json:"content" binding:"required"`
}

func (r ApiArticleDTO) TranslateError(err error) map[string]string {
	result := make(map[string]string)

	// 判断是不是验证错误
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		result["other"] = err.Error()
		return result
	}

	// 遍历所有验证错误
	for _, e := range errs {
		field := e.Field()
		tag := e.Tag()

		switch field {
		case "ID":
			switch tag {
			case "required":
				result["id"] = "ID不能为空"
			}
		case "Title":
			switch tag {
			case "required":
				result["title"] = "文章标题不能为空"
			}
		case "CategoryID":
			switch tag {
			case "required":
				result["categoryID"] = "文章分类ID不能为空"
			}
		case "Tags":
			switch tag {
			case "required":
				result["tags"] = "文章标签不能为空"
			}
		case "Summary":
			switch tag {
			case "required":
				result["summary"] = "文章摘要不能为空"
			}
		case "Content":
			switch tag {
			case "required":
				result["content"] = "文章内容不能为空"
			}
		}
	}
	return result
}
