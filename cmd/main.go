package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

func main () {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	envApiKey := "NOTION_API_KEY"
	envPageId := "NOTION_PAGE_ID"
	notionApiKey := os.Getenv(envApiKey)
	notionPageId := os.Getenv(envPageId)

	client := notionapi.NewClient(notionapi.Token(notionApiKey))

	fmt.Println(notionApiKey, notionPageId, client)

	page, err := client.Page.Get(context.Background(), notionapi.PageID(notionPageId))

	if err != nil {
		fmt.Println("Error connecting page", err)
	}

	fmt.Println(page)
}