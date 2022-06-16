package configs

import (
	"hackernew-scrap/core/errors"

	"github.com/kelseyhightower/envconfig"
)

// DBConfig struct for config database.
type DBConfig struct {
	Connection string `envconfig:"DB_CONNECTION" default:"postgres"`
	Host       string `envconfig:"DB_HOST" default:"postgres"`
	Port       string `envconfig:"DB_PORT" default:"5432"`
	User       string `envconfig:"DB_USER" default:"postgres"`
	Password   string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName     string `envconfig:"DB_DBNAME" default:"postgres"`
	SSLMode    string `envconfig:"DB_SSLMODE" default:"disable"`
}

// NewDBConfig return config for database.
func NewDBConfig() (*DBConfig, error) {
	var dbconfig DBConfig

	err := envconfig.Process("", &dbconfig)

	return &dbconfig, errors.Wrap(err)
}
