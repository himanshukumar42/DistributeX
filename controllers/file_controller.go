package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himanshukumar42/DistributeX/services"
	"github.com/himanshukumar42/DistributeX/utils"
)

func UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file") 
	if err != nil {
		utils.Logger.Error("failed to get file from request: ", err)
		utils.ResponseWithError(c, http.StatusBadRequest, "invalid file")
		return
	}

	fileID, err := services.UploadFile(fileHeader)
	if err != nil {
		utils.Logger.Error("failed to upload file: ", err)
		utils.ResponseWithError(c, http.StatusInternalServerError, "failed to upload the file")
		return
	}

	utils.ResponseWithSuccess(c, http.StatusOK, "file uploaded successfully", gin.H{"file_id": fileID})
}

func GetFiles(c *gin.Context) {
	files, err := services.GetFiles()
	if err != nil {
		utils.Logger.Error("failed to fetch files: ", err)
		utils.ResponseWithError(c, http.StatusInternalServerError, "failed to fetch files")
		return
	}
	utils.ResponseWithSuccess(c, http.StatusOK, "Files retrieved Successfully", files)
}

func DownloadFiles(c *gin.Context) {
	fileID := c.Param("id")
	fileData, filename, err := services.DownloadFile(fileID)
	if err != nil {
		utils.Logger.Error("failed to download file: ", err)
		utils.ResponseWithError(c, http.StatusInternalServerError, "failed to download the file")
		return
	}
	fmt.Println(filename)
	utils.Logger.Info("*******************", filename)
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/octet-stream", fileData)
}