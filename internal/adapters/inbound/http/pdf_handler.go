package http

import (
	"pdf-tailwindcss/internal/application"

	"github.com/gofiber/fiber/v2"
)

type PDFHandler struct {
	service *application.PDFService
}

func NewPDFHandler(s *application.PDFService) *PDFHandler {
	return &PDFHandler{service: s}
}

func (h *PDFHandler) RegisterURLRoutes(app *fiber.App) {
	app.Post("/pdf", h.GeneratePDF)
}

func (h *PDFHandler) GeneratePDF(c *fiber.Ctx) error {
	var body struct {
		HTMLContent string `json:"htmlContent"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	pdf, err := h.service.GeneratePDF(body.HTMLContent)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=output.pdf")
	return c.Send(pdf)
}
