package model

import (
	"database/sql"
	"fmt"
	"github.com/endymion/go-base/task-04/common/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var dataBaseSetting = setting.DatabaseSetting

var db *gorm.DB
var sqlDB *sql.DB

func SetUp() {
	//模版字符串构建dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		dataBaseSetting.User,
		dataBaseSetting.Password,
		dataBaseSetting.Host,
		dataBaseSetting.Port,
		dataBaseSetting.Name,
		dataBaseSetting.Conf)

	// 根据databaseSetting.Type 使用gorm 连接数据库
	var err error
	switch dataBaseSetting.Type {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             200 * time.Millisecond,
					LogLevel:                  logger.Info, // 输出所有 SQL
					IgnoreRecordNotFoundError: true,
					Colorful:                  true,
				},
			),
		})
	default:
		log.Fatalf("数据库类型错误, %s", dataBaseSetting.Type)
	}
	if err != nil {
		log.Fatal("数据库连接失败")
	}
	if setting.ServerSetting.RunMode == "debug" {
		db = db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Info)})
	}
	sqlDB, err = db.DB()
	if err != nil {
		log.Fatal("数据库连接失败")
	}
	sqlDB.SetMaxIdleConns(dataBaseSetting.MaxIdleConns)
	sqlDB.SetMaxOpenConns(dataBaseSetting.MaxOpenConns)
}

func CloseDB() {
	// 关闭gorm
	err := sqlDB.Close()
	if err != nil {
		log.Printf("数据库关闭失败: %v", err)
	}
}
