package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

func initRouter(app *fiber.App) {
	app.Static("/", "./public")
	app.Get("/api", getMonitors)
}

func main() {
	app := fiber.New()
	app.Use(compress.New())
	initRouter(app)
	cf := getConfig()
	_ = app.Listen(fmt.Sprintf("%v:%v", cf.Host, cf.Port))
}
