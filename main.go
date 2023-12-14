package main

import (
	"flag"
	"log"

	"template-manager/grpc"
	"template-manager/rest"
)

func main() {
	var server, port string
	flag.StringVar(&server, "server", "rest", "grpc or rest")
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.Parse()

	if server == "grpc" {
		app := grpc.New()
		log.Fatal(app.Listen(port))
	} else {
		app := rest.New()
		log.Fatal(app.Listen(port))
	}
}
