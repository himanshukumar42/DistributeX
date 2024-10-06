package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}


func CreateErrorResponse(message string) Response {
	return Response{
		Status: "error",
		Message: message,
	}
}

func CreateSuccessResponse(message string, data interface{}) Response {
	return Response{
		Status: "success",
		Message: message,
		Data: data,
	}
}

func ResponseWithError(c *gin.Context, code int,  message string) {
	response := CreateErrorResponse(message)
	c.JSON(code,response)
}

func ResponseWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	response := CreateSuccessResponse(message, data)
	c.JSON(code, response)
}