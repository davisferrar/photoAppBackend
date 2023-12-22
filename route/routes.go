package route

import (
	"github.com/chhoengichen/distributed-system-mission1/handler"
	"github.com/gofiber/fiber/v2"
)

func DefineRoutes(app *fiber.App) {
	/* -------------------------------------------------------------------------- */
	/*                                Define routes                               */
	/* -------------------------------------------------------------------------- */
	app.Get("/display", handler.GetImage)
	app.Post("/upload", handler.UploadImage)
	app.Delete("/delete", handler.DeleteImage)
	app.Static("/client/public", "./client/public")
}
