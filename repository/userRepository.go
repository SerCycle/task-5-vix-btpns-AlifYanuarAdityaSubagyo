package repository

import (
	"github.com/SerCycle/BTPNFinalProject/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.User, error)
	FindByID(ID int) (model.User, error)
	Create(user model.User) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FindAll() ([]model.User, error) {
	var users []model.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) FindByID(ID int) (model.User, error) {
	var user model.User

	err := r.db.Find(&user, ID).Error

	return user, err
}

func (r *userRepository) Create(user model.User) (model.User, error) {
	err := r.db.Find(&user).Error

	return user, err
}
