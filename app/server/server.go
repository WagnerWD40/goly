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

	router.Get("/r/:redirect", controllers.Redirect)

	router.Get("/goly", controllers.GetAllRedirects)
	router.Get("/goly/:id", controllers.GetGoly)
	router.Post("/goly", controllers.CreateGoly)
	router.Put("/goly", controllers.UpdateGoly)
	router.Delete("/goly/:id", controllers.DeleteGoly)

	router.Listen(":3000")
}
