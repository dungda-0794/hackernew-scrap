package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Connection string `envconfig:"DB_CONNECTION" default:"postgres"`
	Host       string `envconfig:"DB_HOST" default:"postgres"`
	Port       string `envconfig:"DB_PORT" default:"5432"`
	User       string `envconfig:"DB_USER" default:"postgres"`
	Password   string `envconfig:"DB_PASSWORD" default:"postgres"`
	DBName     string `envconfig:"DB_DBNAME" default:"postgres"`
	SSLMode    string `envconfig:"DB_SSLMODE" default:"disable"`
}

func NewDBConfig() (*DBConfig, error) {
	var dbconfig DBConfig

	err := envconfig.Process("", &dbconfig)

	return &dbconfig, err
}
