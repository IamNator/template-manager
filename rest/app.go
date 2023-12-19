package rest

import (
	"github.com/gofiber/fiber/v2"
	"template-manager/app"
	"template-manager/config"
)

type Middleware interface {
	FiberAuthMiddleware(c *fiber.Ctx) error
}
type server struct {
	conf       *config.Config
	app        *app.App
	middleware Middleware
}

// New creates a new fiber app
func New(conf *config.Config, app *app.App, middleware Middleware) *server {
	return &server{
		conf:       conf,
		app:        app,
		middleware: middleware,
	}
}

func (s server) Listen(port string) error {
	app := fiber.New()

	app.Use(
		s.middleware.FiberAuthMiddleware,
	)

	// Serve static files from the "public" directory
	app.Static("/", "./public")

	// Setup route for the API health check
	app.Get("/api/health", health)
	app.Get("/api/stats", stats)

	// Define API endpoints for managing users
	app.Post("/api/user/login", s.Login)
	app.Post("/api/user/signup", s.Signup)
	app.Post("/api/user/logout", s.Logout)

	// Define API endpoints for managing keys
	app.Post("/api/key", s.AddKey)
	app.Get("/api/key", s.ListAccessKeys)
	app.Delete("/api/key/:id", s.DeleteKey)

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
