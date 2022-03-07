package main

import (
	"context"
	"flag"
	"github.com/Qalifah/grey-challenge/transaction/config"
	"github.com/Qalifah/grey-challenge/transaction/database/postgres"
	"github.com/Qalifah/grey-challenge/transaction/handler"
	"github.com/Qalifah/grey-challenge/transaction/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
	"net"
	"os"
)

var configPath *string

func init() {
	configPath = flag.String("config_path", "", "path to config file")
	flag.Parse()
	if configPath == nil {
		log.Fatal("-config_path flag is required")
	}
}

func main() {
	file, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("unable to open config file: %v", err)
	}

	cfg := &config.BaseConfig{}
	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		log.Fatalf("failed to decode config file: %v", err)
	}

	ctx := context.Background()
	conn, err := postgres.New(ctx, cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to create postgres client: %v", err)
	}
	defer conn.Close(ctx)

	transferRepo := postgres.NewTransferRepository(conn)
	ctrl := handler.New(transferRepo)
	lis, err := net.Listen("tcp", cfg.ServeAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterTransactionServer(s, ctrl)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
