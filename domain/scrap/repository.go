package scrap

import "hackernew-scrap/models"

// Repository for news repository interface.
type Repository interface {
	Create(news models.News) (*models.News, error)
	Get(id string) (*models.News, error)
	CheckExists(id string) (bool, error)
}
