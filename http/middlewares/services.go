package main

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (resp *helloworld.HelloResponse, err error) {
	if req.Name == "error" {
		err = errors.BadRequest("custom_error", fmt.Sprintf("invalid argument %s", req.Name))
		return
	}
	if req.Name == "panic" {
		panic("server panic")
	}
	resp.Message = fmt.Sprintf("Hello %+v", req.Name)
	return
}
