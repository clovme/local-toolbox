package models

import (
	"time"
	"toolbox/pkg/utils"

	"gorm.io/gorm"
)

// Article Article
type Article struct {
	ID         int64      `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:ID"`
	Title      string     `gorm:"type:varchar(100);comment:文章标题"`
	CategoryID int64      `gorm:"type:bigint;comment:文章分类"`
	Tags       string     `gorm:"type:varchar(100);comment:标签，逗号分割"`
	Summary    string     `gorm:"type:varchar(300);comment:摘要信息"`
	Content    string     `gorm:"type:text;comment:文章内容"`
	CreatedAt  *time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
	UpdatedAt  *time.Time `gorm:"autoUpdateTime:nano;comment:更新时间"`
}

// BeforeCreate 执行 gorm 创建前操作
// 自动生成ID
func (r *Article) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

// TableName 表名
func (*Article) TableName() string {
	return "article"
}

// TableComment 表注释
func (r *Article) TableComment() string {
	return "Article"
}
