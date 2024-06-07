package spotify

import (
	"os/exec"
)

func toggleShuffle() (output string, error error) {
	cmd := exec.Command("spotify", "toggle", "shuffle")
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(stdout), nil
}

func PlayAlbum(s string) (output string, error error) {
	cmd := exec.Command("spotify", "play", s)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}

	toggleSuffleStdout, err := toggleShuffle()

	if err != nil {
		return "", err
	}

	if toggleSuffleStdout == "Spotify shuffling set to true" {
		toggleShuffle()
	}

	return string(stdout), nil
}