package models

type FileMetadata struct {
	FileID string `json:"file_id"`
	Filename string `json:"filename"`
	PartCount int `json:"part_count"`
}