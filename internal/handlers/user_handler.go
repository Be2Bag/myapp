package handlers

import (
	"strconv"

	"github.com/be2bag/myapp/internal/core/domain"
	"github.com/be2bag/myapp/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(app *fiber.App, userService ports.UserService) {
	handler := &userHandler{userService: userService}

	app.Post("/users", handler.RegisterUser)
	app.Delete("/users/:id", handler.RemoveUser)
	app.Put("/users/:id", handler.ModifyUser)
	app.Get("/users/:id", handler.FindUserByID)
	app.Get("/users", handler.ListUsers)
}

func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	createdUser, err := h.userService.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser)
}

func (h *userHandler) RemoveUser(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	err := h.userService.RemoveUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *userHandler) ModifyUser(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	user.ID = uint(id)

	updatedUser, err := h.userService.ModifyUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(updatedUser)
}

func (h *userHandler) FindUserByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	user, err := h.userService.FindUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	return c.JSON(user)
}

func (h *userHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}
