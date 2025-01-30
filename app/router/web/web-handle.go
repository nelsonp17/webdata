package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nelsonp17/webdata/app/database/sqlc"
)

type Handler struct {
	Repo sqlc.Repo
}

func (h Handler) HomeView(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"LinkGooglePlay": "https://play.google.com/store/apps/details?id=com.nelsonp17.dolar",
		"LinkAppStore":   "https://apps.apple.com/us/app/dolar-argentina/id1560000000",
		"LinkDownload":   c.BaseURL() + "/assets/download",

		"Facebook":   "https://www.facebook.com/nelsonp17",
		"Instragram": "https://www.instagram.com/nelsonp17",
	})
}

func (h Handler) DownloadApp(c *fiber.Ctx) error {
	return c.SendFile("./frontend/assets/calculapp.apk")
}
