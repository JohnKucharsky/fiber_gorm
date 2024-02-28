package handler

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/model"
	"github.com/gofiber/fiber/v2"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	} `json:"user"`
}

func (r *userRegisterRequest) bind(
	c *fiber.Ctx,
	u *model.User,
	v *Validator,
) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if err := v.Validate(r); err != nil {
		return err
	}

	u.UserName = r.User.Username
	u.Email = r.User.Email
	u.Password = r.User.Password

	return nil
}
