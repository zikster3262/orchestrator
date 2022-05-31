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
		Id:    uint64(3),
		Token: "dasdasdasdas12easdas",
	}
	ctx := context.Background()
	response, err := client.RegisterWorker(ctx, request)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Message)
}
