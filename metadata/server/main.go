package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/metadata"
	mmd "github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (resp *helloworld.HelloResponse, err error) {
	var extra string
	if md, ok := metadata.FromServerContext(ctx); ok {
		extra = md.Get("x-md-global-extra")
	}
	resp = &helloworld.HelloResponse{Message: fmt.Sprintf("Hello %s extra_meta: %s", req.Name, extra)}
	return
}

func main() {
	s := &server{}

	httpSvr := http.NewServer(http.Address(":8000"), http.Middleware(mmd.Server()))
	grpcSvr := grpc.NewServer(grpc.Address(":9000"), grpc.Middleware(mmd.Server()))

	helloworld.RegisterGreeterServer(grpcSvr, s)
	helloworld.RegisterGreeterHTTPServer(httpSvr, s)

	app := kratos.New(kratos.Name("metadata"), kratos.Server(httpSvr, grpcSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
