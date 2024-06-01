package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type GameData struct {
	TotalKills int
	Players    map[string]bool
	Kills      map[string]int
}

func main() {
	file, err := os.Open("qgames.txt")
	if err != nil {
		fmt.Printf("Erro ao abrir o arquivo: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameData := make(map[int]*GameData)
	currentGame := 0

	for scanner.Scan() {
		line := scanner.Text()

		// Verify if a new game starts
		if strings.Contains(line, "InitGame:") {
			currentGame++
			gameData[currentGame] = &GameData{
				Players: make(map[string]bool),
				Kills:   make(map[string]int),
			}
			continue
		}

		// Regex to capture kill events
		re := regexp.MustCompile(`(\d+:\d+) Kill: (\d+) (\d+) (\d+): (.*) killed (.*) by (.*)`)
		matches := re.FindStringSubmatch(line)

		if matches != nil {
			killer := matches[5]
			victim := matches[6]
			mod := matches[7]

			if mod != "TRIGGER_HURT" && currentGame > 0 {
				game := gameData[currentGame]

				game.TotalKills++

				// Handle <world> kills a player
				if killer == "<world>" {
					if victim != "<world>" {
						game.Kills[victim]--
						game.Players[victim] = true
					}
				} else {
					game.Kills[killer]++
					game.Players[killer] = true
					game.Players[victim] = true
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Erro ao ler o arquivo: %v\n", err)
		return
	}

	// Print results
	for gameNumber, data := range gameData {
		fmt.Printf("Game %d:\n", gameNumber)
		fmt.Printf("Total Kills: %d\n", data.TotalKills)
		fmt.Println("Players:")
		for player := range data.Players {
			fmt.Printf("- %s: %d kills\n", player, data.Kills[player])
		}
		fmt.Println()
	}
}
