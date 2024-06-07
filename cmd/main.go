package main

import (
	"fmt"

	env "github.com/abroudoux/random-album/internal/env"
	notion "github.com/abroudoux/random-album/internal/notion"
	spotify "github.com/abroudoux/random-album/internal/spotify"
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
	stdout, err := spotify.PlayAlbum(randomAlbum)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(stdout)
}
