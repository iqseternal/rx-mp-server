package db

import (
	"demo/internal/models"
	"demo/internal/pkg/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var RdPg *gorm.DB

func init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.RdPg.Host, config.RdPg.User, config.RdPg.Password, config.RdPg.DbName, config.RdPg.Port, config.RdPg.SslMode, config.RdPg.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	//refactoringRdClient(db)

	RdPg = db
}

// refactoringRdClient 重新建立 rapid.client 表结构
func refactoringRdClient(db *gorm.DB) {
	if true {
		fmt.Println("重构表结构? 如果确信操作, 请注释 return 块")
		return
	}

	_ = db.Migrator().DropTable(&models.RdClientUser{})
	_ = db.AutoMigrate(&models.RdClientUser{})

	_ = db.Migrator().DropTable(&models.RdClientOrganization{})
	_ = db.AutoMigrate(&models.RdClientOrganization{})

	_ = db.Migrator().DropTable(&models.RdClientUserOrganization{})
	_ = db.AutoMigrate(&models.RdClientUserOrganization{})

	_ = db.Migrator().DropTable(&models.RdClientRole{})
	_ = db.AutoMigrate(&models.RdClientRole{})

	_ = db.Migrator().DropTable(&models.RdClientUserRole{})
	_ = db.AutoMigrate(&models.RdClientUserRole{})

	_ = db.Migrator().DropTable(&models.RdClientPermission{})
	_ = db.AutoMigrate(&models.RdClientPermission{})

	_ = db.Migrator().DropTable(&models.RdClientRolePermission{})
	_ = db.AutoMigrate(&models.RdClientRolePermission{})
}
