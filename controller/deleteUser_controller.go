package controller

import (
	"log"
	"nds_provision/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	user_id, err := strconv.Atoi(id)
	if id == "" || err != nil {
		log.Println("ERR CONTROLLER: id invalid")
		c.Error(helper.ErrInvalidId)
		return
	}

	err = h.UR.DeleteUser(uint(user_id))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user with id:" + id + " deleted succesfully"})
}
