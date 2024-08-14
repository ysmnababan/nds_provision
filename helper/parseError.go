package helper

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNoUser    = errors.New("user doesn't exist")
	ErrQuery     = errors.New("query execution failed")
	ErrInvalidId = errors.New("invalid id")
	ErrBindJSON      = errors.New("unable to bind json")

	ErrNoRows        = errors.New("no rows in result set")
	ErrScan          = errors.New("row scanning failed")
	ErrUserExists    = errors.New("user already exist")
	ErrRowsAffected  = errors.New("unable to get affected row")
	ErrNoAffectedRow = errors.New("rows affected is 0")
	ErrLastInsertId  = errors.New("unable to get last insert id")
	ErrNoUpdate      = errors.New("data already exists")
	ErrParam         = errors.New("error or missing parameter")
	ErrCredential    = errors.New("password or email doesn't match")
)

func ParseError(err error, c *gin.Context) {
	var code int
	switch {
	case errors.Is(err, ErrInvalidId):
		fallthrough
	case errors.Is(err, ErrNoUser):
		code = http.StatusBadRequest
	case errors.Is(err,ErrBindJSON):
		fallthrough
	case errors.Is(err, ErrQuery):
		code = http.StatusInternalServerError
	}

	c.JSON(code, gin.H{"error": err.Error()})
}
