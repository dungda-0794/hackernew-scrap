package scrap

import "hackernew-scrap/models"

type Repository interface {
	Create(news models.News) (*models.News, error)
	Get(Id string) (*models.News, error)
}
