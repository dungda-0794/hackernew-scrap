package repository

import (
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/models"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewRepsitory(db *gorm.DB) scrap.Repository {
	return &repository{DB: db}
}

func (r *repository) Create(news models.News) (*models.News, error) {
	err := r.DB.Create(&news).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}
func (r *repository) Get(Id string) (*models.News, error) {
	var news models.News
	err := r.DB.Where("id_external = ?", Id).First(&news).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}
