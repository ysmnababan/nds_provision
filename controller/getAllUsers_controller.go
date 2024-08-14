package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserController) GetAllUsers(c *gin.Context) {
	res, err := h.UR.GetAllUsers()
	if err != nil {
		c.Error(err)
		return
	}
	for _, val := range res {
		val.Pwd = "***"
	}

	c.JSON(http.StatusOK, gin.H{"users": res})
}
