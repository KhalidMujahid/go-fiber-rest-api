package main

import "log"
import "github.com/gofiber/template/html/v2"
import "github.com/gofiber/fiber/v2"
import "rest-api/routes"


func main(){

  engine := html.New("./views",".html")

  app := fiber.New(fiber.Config{
    Views: engine,
  })

  app.Static("/", "./public")

  routes.SetupRoutes(app)

  log.Fatal(app.Listen(":3001"))
}