package server

import (
	"app/server/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/goly", controllers.GetAllRedirects)
	router.Get("/goly/:id", controllers.GetGoly)
	router.Post("/goly", controllers.CreateGoly)

	router.Listen(":3000")
}
