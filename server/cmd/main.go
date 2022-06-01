package main

import (
	"net"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	cfg "github.com/zikster3262/orchestrator-server/pkg/config"
	"github.com/zikster3262/orchestrator-server/pkg/db"
	srv "github.com/zikster3262/orchestrator-server/pkg/srv"
	"github.com/zikster3262/orchestrator-server/pkg/srv/pb"
	"github.com/zikster3262/orchestrator-server/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	c := cfg.LoadConfig()

	h := db.Init()

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "orchestrator-grpc",
		ExpirationHours: 24 * 365,
	}

	s := srv.Server{
		H:   h,
		C:   &c,
		Jwt: jwt,
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Error().Msg("Failed to create a listener.")
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Error().Msg("Failed to serve.")
	}

	log.Info().Msg("GRPC server has been started.")

	r := gin.Default()

	r.Run(":8080")
}
