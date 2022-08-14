package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stheven26/controllers"
)

func Routes(app *fiber.App) {
	app.Post("/api/v1/register", controllers.Register)
	app.Post("/api/v1/login", controllers.Login)
	app.Post("/api/v1/logout", controllers.Logout)
	app.Get("/api/v1/data", controllers.GetData)
}
