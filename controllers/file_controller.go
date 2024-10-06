package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/himanshukumar42/DistributeX/services"
	"github.com/himanshukumar42/DistributeX/utils"
)

// @Summary Upload a file
// @Description Upload a new file to the server
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 201 {object} utils.Response // Reference your utils.Response struct
// @Failure 400 {object} utils.Response // Reference your utils.ErrorResponse struct
// @Failure 500 {object} utils.Response
// @Router /api/v1/upload [post]
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

// @Summary Get all files
// @Description Retrieve a list of all uploaded files
// @Tags Files
// @Produce json
// @Success 200 {object} utils.Response // Reference your utils.Response struct
// @Failure 500 {object} utils.Response
// @Router /api/v1/files [get]
func GetFiles(c *gin.Context) {
	files, err := services.GetFiles()
	if err != nil {
		utils.Logger.Error("failed to fetch files: ", err)
		utils.ResponseWithError(c, http.StatusInternalServerError, "failed to fetch files")
		return
	}
	utils.ResponseWithSuccess(c, http.StatusOK, "Files retrieved Successfully", files)
}


// @Summary Download a file
// @Description Download a file by its ID
// @Tags Files
// @Produce application/octet-stream
// @Param id path string true "File ID"
// @Success 200 {file} []byte "File data"
// @Failure 404 {object} utils.Response // Reference your utils.ErrorResponse struct
// @Failure 500 {object} utils.Response
// @Router /api/v1/download/{id} [get]
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