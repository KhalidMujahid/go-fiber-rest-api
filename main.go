package main

import "log"
import "github.com/gofiber/fiber/v2"
import "github.com/gofiber/template/html/v2"

func main(){

  engine := html.New("./views",".html")

  app := fiber.New(fiber.Config{
    Views: engine,
  })

  app.Static("/", "./public") 

  app.Get("/", func(ctx *fiber.Ctx) error {
    return ctx.Status(200).Render("index", fiber.Map{
      "Heading":"Hello World",
    })
  })

  app.Get("/home", func(ctx *fiber.Ctx) error {
    return ctx.Status(200).Render("dashboard", fiber.Map{})
  })

  log.Fatal(app.Listen(":3000"))
}