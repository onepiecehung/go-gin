package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Declare variables for environment values
var (
	Port string
)

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

// Initialize the environment variables in an init() function
func init() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Assign environment variables to the variables
	Port = os.Getenv("PORT")

	DatabaseSetting.Type = os.Getenv("DB_TYPE")
	DatabaseSetting.User = os.Getenv("DB_USER")
	DatabaseSetting.Password = os.Getenv("DB_PASSWORD")
	DatabaseSetting.Host = os.Getenv("DB_HOST")
	DatabaseSetting.Name = os.Getenv("DB_NAME")
	DatabaseSetting.TablePrefix = os.Getenv("DB_TABLE_PREFIX")
}
