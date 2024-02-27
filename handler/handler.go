package handler

import "github.com/JohnKucharsky/real_world_fiber_gorm/user"

type Handler struct {
	userStore user.Store
	validator *Validator
}

func NewHandler(us user.Store) *Handler {
	v := NewValidator()

	return &Handler{
		userStore: us,
		validator: v,
	}
}
