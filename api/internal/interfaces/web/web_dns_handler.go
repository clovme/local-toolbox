package web

import (
	"strings"
	dnsService "toolbox/internal/application/dns"
	"toolbox/internal/core"
	"toolbox/internal/schema/dto/dns"
	"toolbox/pkg/config"
	"toolbox/pkg/constants"
	"toolbox/pkg/copyright"
	"toolbox/pkg/enums/code"
	"toolbox/pkg/enums/icon"
	"toolbox/pkg/enums/position"
	"toolbox/pkg/enums/status"
	"toolbox/pkg/utils"

	"github.com/gin-gonic/gin"
)

type DnsHandler struct {
	Service *dnsService.WebDnsService
}

// GetViewsIndexHandler
// @Type			web
// @Group 			dnsView
// @Router			/ [GET]
// @Name			indexView
// @Summary			首页视图
func (r *DnsHandler) GetViewsIndexHandler(c *core.Context) {
	c.HTML("index.html", nil)
}

// GetEnumsHandler
// @Type			api
// @Group 			dnsApi
// @Router			/enums [GET]
// @Name			enumsMap
// @Summary			获取枚举映射
func (r *DnsHandler) GetEnumsHandler(c *core.Context) {
	c.JsonSuccess(gin.H{
		"code":     code.ValueMap(),
		"status":   status.ValueMap(),
		"icon":     icon.ValueMap(),
		"position": position.ValueMap(),
	})
}

// CopyrightHandler
// @Type			api
// @Group 			dnsApi
// @Router			/copyright [GET]
// @Name			copyright
// @Summary			版权
func (r *DnsHandler) CopyrightHandler(c *core.Context) {
	c.JsonSuccess(copyright.NewCopyright())
}

// PageHandler
// @Type			api
// @Group 			dnsApi
// @Router			/list [GET]
// @Name			dnsList
// @Summary			获取DNS列表
func (r *DnsHandler) PageHandler(c *core.Context) {
	query, ok := c.GetQuery("orderBy")
	if !ok {
		query = "createdAt|asc"
	}
	dataList, err := r.Service.ServiceFindDnsList(query)
	if err != nil {
		c.JsonFail(code.ServiceQueryError, "获取DNS列表失败", err)
		return
	}
	c.JsonSuccess(dataList)
}

// SaveHandler
// @Type			api
// @Group 			dnsApi
// @Router			/save [POST]
// @Name			dnsSave
// @Summary			保存DNS数据
func (r *DnsHandler) SaveHandler(c *core.Context) {
	var data dns.WebSaveData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JsonFailDesc(code.ServiceInsertError, err)
		return
	}
	c.JsonSuccess(r.Service.ServiceSaveDnsData(data))
}

// DeleteHandler
// @Type			api
// @Group 			dnsApi
// @Router			/delete [DELETE]
// @Name			dnsDelete
// @Summary			删除DNS数据
func (r *DnsHandler) DeleteHandler(c *core.Context) {
	var row dns.WebDeleteRow
	if err := c.ShouldBindJSON(&row); err != nil {
		c.JsonFailDesc(code.ServiceDeleteError, err)
		return
	}
	if err := r.Service.ServiceDeleteDnsData(row.ID); err != nil {
		c.JsonFailDesc(code.ServiceDeleteError, err)
		return
	}
	c.JsonSuccess("数据删除成功")
}

// ServiceRunningHandler
// @Type			api
// @Group 			dnsApi
// @Router			/service/running/{iface} [POST]
// @Name			dnsServiceRunning
// @Summary			启动DNS服务
func (r *DnsHandler) ServiceRunningHandler(c *core.Context) {
	cfg := config.GetConfig()
	if strings.EqualFold(cfg.Server.DNSRunning, constants.DNSRunning) {
		c.JsonDnsStatus(constants.DNSRunning)
		return
	}
	cfg.Server.Iface = c.Param("iface")
	cfg.Server.DNSRunning = constants.DNSRunning

	dnsProxy, err := r.Service.ServiceNewDNSProxy()
	if err != nil {
		c.JsonFailDesc(code.RequestNotFound, err)
		return
	}
	if err := dnsProxy.SetLocalDNS(cfg.Server.Iface); err != nil {
		c.JsonFail(code.RequestNotFound, "设置系统DNS(127.0.0.1)失败", err)
		return
	}

	go dnsProxy.StartDnsServer()
	<-dnsProxy.CtxStatus.Done()

	if !dnsProxy.IsRunning {
		c.JsonDnsStatus(constants.DNSStop)
		return
	}

	c.JsonDnsStatus(constants.DNSRunning)
	config.SaveToIni()
}

// ServiceStopHandler
// @Type			api
// @Group 			dnsApi
// @Router			/service/stop/{iface} [POST]
// @Name			dnsServiceStop
// @Summary			禁用DNS服务
func (r *DnsHandler) ServiceStopHandler(c *core.Context) {
	cfg := config.GetConfig()

	if strings.EqualFold(cfg.Server.DNSRunning, constants.DNSStop) {
		c.JsonDnsStatus(constants.DNSStop)
		return
	}
	cfg.Server.DNSRunning = constants.DNSStop
	cfg.Server.Iface = c.Param("iface")

	dnsProxy, err := r.Service.ServiceNewDNSProxy()
	if err != nil {
		c.JsonFail(code.RequestNotFound, "获取网卡列表失败", err)
		return
	}
	if err := dnsProxy.RestoreDNS(cfg.Server.Iface); err != nil {
		c.JsonFail(code.RequestNotFound, "系统代理DNS设置失败", err)
		return
	}

	c.JsonDnsStatus(constants.DNSStop)
	config.SaveToIni()
}

// GetNetIfaceHandler
// @Type			api
// @Group 			dnsApi
// @Router			/network/interfaces [GET]
// @Name			dnsNetIface
// @Summary			获取网络接口列表
func (r *DnsHandler) GetNetIfaceHandler(c *core.Context) {
	ifaces, err := utils.GetNetworkInterfaces()
	if err != nil {
		c.JsonFail(code.ServiceQueryError, "获取网卡列表失败", err)
		return
	}
	cfg := config.GetConfig()
	c.JsonSuccess(gin.H{
		"iface":   cfg.Server.Iface,
		"running": cfg.Server.DNSRunning,
		"ifaces":  ifaces,
	})
}
