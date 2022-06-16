package usecase

import (
	"hackernew-scrap/core/errors"
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/models"
)

type newsUsecase struct {
	repository scrap.Repository
}

// NewNewsUsecase for create news usecase.
func NewNewsUsecase(repository scrap.Repository) scrap.NewsUsecase {
	return &newsUsecase{repository: repository}
}

func (n *newsUsecase) CreateNews(news models.News) (*models.News, error) {
	result, err := n.repository.Create(news)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}
