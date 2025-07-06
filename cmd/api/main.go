package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jakkapan-a/fiber-next-ecommerce/config"
	"github.com/jakkapan-a/fiber-next-ecommerce/internal/adapters/db"
)

func main() { // Initialize the API server

	app := fiber.New()

	cfg := config.LoadConfig() // Connect to the database using the configuration

	db.Connect(cfg) // Connect to the database

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Server is running",
			"status":  "success",
		})
	})

	err := app.Listen(":3000") // Start the server on port 3000
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}

	log.Println("Server is running on http://localhost:3000")
}
