package routes

import (
	"geolocation/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupUsers(app *fiber.App) {
	app.Get("/user/:id", controllers.Get)
	app.Get("/users", controllers.List)
	app.Delete("/user/:id", controllers.Delete)
	app.Put("/user/:id", controllers.Update)
	app.Post("/user", controllers.Create)
	app.Post("/answer", controllers.SendAnswer)
	app.Post("/user/score", controllers.SendScore)
	app.Get("/users/score", controllers.GetScore)
	app.Get("/azure", controllers.GetAzure)

}
