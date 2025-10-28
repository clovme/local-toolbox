package category

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ApiCategoryAddDTO struct {
	ID      *int64 `json:"id,string"`
	Title   string `json:"title" binding:"required"`
	Pid     *int64 `json:"pid,string"`
	Sort    int    `json:"sort"`
	DocSort string `json:"docSort"`
}

func (r ApiCategoryAddDTO) TranslateError(err error) map[string]string {
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
		case "name":
			switch tag {
			case "required":
				result["name"] = "分类名称不能为空"
			}
		}
	}
	return result
}
