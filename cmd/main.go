package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const notionAPIVersion string = "2022-06-28"
type NotionResponse struct {
	Object     string    `json:"object"`
	Results    []Block   `json:"results"`
	NextCursor *string   `json:"next_cursor"`
	HasMore    bool      `json:"has_more"`
}

type Block struct {
	Type string `json:"type"`
	ToDo *ToDo  `json:"to_do"`
}

type ToDo struct {
	Checked  bool       `json:"checked"`
	RichText []RichText `json:"rich_text"`
}

type RichText struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
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

func GetTodos(notionPageId string, notionApiKey string, notionAPIVersion string) ([]string, error) {
	url := "https://api.notion.com/v1/blocks/" + notionPageId + "/children"
	var bearer = "Bearer " + notionApiKey

	client := &http.Client{}
	var todos []string
	var startCursor *string
	hasMore := true

	for hasMore {
		req, err := http.NewRequest("GET", url, nil)

		if err != nil {
			return nil, fmt.Errorf("error creating request: %v", err)
		}

		req.Header.Add("Authorization", bearer)
		req.Header.Add("Notion-Version", notionAPIVersion)

		if startCursor != nil {
			q := req.URL.Query()
			q.Add("start_cursor", *startCursor)
			req.URL.RawQuery = q.Encode()
		}

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

		for _, block := range notionResponse.Results {
			if block.Type == "to_do" && block.ToDo != nil && !block.ToDo.Checked {
				for _, richText := range block.ToDo.RichText {
					todos = append(todos, richText.Text.Content)
				}
			}
		}

		hasMore = notionResponse.HasMore
		startCursor = notionResponse.NextCursor
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

	todos, err := GetTodos(notionPageId, notionApiKey, notionAPIVersion)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	randomAlbum := ChooseRandomAlbum(todos)

	fmt.Println(randomAlbum)
}
