package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const notionAPIVersion string = "2022-06-28"

func Load() (string, string) {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	envApiKey := "NOTION_API_KEY"
	envPageId := "NOTION_PAGE_ID"
	notionApiKey := os.Getenv(envApiKey)
	notionPageId := os.Getenv(envPageId)

	return notionApiKey, notionPageId
}

func GetPage(notionPageId string, notionApiKey string, notionAPIVersion string) (string, error) {
	url := "https://api.notion.com/v1/blocks/" + notionPageId + "/children"

	var bearer = "Bearer " + notionApiKey

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", notionAPIVersion)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

    if err != nil {
        log.Println("Error while reading the response bytes:", err)
    }

	return string([]byte(body)), nil
}

func main () {
	notionApiKey, notionPageId := Load()

	res, err := GetPage(notionPageId, notionApiKey, notionAPIVersion)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}