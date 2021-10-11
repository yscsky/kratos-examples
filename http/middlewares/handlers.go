package main

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

func sayHelloHandler(ctx http.Context) (err error) {
	var req helloworld.HelloRequest
	if err = ctx.BindQuery(&req); err != nil {
		return
	}
	if err = ctx.BindVars(&req); err != nil {
		return
	}
	http.SetOperation(ctx, "/helloworld.Greeter/SayHello")
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return &helloworld.HelloResponse{Message: "test:" + req.(*helloworld.HelloRequest).Name}, nil
	})
	err = ctx.Returns(h(ctx, &req))
	return
}
