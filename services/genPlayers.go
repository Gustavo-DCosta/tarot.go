package services

import (
	"log"
	"os"
)

func GeneratePlayers() []string {
	playersList := []string{
		"player1",
		"player2",
		"player3",
		"player4",
		"player5",
	}

	for _, p := range playersList {
		file, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}

	return playersList
}
