package main

import (
	// "fmt"
	"uber-data-analytics/config"
	"uber-data-analytics/pkg"
	"uber-data-analytics/public"

	"github.com/gofiber/fiber/v2"
)

func main(){

	config.Initconfig()
	pkg.Init() 
	app:=fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	public.MountRoutes(app)

	app.Listen(":"+ config.Cfg.Port)
}