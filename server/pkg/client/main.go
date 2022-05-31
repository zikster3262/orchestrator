package main

import (
	"context"
	"fmt"

	"github.com/zikster3262/orchestrator-server/pkg/srv/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewServiceClient(conn)
	request := &pb.RegisterRequest{
		Message: "Ping me",
		Id:      uint64(20),
	}
	ctx := context.Background()
	response, err := client.RegisterWorker(ctx, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Message)
}
