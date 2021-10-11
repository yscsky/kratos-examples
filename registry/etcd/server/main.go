package main

import (
	"context"
	"fmt"
	"log"

	etcd "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
	etcdcli "go.etcd.io/etcd/client/v3"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	return &helloworld.HelloResponse{Message: fmt.Sprintf("Hello %+v", req.Name)}, nil
}

func main() {
	client, err := etcdcli.New(etcdcli.Config{
		Endpoints: []string{"192.168.101.236:2379"},
	})
	if err != nil {
		log.Fatal(err)
	}

	s := &server{}

	httpSvr := http.NewServer(http.Address(":8000"), http.Middleware(recovery.Recovery()))
	grpcSvr := grpc.NewServer(grpc.Address(":9000"), grpc.Middleware(recovery.Recovery()))

	helloworld.RegisterGreeterServer(grpcSvr, s)
	helloworld.RegisterGreeterHTTPServer(httpSvr, s)

	r := etcd.New(client)
	app := kratos.New(kratos.Name("etcd"), kratos.Server(httpSvr, grpcSvr), kratos.Registrar(r))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
