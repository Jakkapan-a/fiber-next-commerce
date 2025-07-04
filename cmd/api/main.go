package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() { // Initialize the API server
	app := fiber.New()

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
