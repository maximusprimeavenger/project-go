package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shop/golang/controllers"
)

func AuthRoutes(incomingRoute *gin.Engine) {
	incomingRoute.POST("users/signup", controller.Signup())
	incomingRoute.POST("users/login", controller.Login())
}
