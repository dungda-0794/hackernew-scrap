package infrastructure

import (
	"hackernew-scrap/migration"
	"log"
	"os"

	"gorm.io/gorm"
)

var (
	//logger
	InfoLog *log.Logger
	ErrLog  *log.Logger

	//database
	DB *gorm.DB
)

func init() {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLog = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	var err error
	DB, err = ConnectDatabase()
	if err != nil {
		ErrLog.Fatal("fail to init db: ", err)
	}

	err = migration.Migration(DB)
	if err != nil {
		ErrLog.Fatal("fail to migration: ", err)
	}
}
