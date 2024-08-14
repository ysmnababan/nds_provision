package repository

import (
	"log"
	"nds_provision/helper"
	"nds_provision/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (r *Repo) isUserExist(email string) (bool, error) {
	var user models.User
	res := r.DB.Where("email=?", email).First(&user)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return false, nil
		}
		log.Println("ERR REPO:", res.Error)
		return false, helper.ErrQuery
	}
	return true, nil
}

func (r *Repo) CreateUser(in *models.User) (*models.User, error) {
	// ensure no duplicate email
	isExist, err := r.isUserExist(in.Email)
	if err != nil {
		return nil, err
	}
	if isExist {
		return nil, helper.ErrUserExists
	}

	// hashing pwd
	hashedpwd, _ := bcrypt.GenerateFromPassword([]byte(in.Pwd), bcrypt.DefaultCost)
	in.Pwd = string(hashedpwd)
	res := r.DB.Create(in)
	if res.Error != nil {
		log.Println("ERR REPO: ", res.Error)
		return nil, helper.ErrQuery
	}

	in.Pwd = "***"
	return in, nil
}
