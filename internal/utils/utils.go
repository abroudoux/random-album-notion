package uils

import "math/rand"

func ChooseRandomAlbum(albums []string) string {
	randomNb := rand.Intn(len(albums))
	randomAlbum := albums[randomNb]

	return randomAlbum
}