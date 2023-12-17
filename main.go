package main

import (
	"flag"
	"log"
	"os"
	"template-manager/database"

	"template-manager/app"
	"template-manager/config"
	"template-manager/email/mailgun"
	"template-manager/grpc"
	"template-manager/rest"
)

func main() {
	var server, port string
	flag.StringVar(&server, "server", "rest", "grpc or rest")
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.Parse()

	conf := loadConfig()

	database.New(
		conf.GetString("POSTGRES_DSN"),
	)

	mj := mailgun.New(conf.GetString("MAILJET_DOMAIN"), conf.GetString("MAILJET_APIKEY"), conf.GetString("MAILJET_SENDER"))

	application := app.New(mj)

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
		SetEnv("MAILJET_SENDER", os.Getenv("MAILJET_SENDER")).
		SetEnv("POSTGRES_DSN", os.Getenv("POSTGRES_DSN"))

	return conf
}
