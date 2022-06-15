package scrap

import "hackernew-scrap/models"

type NewsUsecase interface {
	CreateNews(news models.News) (*models.News, error)
}
