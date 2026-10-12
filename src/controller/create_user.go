package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/config/errors"
	"github.com/glesiosantos/crud_go_mongodb/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := errors.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, errors = %s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(userRequest)
}
