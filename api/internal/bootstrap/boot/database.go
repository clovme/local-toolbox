package boot

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
	"toolbox/internal/infrastructure/query"
	"toolbox/pkg/constants"
	"toolbox/pkg/logger"
	"toolbox/pkg/logger/log"
	"toolbox/pkg/utils/file"
)

// InitializationDB 打开数据库连接, 并设置连接池, 数据库链接统一入口
// 返回值：
//   - *gorm.DB 数据库连接对象
func InitializationDB() *gorm.DB {
	dbPath, err := file.GetFileAbsPath(constants.DataPath, "data.db")
	if err != nil {
		return nil
	}
	var dsn = sqlite.Open(dbPath)

	db, err := gorm.Open(dsn, logger.GetGormLogger())
	if err != nil {
		log.Error().Err(err).Msg("打开SQLite失败")
		os.Exit(-1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Error().Err(err).Msg("获取底层 sql.DB 失败")
		os.Exit(-1)
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	query.SetDefault(db)
	return db
}
