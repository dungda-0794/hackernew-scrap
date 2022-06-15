package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type News struct {
	Id         string
	IdExternal string `gorm:"index:idx_external_id,unique"`
	Title      string
	Author     string
	Link       string
	Point      int
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}

func (n *News) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	n.Id = id.String()

	return
}
