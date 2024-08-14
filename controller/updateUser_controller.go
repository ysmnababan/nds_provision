package controller

import (
	"log"
	"nds_provision/helper"
	"nds_provision/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user_id, err := strconv.Atoi(id)
	if id == "" || err != nil {
		c.Error(helper.ErrInvalidId)
		return
	}

	var in models.EditUser
	if c.ShouldBindJSON(&in) != nil {
		log.Println("ERR CONTROLLER: binding")
		c.Error(helper.ErrBindJSON)
		return
	}

	res, err := h.UR.EditUser(uint(user_id), &in)
	if err != nil {
		c.Error(err)
		return
	}
	res.Pwd = "***"
	c.JSON(http.StatusOK, gin.H{"Updated User": res})
}
