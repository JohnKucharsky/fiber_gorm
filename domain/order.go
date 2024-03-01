package domain

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type OrderStore interface {
	Create(u OrderRequest) (int, error)
	GetMany() ([]*Order, error)
	GetOne(id int) (*Order, error)
	Update(u OrderRequest, id int) error
	Delete(id int) error
}

type Order struct {
	ID        uint      `json:"id" db:"id"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Product   Product   `json:"product" db:"product"`
	User      User      `json:"user" db:"user"`
}

type OrderRequest struct {
	ProductID int `json:"product_id" validate:"required"`
	UserID    int `json:"user_id" validate:"required"`
}

func (r *OrderRequest) Bind(
	c *fiber.Ctx,
) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	if err := NewValidator().Validate(r); err != nil {
		return err
	}

	return nil
}
