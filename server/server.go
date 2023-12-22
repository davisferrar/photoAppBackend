package server

import (
	"fmt"
	"log"
	"os"

	"github.com/chhoengichen/distributed-system-mission1/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func StartServer() {
	// Load environment variables from a .env file
	godotenv.Load()

	// Get the PORT value from environment variables
	portStriing := os.Getenv("PORT")
	if portStriing == "" {
		log.Fatal("Port not found in env file")
	}

	// Create a new Fiber app instance
	app := fiber.New(fiber.Config{
		// Set a body size limit of 25MB for incoming requests
		BodyLimit: 25 * 1024 * 2014,
	})

	// Set up CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,DELETE",
	}))
	// Define allowed methods for specific IPs
	// allowedMethods := map[string][]string{
	// 	"*":         {"GET"},                // Allowed methods for IP 127.0.0.1
	// 	"127.0.0.1": {"GET", "POST", "PUT"}, // Allowed methods for IP 192.168.1.1
	// }

	// Apply middleware to restrict methods by IP
	// app.Use(RestrictMethodByIP(allowedMethods))

	route.DefineRoutes(app)

	fmt.Println("Starting on port", portStriing)
	// Start the Fiber app and listen on the specified port
	log.Fatal(app.Listen(":" + portStriing))
}

func RestrictMethodByIP(allowedMethods map[string][]string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientIP := c.IP() // Get client's IP address

		// Get the allowed methods for the client's IP
		allowed, exists := allowedMethods[clientIP]
		if !exists {
			return c.Status(fiber.StatusForbidden).SendString("Access denied")
		}

		// Check if the request method is allowed for the client's IP
		methodAllowed := false
		for _, method := range allowed {
			if method == c.Method() {
				methodAllowed = true
				break
			}
		}

		if !methodAllowed {
			return c.Status(fiber.StatusMethodNotAllowed).SendString("Method not allowed for this IP")
		}

		// Continue with the next handler if the method is allowed for the IP
		return c.Next()
	}
}
