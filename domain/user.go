package domain

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type UserStore interface {
	Create(u UserRequest) (*User, error)
	GetMany() ([]*User, error)
	GetOne(id int) (*User, error)
	Update(u UserRequest, id int) (*User, error)
	Delete(id int) (*User, error)
}

type User struct {
	ID        uint      `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  *string   `json:"last_name" db:"last_name"`
}

type UserRequest struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name"`
}

func (r *UserRequest) Bind(
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
