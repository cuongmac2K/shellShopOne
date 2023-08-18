package shell

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"os/exec"
)

func CreateShell(c *fiber.Ctx) (err error){
	_ = exec.CommandContext(
		context.Background(),
		"docker",
		"run",
		"--name",
		"hello-world",
		//"-it",
		"alpine:latest")
	return
}
