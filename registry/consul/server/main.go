package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/hashicorp/consul/api"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: fmt.Sprintf("Hello %+v", req.Name)}, nil
}

func main() {
	logger := log.NewStdLogger(os.Stdout)
	log := log.NewHelper(logger)

	cfg := api.DefaultConfig()
	cfg.Address = "192.168.101.236:8500"
	consulClient, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := &server{}

	httpSvr := http.NewServer(http.Address(":8000"), http.Middleware(recovery.Recovery(), logging.Server(logger)))
	grpcSvr := grpc.NewServer(grpc.Address(":9000"), grpc.Middleware(recovery.Recovery(), logging.Server(logger)))

	helloworld.RegisterGreeterServer(grpcSvr, s)
	helloworld.RegisterGreeterHTTPServer(httpSvr, s)

	r := consul.New(consulClient)
	app := kratos.New(kratos.Name("consul"), kratos.Server(httpSvr, grpcSvr), kratos.Registrar(r))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
