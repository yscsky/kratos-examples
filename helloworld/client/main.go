package main

import (
	"context"
	"log"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/yscsky/kratos-examples/helloworld/helloworld"
)

func main() {
	callHTTP()
	callGRPC()
}

func callHTTP() {
	ctx := context.Background()
	conn, err := http.NewClient(ctx, http.WithMiddleware(recovery.Recovery()), http.WithEndpoint(":8000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterHTTPClient(conn)
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "God"})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("[http] SayHello %s\n", resp.Message)

	if _, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "error"}); err != nil {
		log.Printf("[http] SayHello error: %v\n", err)
	}
	if errors.IsBadRequest(err) {
		log.Printf("[http] SayHello error is invalid argument: %v\n", err)
	}
}

func callGRPC() {
	ctx := context.Background()
	conn, err := grpc.DialInsecure(ctx, grpc.WithMiddleware(recovery.Recovery()), grpc.WithEndpoint(":9000"))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := helloworld.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: "God"})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("[http] SayHello %s\n", resp.Message)

	if _, err = client.SayHello(ctx, &helloworld.HelloRequest{Name: "error"}); err != nil {
		log.Printf("[http] SayHello error: %v\n", err)
	}
	if errors.IsBadRequest(err) {
		log.Printf("[http] SayHello error is invalid argument: %v\n", err)
	}
}
