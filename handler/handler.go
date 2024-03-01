package handler

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/domain"
)

type Handler struct {
	userStore    domain.UserStore
	productStore domain.ProductStore
	orderStore   domain.OrderStore
}

func NewHandler(
	us domain.UserStore,
	ps domain.ProductStore,
	os domain.OrderStore,
) *Handler {
	return &Handler{
		userStore:    us,
		productStore: ps,
		orderStore:   os,
	}
}
