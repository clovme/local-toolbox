package models

import (
	"time"
	"toolbox/pkg/utils"

	"gorm.io/gorm"
)

// FileRecord Files
type FileRecord struct {
	ID        int64      `gorm:"primaryKey;type:bigint;autoIncrement:false;comment:ID"`
	Url       string     `gorm:"type:varchar(100);comment:文件Url"`
	Name      string     `gorm:"type:varchar(50);comment:文件名"`
	Type      string     `gorm:"type:varchar(20);comment:文件类型"`
	Size      int64      `gorm:"type:bigint;comment:文件大小"`
	Hash      string     `gorm:"type:varchar(64);unique;comment:文件哈希"`
	CreatedAt *time.Time `gorm:"autoCreateTime:nano;comment:创建时间"`
}

// BeforeCreate 执行 gorm 创建前操作
// 自动生成ID
func (r *FileRecord) BeforeCreate(tx *gorm.DB) (err error) {
	if r.ID == 0 {
		r.ID = utils.GenerateID()
	}
	return
}

// TableName 表名
func (*FileRecord) TableName() string {
	return "file_record"
}

// TableComment 表注释
func (r *FileRecord) TableComment() string {
	return "文件信息表"
}
