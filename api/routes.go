package api

import (
	"github.com/gin-gonic/gin"
	"github.com/himanshukumar42/DistributeX/controllers"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/upload", controllers.UploadFile)
		v1.GET("/files", controllers.GetFiles)
		v1.GET("/download/:id", controllers.DownloadFiles)
	}
}