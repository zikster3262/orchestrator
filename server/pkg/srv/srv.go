package srv

import (
	"net"

	"github.com/zikster3262/orchestrator-server/pkg/config"
	"github.com/zikster3262/orchestrator-server/pkg/srv/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.ServiceServer
}

func ServeGrpc(c *config.Config) error {
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterServiceServer(srv, &Server{})

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
	return err
}
