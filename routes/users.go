package routes

import "github.com/gofiber/fiber/v2"
import "rest-api/controllers"

type Person struct {
  Fname string `json:"fname"`
  Lname string `json:"lname"`
}

func SetupRoutes(app *fiber.App) {

  // render home page(index)
  app.Get("/", func(ctx *fiber.Ctx) error {
    return ctx.Status(200).Render("index", fiber.Map{})
  })

  users := app.Group("/users")
  users.Get("/",controllers.GetUsers)
  users.Get("/:id", controllers.GetUser)
  users.Post("/", controllers.AddUser)
  users.Delete("/:id", controllers.RemoveUser)
}
