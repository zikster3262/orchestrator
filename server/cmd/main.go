package main

import (
	"log"

	cfg "github.com/zikster3262/orchestrator-server/pkg/config"
	srv "github.com/zikster3262/orchestrator-server/pkg/srv"
)

func main() {
	c := cfg.LoadConfig()

	err := srv.ServeGrpc(&c)
	if err != nil {
		log.Fatal(err)
	}
}
