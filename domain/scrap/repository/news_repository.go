package repository

import (
	"errors"
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/models"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

// NewRepsitory return news repository
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

func (r *repository) Get(id string) (*models.News, error) {
	var news models.News
	err := r.DB.Where("id_external = ?", id).First(&news).Error
	if err != nil {
		return nil, err
	}

	return &news, nil
}

func (r *repository) CheckExists(id string) (bool, error) {
	var news models.News
	result := r.DB.Where("id_external = ?", id).First(&news)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}

	return result.RowsAffected > 0, nil
}
