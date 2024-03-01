package handler

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
	"github.com/JohnKucharsky/real_world_fiber_gorm/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	var req domain.UserRequest
	if err := req.Bind(c); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.userStore.Create(req)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {
	res, err := h.userStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetOneUser(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.userStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.UserRequest
	if err := req.Bind(c); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.userStore.Update(req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.userStore.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}
