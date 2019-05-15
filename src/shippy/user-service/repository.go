package main

import (
	"github.com/jinzhu/gorm"
	userProto "shippy/user-service/proto/user"
)

type Repository interface {
	Get(id string) (*userProto.User, error)
	GetAll() ([]*userProto.User, error)
	Create(*userProto.User) error
	GetByEmail(email string) (*userProto.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Get(id string) (*userProto.User, error) {
	var u *userProto.User
	u.Id = id
	if err := repo.db.First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) GetAll() ([]*userProto.User, error) {
	var users []*userProto.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Create(u *userProto.User) error {
	if err := repo.db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmail(email string) (*userProto.User, error) {
	user := &userProto.User{}
	if err := repo.db.Where("email = ? ", email).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
