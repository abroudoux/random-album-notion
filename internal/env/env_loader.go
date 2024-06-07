package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() (string, string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	envAPIKey := "NOTION_API_KEY"
	envPageId := "NOTION_PAGE_ID"
	notionApiKey := os.Getenv(envAPIKey)
	notionPageId := os.Getenv(envPageId)

	return notionApiKey, notionPageId
}