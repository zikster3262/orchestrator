package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zikster3262/orchestrator-client/pkg/config"
	"github.com/zikster3262/orchestrator-client/pkg/srv/pb"
	"google.golang.org/grpc"
)

func main() {

	c := config.LoadConfig()

	conn, err := grpc.Dial(c.Server, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewServiceClient(conn)

	for {
		request := &pb.RegisterRequest{
			Id:       uint64(3),
			Workerid: c.Hostname,
			Token:    c.JWTSecretKey,
		}

		res, err := client.RegisterWorker(context.Background(), request)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.Status, res.Message)
		time.Sleep(time.Second * 5)
	}

}
