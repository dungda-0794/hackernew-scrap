package migration

import (
	"hackernew-scrap/core/errors"
	"hackernew-scrap/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// Migration database.
func Migration(db *gorm.DB) error {
	if err := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202206150001",
			Migrate: func(tx *gorm.DB) error {
				return errors.Wrap(tx.AutoMigrate(&models.News{}))
			},
			Rollback: func(tx *gorm.DB) error {
				return errors.Wrap(tx.Migrator().DropTable("news"))
			},
		},
	}).Migrate(); err != nil {
		return errors.Wrap(err)
	}

	return nil
}
