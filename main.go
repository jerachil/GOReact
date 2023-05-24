package main

import (
	"os"
	"goreact/initializers"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func init() { // init always runbefore main
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDB()
}

func main() {
	//Load Templates
	engine := html.New("./views", ".html")

	 // Setup app
	app := fiber.New(fiber.Config{
        Views: engine,
    })

	// Routes
    Routes(app)

	// Configure app
    app.Static("/", "./public")

	// Start app
    app.Listen(":" + os.Getenv("PORT"))
}
