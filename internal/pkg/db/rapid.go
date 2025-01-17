package db

import (
	"demo/internal/pkg/config"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgRapid *gorm.DB

func Init(config *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.PG.Host, config.PG.User, config.PG.Password, config.PG.DbName, config.PG.Port, config.PG.SslMode, config.PG.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err == nil {
		PgRapid = db
	}
}
