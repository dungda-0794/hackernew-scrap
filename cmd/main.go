package main

import (
	newsRepository "hackernew-scrap/domain/scrap/repository"
	newsUsecase "hackernew-scrap/domain/scrap/usecase"
	"hackernew-scrap/infrastructure"
)

func main() {

	infrastructure.InitGloblalVariable()

	newsRepository := newsRepository.NewRepsitory(infrastructure.DB)
	_ = newsUsecase.NewNewsUsecase(newsRepository)
	newsUsecase.NewCronJob()

	infrastructure.InfoLog.Println("run server")

}
