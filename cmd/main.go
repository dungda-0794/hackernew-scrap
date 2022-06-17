package main

import (
	newsRepository "hackernew-scrap/domain/scrap/repository"
	newsSchedule "hackernew-scrap/domain/scrap/schedule"
	newsUsecase "hackernew-scrap/domain/scrap/usecase"
	"hackernew-scrap/infrastructure"
)

func main() {

	infrastructure.InitGloblalVariable()

	infrastructure.InfoLog.Println("run server")
	newsRepository := newsRepository.NewRepsitory(infrastructure.DB)
	newsUsecase := newsUsecase.NewNewsUsecase(newsRepository)
	newsSchedule := newsSchedule.NewNewsSchedule(newsUsecase)
	newsSchedule.CronJob()

}
