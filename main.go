package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"template-manager/database"
	"template-manager/email/mailjet"

	"template-manager/app"
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

	db, err := database.New(
		conf.GetString("POSTGRES_DSN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	mj := mailjet.New(
		conf.GetString("MAILJET_PUBLIC_KEY"),
		conf.GetString("MAILJET_PRIVATE_KEY"),
		conf.GetString("MAILJET_DEFAULT_SENDER"),
	)
	logger := slog.New(&slog.JSONHandler{})

	application := app.New(mj, logger, db.Client)

	if server == "grpc" {
		grpcApp := grpc.New(conf)
		log.Fatal(grpcApp.Listen(port))
	} else {
		restApp := rest.New(conf, application)
		log.Fatal(restApp.Listen(port))
	}
}

func loadConfig() *config.Config {

	conf := config.New().
		SetEnv("MAILJET_DOMAIN", os.Getenv("MAILJET_DOMAIN")).
		SetEnv("MAILJET_APIKEY", os.Getenv("MAILJET_APIKEY")).
		SetEnv("MAILJET_DEFAULT_SENDER", os.Getenv("MAILJET_DEFAULT_SENDER")).
		SetEnv("POSTGRES_DSN", os.Getenv("POSTGRES_DSN"))

	return conf
}
