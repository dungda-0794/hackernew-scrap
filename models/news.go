package models

import (
	"hackernew-scrap/core/errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// News is struct for news model.
type News struct {
	ID         string
	IDExternal string `gorm:"index:idx_external_id,unique"`
	Title      string
	Author     string
	Link       string
	Point      int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

// BeforeCreate is hook executed berfore create record .
func (n *News) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	if err != nil {
		return errors.Wrap(err)
	}

	n.ID = id.String()

	return nil
}
