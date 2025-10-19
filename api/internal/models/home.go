package models

import (
	"time"
	"toolbox/pkg/enums/icon"
	"toolbox/pkg/enums/position"
	"toolbox/pkg/utils"

	"gorm.io/gorm"
)

// Home 首页
type Home struct {
	ID        int64             `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:ID"`
	Icon      string            `gorm:"type:varchar(20);not null;comment:图标"`
	Name      string            `gorm:"type:varchar(10);not null;comment:名称"`
	Path      string            `gorm:"type:varchar(10);not null;comment:路径"`
	Sort      int               `gorm:"type:int;default:0;comment:排序值，值越大越靠前，默认0"`
	IconType  icon.Icon         `gorm:"type:varchar(20);not null;comment:图标类型"`
	Position  position.Position `gorm:"type:varchar(20);not null;comment:位置类型"`
	CreatedAt *time.Time        `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt *time.Time        `gorm:"autoUpdateTime:nano;comment:更新时间"`
}

// BeforeCreate 执行 gorm 创建前操作
// 自动生成ID
func (r *Home) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

// TableName 表名
func (*Home) TableName() string {
	return "home"
}

// TableComment 表注释
func (r *Home) TableComment() string {
	return "首页"
}
