package handler

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/model"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (h *Handler) SignUp(c *fiber.Ctx) error {
	var u model.User

	req := &userRegisterRequest{}
	if err := req.bind(c, &u, h.validator); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	if err := h.userStore.Create(&u); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err.Error())
	}

	return c.Status(http.StatusCreated).JSON(newUserResponse(&u))
}
