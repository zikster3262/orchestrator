package srv

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zikster3262/orchestrator-server/pkg/config"
	"github.com/zikster3262/orchestrator-server/pkg/db"
	"github.com/zikster3262/orchestrator-server/pkg/models"
	"github.com/zikster3262/orchestrator-server/pkg/srv/pb"
	"github.com/zikster3262/orchestrator-server/pkg/utils"
)

type Server struct {
	C *config.Config
	pb.ServiceServer
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) RegisterWorker(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error) {
	var worker models.Worker

	if result := s.H.DB.Where(&models.Worker{Id: int64(request.Id)}).First(&worker); result.Error == nil {
		return &pb.RegisterResponse{
			Status:  http.StatusConflict,
			Message: fmt.Sprintf("Worker exists with ID %v exists.", request.Id),
		}, nil
	}

	worker.Id = int64(request.Id)
	worker.Token = utils.HashPassword(request.Token)

	s.H.DB.Create(&worker)
	log.Info().Msg(fmt.Sprintf("Worker with ID: %d has been created and was saved to database.", request.Id))

	return &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: fmt.Sprintf("Worker with ID: %d has been registered.", request.Id),
	}, nil
}

func (s *Server) ValidateAuth(ctx context.Context, request *pb.AuthRequest) (response *pb.AuthResponse, err error) {
	token := request.GetToken()
	return &pb.AuthResponse{
		Message: token,
	}, nil
}
