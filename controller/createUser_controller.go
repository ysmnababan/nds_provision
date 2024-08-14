package controller

import (
	"log"
	"nds_provision/helper"
	"nds_provision/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *UserController) CreateUser(c *gin.Context) {
	// get user data
	var in models.CreateUser
	err := c.ShouldBindJSON(&in)
	if err != nil {
		log.Println("ERR CONTROLLER:", err)
		c.Error(helper.ErrBindJSON)
		return
	}

	//validate user
	if in.Email == "" || in.Name == "" || in.Pwd == "" {
		log.Println("ERR CONTROLLER: invalid body")
		c.Error(helper.ErrParam)
		return
	}

	req := &models.User{
		Name:  in.Name,
		Email: in.Email,
		Pwd:   in.Pwd,
	}
	// validate birthday
	if in.Birthday != "" {
		parsedTime, err := time.Parse("2006-01-02", in.Birthday)
		if err != nil {
			log.Println("ERR CONTROLLER:", err)
			c.Error(helper.ErrParam)
			return
		}
		req.Birthday = &parsedTime
	}
	res, err := h.UR.CreateUser(req)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"New User: ": res})
}
