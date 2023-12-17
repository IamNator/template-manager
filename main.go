package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"template-manager/app/session"
	"template-manager/database"
	"template-manager/email/mailjet"
	"template-manager/entity"
	"template-manager/rest/middleware"

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
	err = db.Client.AutoMigrate(&entity.Account{}, entity.Key{}, entity.Session{})
	if err != nil {
		log.Fatal(err)
	}
	mj := mailjet.New(
		conf.GetString("MAILJET_PUBLIC_KEY"),
		conf.GetString("MAILJET_PRIVATE_KEY"),
		conf.GetString("MAILJET_DEFAULT_SENDER"),
		mailjet.WithName("template manager"),
	)
	logger := slog.New(&slog.JSONHandler{})
	sessionManager := session.New(db.Client, conf, logger)
	midware := middleware.NewAuth(sessionManager)
	
	application := app.New(conf, mj, logger, db.Client, sessionManager)

	if server == "grpc" {
		grpcApp := grpc.New(conf)
		log.Fatal(grpcApp.Listen(port))
	} else {
		restApp := rest.New(conf, application, midware)
		log.Fatal(restApp.Listen(port))
	}
}

func loadConfig() *config.Config {

	conf := config.New().
		SetEnv("MAILJET_DOMAIN", os.Getenv("MAILJET_DOMAIN")).
		SetEnv("MAILJET_PRIVATE_KEY", os.Getenv("MAILJET_PRIVATE_KEY")).
		SetEnv("MAILJET_PUBLIC_KEY", os.Getenv("MAILJET_PUBLIC_KEY")).
		SetEnv("MAILJET_DEFAULT_SENDER", os.Getenv("MAILJET_DEFAULT_SENDER")).
		SetEnv("POSTGRES_DSN", os.Getenv("POSTGRES_DSN")).
		SetEnv("JWT_SIGNING_KEY", os.Getenv("JWT_SIGNING_KEY"))

	return conf
}
