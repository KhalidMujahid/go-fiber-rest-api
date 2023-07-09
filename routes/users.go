package routes

import "github.com/gofiber/fiber/v2"
import "rest-api/controllers"

func SetupRoutes(app *fiber.App) {
  app.Get("/",controllers.GetUsers)
  app.Get("/:id", controllers.GetUser)
  app.Post("/", controllers.AddUser)
  app.Delete("/:id", controllers.RemoveUser)
}
