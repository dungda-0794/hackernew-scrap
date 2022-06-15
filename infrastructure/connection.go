package infrastructure

import (
	"hackernew-scrap/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	dBConfig, err := configs.NewDBConfig()
	if err != nil {
		return nil, err
	}

	args := "host=" + dBConfig.Host +
		" user=" + dBConfig.User +
		" port=" + dBConfig.Port +
		" dbname=" + dBConfig.DBName +
		" password=" + dBConfig.Password +
		" sslmode=" + dBConfig.SSLMode

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
