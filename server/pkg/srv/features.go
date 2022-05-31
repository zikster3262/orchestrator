package srv

import (
	"context"
	"fmt"

	"github.com/zikster3262/orchestrator-server/pkg/srv/pb"
)

func (s *Server) RegisterWorker(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, err error) {
	message, id := request.GetMessage(), request.GetId()
	return &pb.RegisterResponse{
		Message: fmt.Sprintf("Reply of %s with ID: %d is OK :)", message, id),
	}, nil
}
