package database

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"examination_system/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	databaseConfig := config.AppConfig.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		databaseConfig.User,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Port,
		databaseConfig.Name,
		databaseConfig.Charset,
		databaseConfig.ParseTime,
		databaseConfig.Loc,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  logger.Info, // 日志级别
				IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 彩色打印
			},
		),
	})
	if err != nil {
		log.Fatal("无法连接数据库:", err)
	}

	//// 自动迁移模式
	//err = DB.AutoMigrate(&model.Exampaper{}, &model.Examquestion{}, &model.User{}, &model.Teacher{}, &model.Student{})
	//if err != nil {
	//	log.Fatal("自动迁移失败:", err)
	//}
}
