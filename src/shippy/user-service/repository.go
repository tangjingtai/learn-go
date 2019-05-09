package main

import (
	"github.com/jinzhu/gorm"
	userPb "shippy/user-service/proto/user"
)

type Repository interface {
	Get(id string) (*userPb.User, error)
	GetAll() ([]*userPb.User, error)
	Create(*userPb.User) error
	GetByEmailAndPassword(*userPb.User) (*userPb.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (repo *UserRepository) Get(id string) (*userPb.User, error) {
	var u *userPb.User
	u.Id = id
	if err := repo.db.First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (repo *UserRepository) GetAll() ([]*userPb.User, error) {
	var users []*userPb.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *UserRepository) Create(u *userPb.User) error {
	if err := repo.db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) GetByEmailAndPassword(u *userPb.User) (*userPb.User, error) {
	if err := repo.db.Find(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
