package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/shop/golang/routes"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("Fail to set the proxies: %v", err)
	}

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access successfully on api-1"})
	})

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access successfully on api-2"})
	})

	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
