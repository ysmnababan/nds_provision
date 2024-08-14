package repository

import (
	"log"
	"nds_provision/helper"
	"nds_provision/models"
	"time"

	"gorm.io/gorm"
)

func (r *Repo) EditUser(id uint, in *models.EditUser) (*models.User, error) {
	var dbUser models.User
	check := r.DB.First(&dbUser, id)
	if check.Error == gorm.ErrRecordNotFound {
		log.Println("ERR REPO: no user found")
		return nil, helper.ErrNoUser
	}

	// name is changed
	if in.Name != "" {
		dbUser.Name = in.Name
	}

	// birthday is changed
	if in.Birthday != "" {
		parsedTime, err := time.Parse("2006-01-02", in.Birthday)
		if err != nil {
			log.Println("ERR REPO:", err)
			return nil, helper.ErrParam
		}
		dbUser.Birthday = &parsedTime
	}

	// update
	res := r.DB.Save(&dbUser)
	if res.Error != nil {
		log.Println("ERR REPO:", res.Error)
		return nil, helper.ErrQuery
	}

	return &dbUser, nil
}
