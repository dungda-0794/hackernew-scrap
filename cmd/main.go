package main

import (
	newsRepository "hackernew-scrap/domain/scrap/repository"
	newsUsecase "hackernew-scrap/domain/scrap/usecase"
	"hackernew-scrap/infrastructure"
)

func main() {

	newsRepository := newsRepository.NewRepsitory(infrastructure.DB)
	_ = newsUsecase.NewNewsUsecase(newsRepository)

	infrastructure.InfoLog.Println("run server")

}
