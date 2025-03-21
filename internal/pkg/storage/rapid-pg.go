package storage

import (
	"fmt"
	"rx-mp/config"
	rd_client "rx-mp/internal/models/rd/client"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var RdPostgress *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.RdPostgress.Host,
		config.RdPostgress.User,
		config.RdPostgress.Password,
		config.RdPostgress.DbName,
		config.RdPostgress.Port,
		config.RdPostgress.SslMode,
		config.RdPostgress.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//refactoring(db)

	RdPostgress = db
}

// refactoringRdClient 重新建立 rapid.client 表结构
func refactoringRdClient(db *gorm.DB) {
	if true {
		fmt.Println("重构表结构? 如果确信操作, 请注释 return 块")
		return
	}

	_ = db.Migrator().DropTable(&rd_client.User{})
	_ = db.AutoMigrate(&rd_client.User{})

	_ = db.Migrator().DropTable(&rd_client.Organization{})
	_ = db.AutoMigrate(&rd_client.Organization{})

	_ = db.Migrator().DropTable(&rd_client.UserOrganization{})
	_ = db.AutoMigrate(&rd_client.UserOrganization{})

	_ = db.Migrator().DropTable(&rd_client.Role{})
	_ = db.AutoMigrate(&rd_client.Role{})

	_ = db.Migrator().DropTable(&rd_client.UserRole{})
	_ = db.AutoMigrate(&rd_client.UserRole{})

	_ = db.Migrator().DropTable(&rd_client.UserPermission{})
	_ = db.AutoMigrate(&rd_client.UserPermission{})

	_ = db.Migrator().DropTable(&rd_client.RolePermission{})
	_ = db.AutoMigrate(&rd_client.RolePermission{})
}
