package routes

import (
	"ketdevbanban/controllers"

	middleware "ketdevbanban/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//Public Routes
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	// Private Routes
	app.Use(middleware.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Put("/api/users/info", controllers.UpdateInfo)
	app.Put("/api/users/password", controllers.UpdatePassword)

	//User
	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

}
