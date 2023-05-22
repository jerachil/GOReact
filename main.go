package main

import (
	"os"
	"goreact/initializers"
	"goreact/controllers"
	"github.com/gofiber/fiber/v2"
)

func init() { // init always runbefore main
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	 // Setup app
	app := fiber.New()

	// Routes
    app.Get("/", controllers.PostsIndex)
    

	// Start app
    app.Listen(":" + os.Getenv("PORT"))
}
