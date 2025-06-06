package main

import (
	"log"
	"pdf-tailwindcss/internal/adapters/inbound/http"
	"pdf-tailwindcss/internal/adapters/outbound/browserless"
	"pdf-tailwindcss/internal/application"
	"pdf-tailwindcss/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.Load()

	app := fiber.New()

	app.Use(cors.New())

	pdfGenerator := browserless.NewBrowserlessPDFGenerator(
		cfg.BrowserlessHost,
		cfg.BrowserlessToken,
		cfg.TailwindCSSFileURL,
	)
	pdfService := application.NewPDFService(pdfGenerator)
	pdfHanlder := http.NewPDFHandler(pdfService)

	pdfHanlder.RegisterURLRoutes(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}
