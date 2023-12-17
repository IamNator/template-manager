package main

import (
	"flag"
	"log"
	"template-manager/app"
	"template-manager/config"
	"template-manager/email/mailjet"

	"template-manager/grpc"
	"template-manager/rest"
)

func main() {
	var server, port string
	flag.StringVar(&server, "server", "rest", "grpc or rest")
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.Parse()

	conf := loadConfig()

	mj := mailjet.New()

	app := app.New(mj)

	if server == "grpc" {
		grpcApp := grpc.New(conf)
		log.Fatal(grpcApp.Listen(port))
	} else {
		restApp := rest.New(conf, app)
		log.Fatal(restApp.Listen(port))
	}
}

func loadConfig() *config.Config {
	conf := config.New().
		SetEnv("", "")

	return conf
}
