package storage

import (
	"fmt"
	"rx-mp/config"
	rdClient "rx-mp/internal/models/rd/client"
	rdMarket "rx-mp/internal/models/rd/rx_market"

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
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	//refactoring(db)

	RdPostgres = db.Debug()
}

// refactoringTable 重新建立 rapid.client 表结构
func refactoringTable(db *gorm.DB) error {
	if true {
		fmt.Println("重构表结构? 如果确信操作, 请注释 return 块")
		return nil
	}

	err := db.AutoMigrate(
		&rdClient.User{},
		&rdClient.Group{},
		&rdClient.GroupUser{},
		&rdClient.Organization{},
		&rdClient.Role{},
		&rdClient.UserRole{},
		&rdClient.Permissions{},
		&rdClient.RolePermission{},
	)

	if err != nil {
		return err
	}

	err = db.AutoMigrate(
		&rdMarket.Extension{},
		&rdMarket.ExtensionGroup{},
		&rdMarket.ExtensionVersion{},
	)

	if err != nil {
		return err
	}

	return nil
}

func refactoringView(db *gorm.DB) error {
	return db.Transaction(func(tx *gorm.DB) error {

		// 创建视图 ExtensionView
		if err := (&rdMarket.ExtensionView{}).CreateView(db); err != nil {
			return err
		}

		// 创建视图 ExtensionVersionView
		if err := (&rdMarket.ExtensionVersionView{}).CreateView(db); err != nil {
			return err
		}

		return nil
	})
}
