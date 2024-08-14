package main

import (
	"nds_provision/config"
	"nds_provision/models"
)

func main() {
	db := config.Connect()
	db.AutoMigrate(models.User{})
	
}
