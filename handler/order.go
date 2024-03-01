package handler

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
	"github.com/JohnKucharsky/real_world_fiber_gorm/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var req domain.OrderRequest
	if err := req.Bind(c); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	id, err := h.orderStore.Create(req)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.orderStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetOrders(c *fiber.Ctx) error {
	res, err := h.orderStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetOneOrder(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.orderStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateOrder(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.OrderRequest
	if err := req.Bind(c); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	if err := h.orderStore.Update(req, id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	res, err := h.orderStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) DeleteOrder(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	if err := h.orderStore.Delete(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	res, err := h.orderStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}
