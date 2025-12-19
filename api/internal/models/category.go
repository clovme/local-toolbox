package models

import (
	"time"
	"toolbox/pkg/utils"

	"gorm.io/gorm"
)

// Category Category
type Category struct {
	ID        int64      `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:ID"`
	Name      string     `gorm:"type:varchar(10);comment:分类名称(英文)"`
	Title     string     `gorm:"type:varchar(50);comment:分类名称(中文)"`
	DocSort   string     `gorm:"type:varchar(15);default:updatedAt;comment:排序字段"`
	Pid       int64      `gorm:"type:bigint;comment:父ID"`
	IsExpand  bool       `gorm:"default:0;comment:展开"`
	Sort      int        `gorm:"type:int;default:0;comment:排序值，值越大越靠前，默认0"`
	CreatedAt *time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:nano;comment:更新时间"`
}

// BeforeCreate 执行 gorm 创建前操作
// 自动生成ID
func (r *Category) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

// TableName 表名
func (*Category) TableName() string {
	return "category"
}

// TableComment 表注释
func (r *Category) TableComment() string {
	return "Category"
}
