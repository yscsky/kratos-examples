package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

func main() {
	router := gin.New()
	router.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware), gin.Logger())
	router.GET("/helloworld/:name", func(c *gin.Context) {
		name := c.Param("name")
		if name == "error" {
			kgin.Error(c, errors.Unauthorized("auth_error", "no authentication"))
			return
		}
		c.JSON(200, map[string]string{"welcome": name})
	})

	httpSvr := http.NewServer(http.Address(":8000"))
	httpSvr.HandlePrefix("/", router)

	app := kratos.New(kratos.Name("gin"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
