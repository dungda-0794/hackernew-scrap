package scrap

import "hackernew-scrap/models"

// NewsUsecase is interface news's usecase
type NewsUsecase interface {
	CreateNews(news models.News) (*models.News, error)
	FetchData() (bool, error)
}
