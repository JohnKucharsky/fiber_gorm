package handler

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
	"github.com/JohnKucharsky/real_world_fiber_gorm/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var req domain.ProductRequest
	if err := req.Bind(c); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.productStore.Create(req)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) GetProducts(c *fiber.Ctx) error {
	res, err := h.productStore.GetMany()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) GetOneProduct(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.productStore.GetOne(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}

func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	var req domain.ProductRequest
	if err := req.Bind(c); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.productStore.Update(req, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(res)
}

func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	id, err := utils.GetID(c)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	res, err := h.productStore.Delete(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(res)
}
