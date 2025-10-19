package api

import (
	"fmt"
	categoryService "toolbox/internal/application/category"
	"toolbox/internal/core"
	"toolbox/internal/schema/dto/category"
	"toolbox/pkg/enums/code"
)

type CategoryHandler struct {
	Service *categoryService.ApiCategoryService
}

// GetCategoryHandler
// @Type			api
// @Group 			categoryApi
// @Router			/category [GET]
// @Name			categoryGet
// @Summary			获取分类数据
func (r *CategoryHandler) GetCategoryHandler(c *core.Context) {
	data, err := r.Service.ServiceFindCategoryAll()
	if err != nil {
		c.JsonFail(code.ServiceQueryError, "没有获取到数据", err)
		return
	}

	c.JsonDesc(code.Success, data)
}

// DeleteCategoryHandler
// @Type			api
// @Group 			categoryApi
// @Router			/category [DELETE]
// @Name			categoryDelete
// @Summary			删除指定分类
func (r *CategoryHandler) DeleteCategoryHandler(c *core.Context) {
	var p category.DeleteCategoryDTO
	if err := c.ShouldBind(&p); err != nil {
		c.JsonFail(code.Fail, "请求参数绑定失败", err)
		return
	}

	del, update, err := r.Service.ServiceDeleteCategory(*p.ID)
	if err != nil {
		c.JsonFail(code.ServiceDeleteError, "数据删除错误", err)
		return
	}
	c.Json(code.Success, fmt.Sprintf("删除[%d]个分类，更新[%d]偏文章到默认分类", del, update), nil)
}

// PostCategoryHandler
// @Type			api
// @Group 			categoryApi
// @Router			/category [POST]
// @Name			categoryAdd
// @Summary			新增分类
func (r *CategoryHandler) PostCategoryHandler(c *core.Context) {
	var dto category.ApiCategoryAddDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JsonFail(code.Fail, "请求参数绑定失败", err)
		return
	}

	if err := r.Service.ServiceAddCategory(dto); err != nil {
		c.JsonFail(code.ServiceDeleteError, err.Error(), err)
		return
	}
	c.JsonSuccess(nil)
}

// PutCategoryHandler
// @Type			api
// @Group 			categoryApi
// @Router			/category [PUT]
// @Name			categoryUpdate
// @Summary			新增分类
func (r *CategoryHandler) PutCategoryHandler(c *core.Context) {
	var dto category.ApiCategoryAddDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JsonFail(code.Fail, "请求参数绑定失败", err)
		return
	}

	if err := r.Service.ServiceUpdateCategory(dto); err != nil {
		c.JsonFail(code.ServiceDeleteError, err.Error(), err)
		return
	}
	c.JsonSuccess(nil)
}
