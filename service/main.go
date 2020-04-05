package main

import (
	"github.com/gin-gonic/gin"
	"user-management-service/api"
)

func main() {

	controller := api.UserController{Router: gin.Default()}
	controller.V1()
	_ = controller.Router.Run("localhost:8080")

}
