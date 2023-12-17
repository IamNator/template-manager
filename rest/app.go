package rest

import (
	"github.com/gofiber/fiber/v2"

	"template-manager/app"
	"template-manager/config"
)

type server struct {
	App *app.App
}

// New creates a new fiber app
func New(conf *config.Config, app *app.App) *server {
	return &server{
		App: app,
	}
}

func (s server) Listen(port string) error {
	app := fiber.New()
	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Setup route for the API health check
	app.Get("/api/health", health)
	app.Get("/api/stats", stats)

	// Define API endpoints for managing keys
	app.Post("/api/key", addKey)
	app.Get("/api/key", getKey)
	app.Delete("/api/key/:id", deleteKey)

	// Define API endpoints for managing users
	app.Post("/api/user/login", s.Login)
	app.Post("/api/user/signup", s.Signup)
	app.Post("/api/user/logout", logout)

	// Define API endpoints for managing templates
	app.Post("/api/template", addTemplate)
	app.Get("/api/template", getTemplate)
	app.Delete("/api/template/:id", deleteTemplate)
	app.Put("/api/template/:id", updateTemplate)

	// Start the server on port 8080
	return app.Listen(port)
}

func health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "pong",
	})
}

func stats(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"grpc":    true,
		"version": "v1.0.0",
		"open":    false, // open source version
	})
}
