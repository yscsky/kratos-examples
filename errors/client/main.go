package main

import (
	"context"
	"flag"
	"log"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/errors/api"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

func main() {
	name := flag.String("name", "", "name")
	flag.Parse()

	callHTTP(*name)
	callGRPC(*name)
}

func callHTTP(name string) {
	ctx := context.Background()
	conn, err := http.NewClient(ctx, http.WithMiddleware(recovery.Recovery()), http.WithEndpoint(":8000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterHTTPClient(conn)
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		if errors.Code(err) == 500 {
			log.Println(err)
		}
		if api.IsUserNotFound(err) {
			log.Println("[http] USER_NOT_FOUND_ERROR", err)
		}
		return
	}
	log.Printf("[http] SayHello %s\n", resp.Message)
}

func callGRPC(name string) {
	ctx := context.Background()
	conn, err := grpc.DialInsecure(ctx, grpc.WithMiddleware(recovery.Recovery()), grpc.WithEndpoint(":9000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		e := errors.FromError(err)
		if e.Message == "USER_NAME_EMPTY" {
			log.Println("[grpc] USER_NAME_EMPTY", err)
		}
		if api.IsUserNotFound(err) {
			log.Println("[grpc] USER_NOT_FOUND_ERROR", err)
		}
		return
	}
	log.Printf("[grpc] SayHello %s\n", resp.Message)
}
