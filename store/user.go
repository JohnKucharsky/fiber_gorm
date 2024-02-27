package store

import (
	"github.com/JohnKucharsky/real_world_fiber_gorm/model"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (us *UserStore) Create(u *model.User) error {
	return us.db.Create(u).Error
}

func (us *UserStore) GetMany() (*[]model.User, error) {
	var m []model.User
	if err := us.db.Find(&m).Error; err != nil {
		return nil, err
	}
	return &m, nil
}
