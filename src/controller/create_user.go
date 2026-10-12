package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/config/errors"
)

func CreateUser(c *gin.Context) {
	err := errors.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
