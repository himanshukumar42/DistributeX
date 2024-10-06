package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/himanshukumar42/DistributeX/models"
	_ "github.com/lib/pq"
)

var db *sql.DB

func RunSQLScript(filePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("could not read schema file: %w", err)
	}

	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		trimmedQuery := strings.TrimSpace(query)
		if trimmedQuery == "" {
			continue
		}

		_, err := db.Exec(trimmedQuery)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v, error: %w", trimmedQuery, err)
		}
	}

	fmt.Println("SQL Schema executed Successfully")
	return nil
}

func InitDB(connStr, schemaFilePath string) error {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	
	if err := db.Ping(); err != nil {
		return err
	}

	// Execute Schema.sql
	if err := RunSQLScript(schemaFilePath); err != nil {
		return err
	}

	fmt.Println("Database initialized successfully with schema!")
	return nil
}

func SaveFilePart(fileID string, partNumber int, data []byte) error {
	query := `INSERT INTO file_parts(file_id, part_number, data) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, fileID, partNumber, data)
	return err
}

func SaveFileMetadata(fileID, filename string, partCount int) error {
	query := `INSERT INTO file_metadata (file_id, filename, part_count) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, fileID, filename, partCount)
	return err
}


func GetAllFiles() ([]models.FileMetadata, error) {
	query := `SELECT file_id, filename, part_count FROM file_metadata`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []models.FileMetadata
	for rows.Next() {
		var file models.FileMetadata
		if err := rows.Scan(&file.FileID, &file.Filename, &file.PartCount); err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func GetFileMetadata(fileID string) (*models.FileMetadata, error) {
	query := `SELECT file_id, filename, part_count FROM file_metadata WHERE file_id=$1`
	row := db.QueryRow(query, fileID)

	var file models.FileMetadata
	if err := row.Scan(&file.FileID, &file.Filename, &file.PartCount); err != nil {
		return nil, err
	}
	return &file, nil
}

func GetFilePart(fileID string, partNumber int) ([]byte, error) {
	query := `SELECT data FROM file_parts WHERE file_id=$1 AND part_number=$2`
	row := db.QueryRow(query, fileID, partNumber)

	var data []byte
	if err := row.Scan(&data); err != nil {
		return nil, err
	}
	return data, nil
}