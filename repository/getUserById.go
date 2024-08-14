package repository

import (
	"log"
	"nds_provision/helper"
	"nds_provision/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepo interface {
	GetUserID(id uint) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	CreateUser(in *models.User) (*models.User, error)

	DeleteUser(id uint) error
}

func (r *Repo) GetUserID(id uint) (*models.User, error) {
	var user models.User

	res := r.DB.First(&user, id)
	if res.Error == gorm.ErrRecordNotFound {
		log.Println("ERR REPO:", res.Error)
		return nil, helper.ErrNoUser
	}

	return &user, nil
}
