package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/shop/golang/controllers"
	"github.com/shop/golang/middleware"
)

func UserRoutes(incomingRoute *gin.Engine) {
	incomingRoute.Use(middleware.Authentificate())
	incomingRoute.GET("/users", controller.GetUsers())
	incomingRoute.GET("/users/user_id", controller.GetUser())
}
