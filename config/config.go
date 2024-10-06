package config

import (
	"fmt"
	"os"

	"github.com/himanshukumar42/DistributeX/repository"
	"github.com/himanshukumar42/DistributeX/utils"
)

type Config struct {
	DBHost string
	DBPort string
	DBUser string
	DBPassword string
	DBName string
}

var AppConfig Config


func LoadConfig() {
	AppConfig.DBHost = getEnv("DB_HOST", "localhost")
	AppConfig.DBPort = getEnv("DB_PORT", "5434")
	AppConfig.DBUser = getEnv("DB_USER", "user")
	AppConfig.DBPassword = getEnv("DB_PASSWORD", "password")
	AppConfig.DBName = getEnv("DB_NAME","new_storage_db")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)

	currentDir, err := os.Getwd()
	if err != nil {
		utils.Logger.Fatal("failed to get current directory: ", err)
	}
	fmt.Println("Current Directory: ", currentDir)
	// schemeFilePath := filepath.Join(currentDir, "Schema.sql")
	if err := repository.InitDB(connStr, "./Schema.sql"); err != nil {
		utils.Logger.Fatal("failed to connect to the database: ", err)
	}

}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
