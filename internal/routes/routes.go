package routes

import (
	"labelin/internal/config"
	"labelin/internal/models"
	"labelin/internal/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg config.Config) {
	// Serve static files
	app.Static("/", "./public")

	// HTMX Views / Frontend (Public/No Auth for simplicity in this demo, but should be protected)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/guests")
	})

	app.Get("/guests", func(c *fiber.Ctx) error {
		var guests []models.Guest
		config.DB.Find(&guests)
		return c.Render("pages/guests", fiber.Map{
			"Guests": guests,
			"Title": "Master Data Tamu",
		}, "layouts/main")
	})

	app.Get("/templates", func(c *fiber.Ctx) error {
		var templates []models.LabelTemplate
		config.DB.Find(&templates)
		return c.Render("pages/templates", fiber.Map{
			"Templates": templates,
			"Title": "Label Templates",
		}, "layouts/main")
	})

	app.Get("/settings", func(c *fiber.Ctx) error {
		var setting models.Setting
		config.DB.First(&setting)
		return c.Render("pages/settings", fiber.Map{
			"Setting": setting,
			"Title": "Pengaturan",
		}, "layouts/main")
	})

	// API Group
	api := app.Group("/api")

	// Guests API
	guests := api.Group("/guests")
	guests.Get("/", GetGuests)
	guests.Post("/", CreateGuest)
	guests.Delete("/:id", DeleteGuest)
	guests.Post("/import", ImportGuests)
	guests.Post("/print", PrintLabels)
	// Settings API
	api.Post("/settings", UpdateSettings)

	// Templates API
	templates := api.Group("/templates")
	templates.Post("/", CreateTemplate)
	templates.Delete("/:id", DeleteTemplate)
}

// CreateTemplate
func CreateTemplate(c *fiber.Ctx) error {
	tpl := new(models.LabelTemplate)
	if err := c.BodyParser(tpl); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&tpl)
	
	if string(c.Request().Header.Peek("Hx-Request")) == "true" {
		c.Response().Header.Set("HX-Refresh", "true")
		return c.SendStatus(200)
	}
	return c.JSON(tpl)
}

func DeleteTemplate(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.LabelTemplate{}, id)
	return c.SendStatus(200)
}

// UpdateSettings
func UpdateSettings(c *fiber.Ctx) error {
	var setting models.Setting
	if err := c.BodyParser(&setting); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	var current models.Setting
	config.DB.First(&current)
	setting.ID = current.ID // Update existing
	
	config.DB.Save(&setting)
	return c.SendStatus(200)
}

// GetGuests
func GetGuests(c *fiber.Ctx) error {
	var guests []models.Guest
	config.DB.Find(&guests)
	return c.JSON(guests)
}

// CreateGuest
func CreateGuest(c *fiber.Ctx) error {
	guest := new(models.Guest)
	if err := c.BodyParser(guest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	config.DB.Create(&guest)

	// Return HTMX fragment or JSON based on accept header
	if string(c.Request().Header.Peek("Hx-Request")) == "true" {
		c.Response().Header.Set("HX-Refresh", "true")
		return c.SendStatus(200)
	}

	return c.JSON(guest)
}

func DeleteGuest(c *fiber.Ctx) error {
	id := c.Params("id")
	config.DB.Delete(&models.Guest{}, id)
	return c.SendStatus(200)
}

func ImportGuests(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "File upload required"})
	}

	excelSvc := services.NewExcelService(config.DB)
	result, err := excelSvc.ImportGuests(file)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(result)
}

func PrintLabels(c *fiber.Ctx) error {
	// Parse request for Guest IDs and Template ID
	type PrintReq struct {
		GuestIDs   []uint `json:"guest_ids"`
		TemplateID uint   `json:"template_id"`
	}
	var req PrintReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var tpl models.LabelTemplate
	if err := config.DB.First(&tpl, req.TemplateID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Template not found"})
	}

	var guests []models.Guest
	if len(req.GuestIDs) > 0 {
		config.DB.Where("id IN ?", req.GuestIDs).Find(&guests)
	} else {
		// print all
		config.DB.Find(&guests)
	}

	if len(guests) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No guests found"})
	}

	pdf, err := services.GeneratePDFLabels(guests, tpl)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate PDF"})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=labels.pdf")
	
	if err := pdf.Output(c.Response().BodyWriter()); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to write PDF"})
	}
	return nil
}
