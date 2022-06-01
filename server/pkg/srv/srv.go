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

	if result := s.H.DB.Where(&models.Worker{Id: request.Id}).First(&worker); result.Error == nil {
		return &pb.RegisterResponse{
			Status:  http.StatusConflict,
			Message: fmt.Sprintf("Worker exists with ID %v exists.", request.Id),
		}, nil
	}

	worker.Id = request.Id
	worker.Workerid = request.Workerid

	s.H.DB.Create(&worker)
	log.Info().Msg(fmt.Sprintf("Worker with ID: %v has been created and was saved to database.", request.Id))

	return &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: fmt.Sprintf("Worker with ID: %v has been registered.", request.Id),
	}, nil
}

func (s *Server) Validate(ctx context.Context, request *pb.ValidateRequest) (response *pb.ValidateResponse, err error) {
	var worker models.Worker

	claims, err := s.Jwt.ValidateToken(request.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	if result := s.H.DB.Where(&models.Worker{Id: claims.Id}).First(&worker); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "Worker not found",
		}, nil
	}

	log.Info().Str("client", worker.Workerid).Msg("Authorized.")

	return &pb.ValidateResponse{
		Status:   http.StatusOK,
		WorkerId: worker.Workerid,
	}, nil
}
