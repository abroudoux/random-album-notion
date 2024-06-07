package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const notionAPIVersion string = "2022-06-28"
type Block struct {
	Object string `json:"object"`
	ID     string `json:"id"`
	Type   string `json:"type"`
	ToDo   *struct {
		RichText []struct {
			Text struct {
				Content string `json:"content"`
			} `json:"text"`
		} `json:"rich_text"`
		Checked bool `json:"checked"`
	} `json:"to_do,omitempty"`
}

type NotionResponse struct {
	Object  string  `json:"object"`
	Results []Block `json:"results"`
}

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

func GetPage(notionPageId string, notionApiKey string, notionAPIVersion string) ([]string, error) {
	url := "https://api.notion.com/v1/blocks/" + notionPageId + "/children"
	var bearer = "Bearer " + notionApiKey

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Notion-Version", notionAPIVersion)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var notionResponse NotionResponse
	err = json.Unmarshal(body, &notionResponse)

	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON response: %v", err)
	}

	var todos []string
	
	for _, block := range notionResponse.Results {
		if block.Type == "to_do" && block.ToDo != nil && !block.ToDo.Checked {
			for _, richText := range block.ToDo.RichText {
				todos = append(todos, richText.Text.Content)
			}
		}
	}

	return todos, nil
}

func ChooseRandomAlbum(albums []string) string {
	randomNb := rand.Intn(len(albums))
	randomAlbum := albums[randomNb]

	return randomAlbum
}

func main() {
	notionApiKey, notionPageId := Load()

	todos, err := GetPage(notionPageId, notionApiKey, notionAPIVersion)

	if err != nil {
		log.Fatalf("Error fetching page: %v", err)
	}

	randomAlbum := ChooseRandomAlbum(todos)

	fmt.Println(randomAlbum)
}
