package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func GetAsString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func GetAsInt(name string, defaultValue int) int {
	valueStr := GetAsString(name, "")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Printf("Error converting environment variable '%s' to int: %v\n", name, err)
		return defaultValue
	}
	return value
}