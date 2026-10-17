package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/glesiosantos/crud_go_mongodb/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:id", userController.FindUserById)
	r.GET("/getUserByEmail/:email", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
