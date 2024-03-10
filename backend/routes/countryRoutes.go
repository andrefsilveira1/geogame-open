package routes

import (
	"geolocation/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupCountries(app *fiber.App) {
	app.Get("/country/:id", controllers.GetCountry)
	app.Delete("/country/:id", controllers.DeleteCountry)
	app.Get("/countries", controllers.ListCountries)
	app.Post("/country", controllers.PostCountry)
	app.Put("/country/:id", controllers.UpdateCountry)
	app.Post("/random/country", controllers.ResponseRandomCountry)
}
