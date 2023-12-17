package main

import (
	"flag"
	"log"
	"template-manager/config"

	"template-manager/grpc"
	"template-manager/rest"
)

func main() {
	var server, port string
	flag.StringVar(&server, "server", "rest", "grpc or rest")
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.Parse()

	conf := loadConfig()

	if server == "grpc" {
		app := grpc.New(conf)
		log.Fatal(app.Listen(port))
	} else {
		app := rest.New(conf)
		log.Fatal(app.Listen(port))
	}
}

func loadConfig() *config.Config {
	conf := config.New().
		SetEnv("", "")

	return conf
}
