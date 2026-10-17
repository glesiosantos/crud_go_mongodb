package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/config/errorsc"
	"github.com/glesiosantos/crud_go_mongodb/src/config/logger"
	"github.com/glesiosantos/crud_go_mongodb/src/view"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("email")

	userDomain, err := uc.service.FindUserByEmail(userEmail)

	if err != nil {
		logger.Error("Error trying to validate email", err, zap.String("journey", "findUserById Controller"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	userId := c.Param("id")

	if _, err := bson.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "findUserById Controller"))
		errorMessage := errorsc.NewBadRequestError("UserId is not a valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserById(userId)

	if err != nil {
		logger.Error("Error trying to validate userId", err, zap.String("journey", "findUserById Controller"))
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
