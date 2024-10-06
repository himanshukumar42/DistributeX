package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

func SplitFile(fileHeader *multipart.FileHeader) ([][]byte, error) {
	file, err := fileHeader.Open()
	if err != nil {
		Logger.Error("Failed to open uploaded file: ", err)
		return nil, err
	}
	defer file.Close()

	const chunkSize = 1024 * 1024 // 1 MB per chunk

	var parts [][]byte
	buffer := make([]byte, chunkSize)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			Logger.Error("failed to read file: ",err)
			return nil, err
		}
		if n == 0 {
			break
		}
		part := make([]byte, n)
		copy(part, buffer[:n])
		parts = append(parts, part)
	}

	Logger.Info("file split into parts", len(parts))
	return parts, nil
}


func MergeFileParts(parts [][]byte) ([]byte, error) {
	var buffer bytes.Buffer
	for _, part := range parts {
		_, err := buffer.Write(part)
		if err != nil {
			Logger.Error("failed to write file part: ", err)
			return nil, err
		}
	}
	mergedFile := buffer.Bytes()
	Logger.Info("merged file size bytes ", len(mergedFile))
	return mergedFile, nil
}

func ReadFileParts(fileID string) ([][]byte, error) {
	return nil, fmt.Errorf("ReadFileParts not implemented")
}
