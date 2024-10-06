package services

import (
	"mime/multipart"
	"sync"

	"github.com/himanshukumar42/DistributeX/models"
	"github.com/himanshukumar42/DistributeX/repository"
	"github.com/himanshukumar42/DistributeX/utils"
)

func UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	fileID := utils.GenerateID()

	parts, err := utils.SplitFile(fileHeader)
	if err != nil {
		return "", err
	}

	if err := repository.SaveFileMetadata(fileID,fileHeader.Filename, len(parts)); err != nil {
		utils.Logger.Errorf("Failed to save metadata for file %s: %v", fileID, err)
		return "", err
	}

	
	var wg sync.WaitGroup
	errChan := make(chan error, len(parts))

	for idx, part := range parts {
		wg.Add(1)
		go func (partData []byte, index int)  {
			defer wg.Done()
			if err := repository.SaveFilePart(fileID, index, partData); err != nil {
				utils.Logger.Errorf("failed to save part %d of file %s: %v", index, fileID, err)
				errChan <- err
			}
		}(part, idx)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return "", <-errChan
	}

	return fileID, nil
}

func GetFiles() ([]models.FileMetadata, error) {
	return repository.GetAllFiles()
}

func DownloadFile(fileID string) ([]byte, string, error) {
	metadata, err := repository.GetFileMetadata(fileID)
	if err != nil {
		return nil, "", err
	}

	parts := make([][]byte, metadata.PartCount)
	var wg sync.WaitGroup
	errChan := make(chan error, metadata.PartCount)

	for i := 0; i < metadata.PartCount; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			part, err := repository.GetFilePart(fileID, index)
			if err != nil {
				utils.Logger.Errorf("failed to retrieve part %d of file %s: %v", index, fileID, err)
				errChan <- err
				return
			}
			parts[index] = part
		}(i)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return nil, "", <-errChan
	}

	mergedFile, err := utils.MergeFileParts(parts)
	if err != nil {
		return nil, "", err
	}
	return mergedFile, metadata.Filename, nil
}