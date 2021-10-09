package main

import (
	"flag"
	"log"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

func main() {
	conf := flag.String("conf", "config.yaml", "config file")
	flag.Parse()
	c := config.New(config.WithSource(file.NewSource(*conf)))
	if err := c.Load(); err != nil {
		panic(err)
	}

	var v struct {
		Service struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"service"`
	}

	if err := c.Scan(&v); err != nil {
		panic(err)
	}
	log.Printf("config: %+v", v)

	name, err := c.Value("service.name").String()
	if err != nil {
		panic(err)
	}
	log.Printf("service: %s", name)

	if err := c.Watch("service.name", func(k string, v config.Value) {
		log.Printf("config changed: %s = %v\n", k, v)
	}); err != nil {
		panic(err)
	}

	<-make(chan struct{})
}
