package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/model/services"
)

func NewUserControllerInterface(service services.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type userControllerInterface struct {
	service services.UserDomainService
}
