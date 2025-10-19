package models

import (
	"time"
	"toolbox/pkg/enums/status"
	"toolbox/pkg/utils"

	"gorm.io/gorm"
)

// DNSTable DNS代理配置项表
type DNSTable struct {
	ID        int64         `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:ID"`
	Protocol  string        `gorm:"type:varchar(10);not null;comment:协议"`
	Domain    string        `gorm:"type:varchar(50);not null;unique;comment:域名"`
	IP        string        `gorm:"type:varchar(20);not null;comment:IP地址"`
	Port      string        `gorm:"type:varchar(6);default:53;comment:端口"`
	Status    status.Status `gorm:"default:1;comment:状态：Enable启用，Disable禁用"`
	CreatedAt *time.Time    `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt *time.Time    `gorm:"autoUpdateTime:nano;comment:更新时间"`
}

// BeforeCreate 执行 gorm 创建前操作
// 自动生成ID
func (r *DNSTable) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

// TableName 表名
func (*DNSTable) TableName() string {
	return "dns"
}

// TableComment 表注释
func (r *DNSTable) TableComment() string {
	return "DNS代理配置表"
}
