package repository

import "nds_provision/models"

func (r *Repo) GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	r.DB.Find(&users)
	return users, nil
}
