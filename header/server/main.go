package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (resp *helloworld.HelloResponse, err error) {
	resp = &helloworld.HelloResponse{}
	info, _ := kratos.FromContext(ctx)
	if tr, ok := transport.FromServerContext(ctx); ok {
		tr.ReplyHeader().Set("app_name", info.Name())
	}
	resp.Message = fmt.Sprintf("Hello %+v", req.Name)
	return
}

func main() {
	s := &server{}

	httpSvr := http.NewServer(http.Address(":8000"), http.Middleware(recovery.Recovery()))
	grpcSvr := grpc.NewServer(grpc.Address(":9000"), grpc.Middleware(recovery.Recovery()))

	helloworld.RegisterGreeterServer(grpcSvr, s)
	helloworld.RegisterGreeterHTTPServer(httpSvr, s)

	app := kratos.New(kratos.Name("header"), kratos.Server(httpSvr, grpcSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
