package main

import (
	"labelin/internal/config"
	"labelin/internal/models"
	"labelin/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to Database
	config.ConnectDB(cfg)

	// Auto Migrate DB
	err := config.DB.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Category{},
		&models.Guest{},
		&models.LabelTemplate{},
		&models.PrintLog{},
		&models.Setting{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate DB: %v", err)
	}

	// Initialize standard HTML template engine
	engine := html.New("./views", ".html")

	// Create Fiber App
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main", // default layout
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Setup Routes
	routes.SetupRoutes(app, cfg)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
