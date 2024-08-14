package controller

import (
	"nds_provision/helper"
	"nds_provision/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UR repository.UserRepo
}

func (h *UserController) GetUserID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.Error(helper.ErrInvalidId)
		return
	}

	user_id, err := strconv.Atoi(id)
	if err != nil {
		c.Error(helper.ErrInvalidId)
		return
	}

	res, err := h.UR.GetUserID(uint(user_id))
	if err != nil {
		c.Error(err)
		return
	}

	res.Pwd = "*****"
	c.JSON(http.StatusOK, gin.H{"user": res})
}
