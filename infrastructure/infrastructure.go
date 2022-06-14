package infrastructure

import (
	"gorm.io/gorm"
	"log"
	"os"
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
}
