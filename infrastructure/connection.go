package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hackernew-scrap/configs"
)

func ConnectDatabase() (*gorm.DB, error) {
	dBConfig, err := configs.NewDBConfig()
	args := "host=" + dBConfig.Host +
		" user=" + dBConfig.User +
		" port=" + dBConfig.Port +
		" dbname=" + dBConfig.DBName +
		" password=" + dBConfig.Password +
		" sslmode=" + dBConfig.SSLMode

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	return db, err
}
