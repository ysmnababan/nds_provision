package main

import (
	"nds_provision/config"
	"nds_provision/controller"
	"nds_provision/helper"
	"nds_provision/models"
	"nds_provision/repository"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := config.Connect()
	db.AutoMigrate(models.User{})

	// create initial data
	// dataSeed(db)

	//initialize router
	router := gin.Default()
	router.Use(gin.Recovery())

	// middleware to handle error
	router.Use(func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			ctx.Next()

			if len(ctx.Errors) > 0 {
				err := ctx.Errors[0].Err
				helper.ParseError(err, ctx)
			}
		}
	}())

	userRepo := &repository.Repo{DB: db}
	userController := &controller.UserController{UR: userRepo}

	//crud endpoint
	router.GET("/user/:id", userController.GetUserID)
	router.GET("/users", userController.GetAllUsers)
	router.POST("/user", userController.CreateUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	// run the server in port 8080
	router.Run(":8080")
}

func dataSeed(db *gorm.DB) {
	temp := time.Now()
	users := []*models.User{
		{Name: "Name1", Email: "email1", Pwd: "pwd1"},
		{Name: "Name2", Email: "email2", Pwd: "pwd2", Birthday: &temp},
		{Name: "Name3", Email: "email3", Pwd: "pwd3"},
	}

	db.Create(users)
}
