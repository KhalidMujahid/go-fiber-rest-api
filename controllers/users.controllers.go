package controllers

import "github.com/gofiber/fiber/v2"
import "rest-api/models"

// Get all users
func GetUsers(ctx *fiber.Ctx) error {
	return ctx.JSON(models.GetUsers())
}

// Get single user
func GetUser(ctx *fiber.Ctx) error {
	param := ctx.Params("id")

	user := models.GetUser(param)

	if user.ID != "" {
		return ctx.Status(200).JSON(models.GetUser(param))
	} else {
		return ctx.Status(404).JSON(fiber.Map{ "message":"User not found!" })
	}

}

// Add user
func AddUser(ctx *fiber.Ctx) error {
	var body models.BodyData

	err := ctx.BodyParser(&body)

	if err != nil {
      return ctx.Status(422).JSON(fiber.Map{ "message": "All fields are required"})
    }

	// send it to the models method

	bodyData := models.AddUser(body)

	if bodyData.ID != "" {
		return ctx.Status(201).JSON(bodyData)
	} else {
		return ctx.Status(400).JSON(fiber.Map{ "message":"Error occured while adding new" })
	}
}

// Update user
func RemoveUser(ctx *fiber.Ctx) error {
	paramId := ctx.Params("id")

	return ctx.Status(200).JSON(fiber.Map{ "message": models.RemoveUser(paramId) })
}