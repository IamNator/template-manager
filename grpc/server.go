package grpc

import (
	"context"
	"log"
	"time"
)

type server struct {
	UnimplementedServiceServer
}

var _ ServiceServer = (*server)(nil)

func (s *server) Download(ctx context.Context, req *TemplateRequest) (*Template, error) {

	// fetch template url from db
	templateName := "Template Name"  // Replace with actual template name
	version := "v1.0"                // Replace with actual version
	createdAt := time.Now().String() // Replace with actual creation time

	// download template from url

	//if vars prefill template

	// convert to bytes
	templateContent := []byte("Your template content goes here")

	// Optionally use the version and vars from the request
	log.Printf("Requested Template ID: %v, Version: %v, Vars: %v", req.TemplateId, req.Version, req.Vars)

	// send template to client
	return &Template{
		Content:   templateContent,
		Name:      templateName,
		Version:   version,
		CreatedAt: createdAt,
	}, nil
}

func New() *server {
	return &server{}
}

// Run starts the gRPC server
func (s server) Listen(port string) error {
	log.Println("Starting gRPC server on port ", port)
	ch := make(chan error)
	<-ch
	return nil
}
