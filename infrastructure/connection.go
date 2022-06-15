package infrastructure

import (
	"hackernew-scrap/configs"
	"hackernew-scrap/core/errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDatabase is function to connect postgresql
func ConnectDatabase() (*gorm.DB, error) {
	dBConfig, err := configs.NewDBConfig()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	args := "host=" + dBConfig.Host +
		" user=" + dBConfig.User +
		" port=" + dBConfig.Port +
		" dbname=" + dBConfig.DBName +
		" password=" + dBConfig.Password +
		" sslmode=" + dBConfig.SSLMode

	db, err := gorm.Open(postgres.Open(args), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return db, errors.Wrap(err)
}
