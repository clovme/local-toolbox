package api

import (
	"fmt"
	"strconv"
	articleService "toolbox/internal/application/article"
	"toolbox/internal/core"
	"toolbox/internal/schema/dto/article"
	"toolbox/pkg/enums/code"
	"toolbox/public"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	Service *articleService.ApiArticleService
}

// ArticleHandler
// @Type			api
// @Group 			articleApi
// @Router			/readme [GET]
// @Name			article
// @Summary			获取Md说明文档
func (r *ArticleHandler) ArticleHandler(c *core.Context) {
	c.JsonDesc(code.Success, string(public.Readme))
}

// PostUploadImagesHandler
// @Type			api
// @Group 			articleApi
// @Router			/upload/images [POST]
// @Name			uploadImages
// @Summary			图片上传
func (r *ArticleHandler) PostUploadImagesHandler(c *core.Context) {
	recordVO, err := r.Service.ServiceUploadImages(c)
	if err != nil {
		c.JsonFail(code.Fail, err.Error(), nil)
		return
	}
	c.JsonSuccess(recordVO)
}

// GetArticleListHandler
// @Type			api
// @Group 			articleApi
// @Router			/article/list [GET]
// @Name			articleListGet
// @Summary			获取文章列表
func (r *ArticleHandler) GetArticleListHandler(c *core.Context) {
	data, err := r.Service.ServiceGetArticleList(c)
	if err != nil {
		c.JsonFail(code.Fail, err.Error(), nil)
		return
	}
	c.Limit(20, 1, 8).JsonSuccess(data)
}

// GetArticleHandler
// @Type			api
// @Group 			articleApi
// @Router			/article [GET]
// @Name			articleGet
// @Summary			获取文章列表
func (r *ArticleHandler) GetArticleHandler(c *core.Context) {
	data, err := r.Service.ServiceGetArticle(c)
	if err != nil {
		c.JsonFail(code.Fail, err.Error(), nil)
		return
	}
	c.JsonSuccess(data)
}

// PostArticleHandler
// @Type			api
// @Group 			articleApi
// @Router			/article [POST]
// @Name			articleAdd
// @Summary			添加文章
func (r *ArticleHandler) PostArticleHandler(c *core.Context) {
	var dto article.ApiArticleDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JsonFailData(code.Fail, "请求参数绑定失败", err, dto.TranslateError(err))
		return
	}

	id, err := r.Service.ServiceAddArticle(dto)
	if err != nil {
		c.JsonFail(code.Fail, err.Error(), nil)
		return
	}
	c.JsonSuccess(gin.H{
		"id": strconv.FormatInt(id, 10),
	})
}

// PutArticleHandler
// @Type			api
// @Group 			articleApi
// @Router			/article [PUT]
// @Name			articleUpdate
// @Summary			添加文章
func (r *ArticleHandler) PutArticleHandler(c *core.Context) {
	var dto article.ApiArticleDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JsonFailData(code.Fail, "请求参数绑定失败", err, dto.TranslateError(err))
		return
	}

	if err := r.Service.ServiceUpdateArticle(dto); err != nil {
		c.JsonFail(code.Fail, err.Error(), nil)
		return
	}
	c.Json(code.Success, fmt.Sprintf("[%s]更新成功！", dto.Title), nil)
}

// DeleteArticleHandler
// @Type			api
// @Group 			articleApi
// @Router			/article [DELETE]
// @Name			articleDelete
// @Summary			添加文章
func (r *ArticleHandler) DeleteArticleHandler(c *core.Context) {
	var dto article.ApiArticleUpdateDTO
	if err := c.ShouldBind(&dto); err != nil {
		c.JsonFailData(code.Fail, "请求参数绑定失败", err, dto.TranslateError(err))
		return
	}

	if err := r.Service.ServiceDeleteArticle(dto); err != nil {
		c.JsonFail(code.Fail, err.Error(), nil)
		return
	}
	c.Json(code.Success, fmt.Sprintf("[%s]删除成功！", dto.Title), nil)
}
