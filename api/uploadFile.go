package api

import (
	"io"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)


func UploadFileHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "File error"})
		return
	}
	defer file.Close()
	
	chunkSize := 1024 * 1024	// 1 MB Chunks
	parts := splitFile(file, chunkSize)

	var wg sync.WaitGroup
	for idx, part := range parts {
		wg.Add(1)
		go func(idx int, part []byte) {
			defer wg.Done()
			uploadPartToDB(idx, part)
		}(idx, part)
	}
}

func splitFile(file io.Reader, chunkSize int) [][]byte {
	var parts [][]byte
	buf := make([]byte, chunkSize)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			parts = append(parts, buf[:n])
		}
		if err != nil {
			break
		}
	}
	return parts
}

func uploadPartToDB(idx int, part []byte) {
	// Logic to store part in the database
}
