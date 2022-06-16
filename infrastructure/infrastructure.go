package infrastructure

import (
	"hackernew-scrap/migration"
	"log"
	"os"

	"gorm.io/gorm"
)

var (
	// InfoLog are global variable for log stdout
	InfoLog *log.Logger
	// ErrLog are global variable for log stdout
	ErrLog *log.Logger

	// DB are global variable for connection database
	DB *gorm.DB
)

// InitGloblalVariable is function to init varable to run server
func InitGloblalVariable() {
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
