package delivery

import (
	"context"
	"hackernew-scrap/domain/scrap"
	"hackernew-scrap/domain/scrap/delivery/cmd/pb"
)

// NewsHandler represent the httphandler for article
type NewsHandler struct {
	AUsecase scrap.NewsUsecase
	pb.UnimplementedScrapPBServer
}

func NewNewsHandler(us scrap.NewsUsecase) *NewsHandler {
	s := &NewsHandler{
		AUsecase: us,
	}
	return s
}

// SendNotify will send article to slack
func (a *NewsHandler) SendNotify(ctx context.Context, in *pb.SendNotifyRequest) (*pb.SendNotifyResponse, error) {
	res, err := a.AUsecase.FetchData()
	if err != nil {
		return &pb.SendNotifyResponse{Success: false}, err
	}
	return &pb.SendNotifyResponse{Success: res}, nil
}
