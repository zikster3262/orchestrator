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
	for {
		client := pb.NewServiceClient(conn)
		request := &pb.RegisterRequest{
			Id:    uint64(3),
			Token: "dasdasdasdas12easdas",
		}
		ctx := context.Background()
		response, err := client.RegisterWorker(ctx, request)
		if err != nil {
			panic(err)
		}
		fmt.Println(response.Message)
		fmt.Println(response.Status)
		time.Sleep(time.Second * 5)
	}

}
