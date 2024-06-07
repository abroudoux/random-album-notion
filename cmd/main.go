package main

import (
	"fmt"

	env "github.com/abroudoux/random-album/internal/env"
	notion "github.com/abroudoux/random-album/internal/notion"
	utils "github.com/abroudoux/random-album/internal/utils"
)

const notionAPIVersion string = "2022-06-28"

func main() {
	notionAPIKey, notionPageId := env.Load()

	todos, err := notion.GetTodos(notionPageId, notionAPIKey, notionAPIVersion)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	randomAlbum := utils.ChooseRandomAlbum(todos)

	fmt.Println(randomAlbum)
}
