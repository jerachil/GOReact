package main

import (
	"goreact/controllers"

	"github.com/gofiber/fiber/v2"
)






func Routes(app *fiber.App) {
	frontendRoutes := []string{
		"/",
		"/about",
	}

	for _, route := range frontendRoutes {
		app.Get(route, controllers.PostsIndex)
	}
}
