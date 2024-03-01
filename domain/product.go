package domain

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type ProductStore interface {
	Create(u ProductRequest) (*Product, error)
	GetMany() ([]*Product, error)
	GetOne(id int) (*Product, error)
	Update(u ProductRequest, id int) (*Product, error)
	Delete(id int) (*Product, error)
}

type Product struct {
	ID           uint      `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	Name         string    `json:"name" db:"name"`
	SerialNumber string    `json:"serial_number" db:"serial_number"`
}

type ProductRequest struct {
	Name         string `json:"name" validate:"required"`
	SerialNumber string `json:"serial_number" validate:"required"`
}

func (r *ProductRequest) Bind(
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
