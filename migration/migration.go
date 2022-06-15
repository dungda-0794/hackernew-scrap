package migration

import (
	"hackernew-scrap/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{

		{
			ID: "202206150001",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.News{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("news")
			},
		},
	})

	if err := m.Migrate(); err != nil {
		return err
	}
	return nil
}
