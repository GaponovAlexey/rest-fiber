package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	app := fiber.New()

	app.Use(basicauth.New(basicauth.Config{
    Users: map[string]string{
        "john":  "doe",
        "admin": "1",
    },
}))

	log.Fatal(app.Listen(":3000"))
}
