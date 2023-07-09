package main

import "log"
import "github.com/gofiber/fiber/v2"
import "rest-api/routes"


func main(){

  app := fiber.New()

  routes.SetupRoutes(app)

  log.Fatal(app.Listen(":3000"))
}