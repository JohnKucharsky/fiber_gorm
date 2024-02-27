package user

import "github.com/JohnKucharsky/real_world_fiber_gorm/model"

type Store interface {
	Create(u *model.User) error
	GetMany() (*[]model.User, error)
}
