package main

import (
	"fmt"
	"github.com/DRJ31/status/service"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func initRouter(app *fiber.App) {
	app.Get("/api", getMonitors)
	app.Use("/", static.New("./public"))
}

func main() {
	app := fiber.New()
	app.Use(compress.New())
	cf := service.GetConfig()
	initRouter(app)
	_ = app.Listen(fmt.Sprintf("%v:%v", cf.Host, cf.Port))
}
