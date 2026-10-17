package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/config/validation"
	"github.com/glesiosantos/crud_go_mongodb/src/controller/dto/request"
	"github.com/glesiosantos/crud_go_mongodb/src/model"
	"github.com/glesiosantos/crud_go_mongodb/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init create user controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidatorUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)

	domainResult, err := uc.service.CreateUser(domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
