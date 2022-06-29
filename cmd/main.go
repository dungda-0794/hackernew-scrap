package main

import (
	"flag"
	"fmt"
	"hackernew-scrap/configs"
	scheduleUsecase "hackernew-scrap/domain/schedules/usecase"
	newsDelivery "hackernew-scrap/domain/scrap/delivery/cmd"
	"hackernew-scrap/domain/scrap/delivery/cmd/pb"
	newsRepository "hackernew-scrap/domain/scrap/repository"
	newsUsecase "hackernew-scrap/domain/scrap/usecase"
	"hackernew-scrap/infrastructure"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	port := os.Getenv("SERVER_PORT")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		infrastructure.ErrLog.Fatalf("failed to listen: %v", err)
	}
	infrastructure.InitGloblalVariable()

	infrastructure.InfoLog.Println("run server")
	newsRepository := newsRepository.NewRepsitory(infrastructure.DB)
	newsUsecase := newsUsecase.NewNewsUsecase(newsRepository)
	newsSchedule := scheduleUsecase.NewScheduleUsecase(newsUsecase)

	certFile := configs.Path("x509/cert.pem")
	keyFile := configs.Path("x509/key.pem")
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		infrastructure.ErrLog.Fatalf("Failed to generate credentials %v", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterScrapPBServer(grpcServer, newsDelivery.NewNewsHandler(newsUsecase))
	go grpcServer.Serve(lis)
	newsSchedule.CronJob()
}
