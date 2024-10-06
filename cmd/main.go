package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/himanshukumar42/DistributeX/api"
	"github.com/himanshukumar42/DistributeX/config"
	"github.com/himanshukumar42/DistributeX/utils"
)


func main() {

	config.LoadConfig()
	utils.SetupLogger()

	
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12*time.Hour,
	}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Health Ok"})
	})
	
	api.SetupRoutes(router)

	utils.Logger.Info("Starting Server.....")
	if err := router.Run(":8080"); err != nil {
		utils.Logger.Fatal("Failed to start server: ", err)
	}
}