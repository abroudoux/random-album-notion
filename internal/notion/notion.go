package notion

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	types "github.com/abroudoux/random-album/internal/types"
)

func GetTodos(notionPageId string, notionAPIKey string, notionAPIVersion string) ([]string, error) {
    url := "https://api.notion.com/v1/blocks/" + notionPageId + "/children"
    var bearer = "Bearer " + notionAPIKey

    client := &http.Client{}
    var todos []string
    var startCursor *string

    for {
        notionResponse, err := fetchNotionResponse(url, bearer, notionAPIVersion, startCursor, client)

        if err != nil {
            return nil, fmt.Errorf("failed to fetch notion response: %v", err)
        }

        todos = appendTodos(todos, notionResponse)

        if !notionResponse.HasMore {
            break
        }

        startCursor = notionResponse.NextCursor
    }

    return todos, nil
}

func fetchNotionResponse(url, bearer, notionAPIVersion string, startCursor *string, client *http.Client) (types.NotionResponse, error) {
    req, err := http.NewRequest("GET", url, nil)

    if err != nil {
        return types.NotionResponse{}, fmt.Errorf("error creating request: %v", err)
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
        return types.NotionResponse{}, fmt.Errorf("error making request: %v", err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        return types.NotionResponse{}, fmt.Errorf("error reading response body: %v", err)
    }

    var notionResponse types.NotionResponse

    if err := json.Unmarshal(body, &notionResponse); err != nil {
        return types.NotionResponse{}, fmt.Errorf("error unmarshalling JSON response: %v", err)
    }

    return notionResponse, nil
}

func appendTodos(todos []string, notionResponse types.NotionResponse) []string {
    for _, block := range notionResponse.Results {
        if block.Type == "to_do" && block.ToDo != nil && !block.ToDo.Checked {
            for _, richText := range block.ToDo.RichText {
                todos = append(todos, richText.Text.Content)
            }
        }
    }

    return todos
}
