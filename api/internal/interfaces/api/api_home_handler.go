package api

import (
	homeService "toolbox/internal/application/home"
	"toolbox/internal/core"
	"toolbox/pkg/enums/code"
)

type HomeHandler struct {
	Service *homeService.ApiHomeService
}

// HomeDataHandler
// @Type			api
// @Group 			homeApi
// @Router			/home/data [GET]
// @Name			homeGet
// @Summary			获取首页数据
func (r *HomeHandler) HomeDataHandler(c *core.Context) {
	iconData, err := r.Service.ServiceFindHomeIconData()
	if err != nil {
		c.JsonFailDesc(code.ServiceQueryError, err)
		return
	}
	c.JsonSuccess(iconData)
}
