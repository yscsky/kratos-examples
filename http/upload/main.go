package main

import (
	"io"
	"log"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func uploadFile(ctx http.Context) error {
	req := ctx.Request()
	fileName := req.FormValue("name")
	file, handler, err := req.FormFile("file")
	if err != nil {
		return err
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0o666)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, file)
	return ctx.String(200, "File "+fileName+" Uploaded successfully")
}

func main() {
	httpSvr := http.NewServer(http.Address(":8000"))
	router := httpSvr.Route("/")
	router.POST("/upload", uploadFile)
	app := kratos.New(kratos.Name("upload"), kratos.Server(httpSvr))
	if err := app.Run(); err != nil {
		log.Fatal(err)
		return
	}
}
