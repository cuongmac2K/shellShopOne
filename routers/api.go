package routers

import (
	"github.com/gofiber/fiber/v2"
	"shellshoponeweb/controller/shell"
)

func ShellApi(app *fiber.App) {
	group:= app.Group("/v0.1")
	shellG:= group.Group("/shell")
	shellG.Post("/create",shell.CreateShell)
}
