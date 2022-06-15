package usecase

import (
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/models"
)

type newsUsecase struct {
	repository scrap.Repository
}

func NewNewsUsecase(repository scrap.Repository) scrap.NewsUsecase {
	return &newsUsecase{repository: repository}
}

func (n *newsUsecase) CreateNews(news models.News) (*models.News, error) {
	result, err := n.repository.Create(news)
	if err != nil {
		return nil, err
	}

	return result, nil
}
