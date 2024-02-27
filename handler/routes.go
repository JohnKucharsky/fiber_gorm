package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Register(r *fiber.App) {
	v1 := r.Group("/api")
	users := v1.Group("/users")
	//users.Get("/", h.GetAll)
	//users.Post("/sign_up", h.SignUp)
	//todo start here
	fmt.Println(users)
}
