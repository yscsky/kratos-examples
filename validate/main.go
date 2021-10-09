package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "github.com/yscsky/kratos-examples/validate/api"
)

type server struct {
	v1.UnimplementedExampleServiceServer
}

func (s *server) TestValidate(ctx context.Context, req *v1.Request) (resp *v1.Respone, err error) {
	resp = &v1.Respone{Message: "ok"}
	return
}

func main() {
	s := &server{}

	httpSvr := http.NewServer(http.Address(":8000"), http.Middleware(validate.Validator()))
	grpcSvr := grpc.NewServer(grpc.Address(":9000"), grpc.Middleware(validate.Validator()))

	v1.RegisterExampleServiceServer(grpcSvr, s)
	v1.RegisterExampleServiceHTTPServer(httpSvr, s)

	app := kratos.New(kratos.Name("validate"), kratos.Server(httpSvr, grpcSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
{
    "id": 1,
    "age": 20,
    "code": 1,
    "score": 12.34,
    "state": true,
    "path": "/hello",
    "phone": "12345678901",
    "explain": "abc",
    "name": "name",
    "card": "1aac34",
    "info": {
        "address": "192.168.0.1"
    }
}
*/
