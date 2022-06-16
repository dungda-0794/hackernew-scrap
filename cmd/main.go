package main

import (
	newsRepository "hackernew-scrap/domain/scrap/repository"
	newsUsecase "hackernew-scrap/domain/scrap/usecase"
	newsSchedule "hackernew-scrap/domain/scrap/schedule"
	"hackernew-scrap/infrastructure"
)

func main() {

	infrastructure.InitGloblalVariable()

	newsRepository := newsRepository.NewRepsitory(infrastructure.DB)
	newsUsecase := newsUsecase.NewNewsUsecase(newsRepository)
	newsSchedule := newsSchedule.NewNewsSchedule(newsUsecase)
	newsSchedule.CronJob()

	infrastructure.InfoLog.Println("run server")

}
