package main

import (
	"fmt"
	scheduleUsecase "hackernew-scrap/domain/schedules/usecase"
	newsDelivery "hackernew-scrap/domain/scrap/delivery/cmd"
	newsRepository "hackernew-scrap/domain/scrap/repository"
	newsUsecase "hackernew-scrap/domain/scrap/usecase"
	"hackernew-scrap/infrastructure"
	"os"

	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	infrastructure.InitGloblalVariable()

	infrastructure.InfoLog.Println("run server")
	e := echo.New()
	newsRepository := newsRepository.NewRepsitory(infrastructure.DB)
	newsUsecase := newsUsecase.NewNewsUsecase(newsRepository)
	newsSchedule := scheduleUsecase.NewScheduleUsecase(newsUsecase)
	newsDelivery.NewNewsHandler(e, newsUsecase)

	e.Start(fmt.Sprintf(":%s", port))

	newsSchedule.CronJob()
}
