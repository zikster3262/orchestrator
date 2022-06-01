package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/zikster3262/orchestrator-client/pkg/config"
	"github.com/zikster3262/orchestrator-client/pkg/models"
	"github.com/zikster3262/orchestrator-client/pkg/srv/pb"
	"github.com/zikster3262/orchestrator-client/pkg/utils"
	"google.golang.org/grpc"
)

func main() {

	c := config.LoadConfig()

	conn, err := grpc.Dial(c.Server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewServiceClient(conn)

	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "orchestrator-grpc",
		ExpirationHours: 24 * 365,
	}
	id := uuid.New().String()

	w := models.Worker{
		Id:       id,
		Workerid: c.Hostname,
	}

	request := &pb.RegisterRequest{
		Id:       w.Id,
		Workerid: w.Workerid,
	}

	token, _ := jwt.GenerateToken(w)
	w.Token = token

	valreg := &pb.ValidateRequest{
		Token: w.Token,
		Id:    w.Id,
	}

	_, err = client.RegisterWorker(context.Background(), request)
	if err != nil {
		log.Err(err).Str("client", c.Hostname).Msg(fmt.Sprintf("Error: %v\n", err))
	} else {
		log.Info().Str("client", c.Hostname).Msg("Registered.")
	}

	for {

		_, err = client.Validate(context.Background(), valreg)
		if err != nil {
			log.Err(err).Str("client", c.Hostname).Msg(fmt.Sprintf("Validate Request was not sent. Error: %v\n", err))
		}

		time.Sleep(time.Second * 5)
	}

}
