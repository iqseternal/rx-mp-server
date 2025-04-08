package storage

import (
	"fmt"
	"rx-mp/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var RdPostgres *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.RdPostgres.Host,
		config.RdPostgres.User,
		config.RdPostgres.Password,
		config.RdPostgres.DbName,
		config.RdPostgres.Port,
		config.RdPostgres.SslMode,
		config.RdPostgres.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	//refactoring(db)

	RdPostgres = db
}

// refactoringRdClient 重新建立 rapid.client 表结构
func refactoringRdClient(db *gorm.DB) {
	if true {
		fmt.Println("重构表结构? 如果确信操作, 请注释 return 块")
		return
	}
}
