package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-di-fiber/domain/dao"
	"go-di-fiber/service"
)

func heathCheck(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Hello, world")
	return c.SendString(msg)
}

func signup(userService *service.UserService, validator *validator.Validate) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var reqDao dao.CreateUserRequest

		if err := c.BodyParser(&reqDao); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if err := validator.Struct(reqDao); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Invalid request")
		}

		if userService.CheckDuplicate(reqDao.Username) {
			return fiber.NewError(fiber.StatusBadRequest, "Username already exists")
		}

		user, err := userService.Register(reqDao.Username, reqDao.Password)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(dao.UserView{Id: user.ID, Username: user.Username})
	}
}
