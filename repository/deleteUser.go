package repository

import (
	"log"
	"nds_provision/helper"
	"nds_provision/models"

	"gorm.io/gorm"
)

func (r *Repo) DeleteUser(id uint) error {
	check := r.DB.First(&models.User{}, id)
	if check.Error == gorm.ErrRecordNotFound {
		log.Println("ERR REPO: no user found")
		return helper.ErrNoUser
	}

	res := r.DB.Delete(&models.User{}, id)
	if res.Error != nil {
		log.Println("REPO ERR: ", res.Error)
		return helper.ErrQuery
	}
	return nil
}
