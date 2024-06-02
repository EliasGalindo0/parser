package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GameData Struct: This struct represents the data for a single game. It contains the following fields: TotalKills, Players, Kills and KillsByMenas
type GameData struct {
	TotalKills   int
	Players      map[string]bool
	Kills        map[string]int
	KillsByMeans map[string]int
}

// The function opens the "qgames.txt" file and creates a new scanner to read the file line by line.
func main() {
	file, err := os.Open("qgames.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// It initializes a games map to store the GameData for each game.
	scanner := bufio.NewScanner(file)
	games := make(map[int]*GameData)
	currentGame := 0

	// The function uses a regular expression (reKill) to parse each line and extract information about kills, including the killer, victim, and the means of the kill.
	reKill := regexp.MustCompile(`(\d+:\d+) Kill: (\d+) (\d+) (\d+): (.*) killed (.*) by (MOD_\w+)`)

	for scanner.Scan() {
		line := scanner.Text()

		// For each kill, the function updates the corresponding GameData struct in the games map.
		if strings.Contains(line, "InitGame:") {
			currentGame++
			games[currentGame] = &GameData{
				Players:      make(map[string]bool),
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
			continue
		}

		matches := reKill.FindStringSubmatch(line)
		if matches != nil {
			killer := matches[5]
			victim := matches[6]
			mod := matches[7]

			game := games[currentGame]
			game.TotalKills++

			if killer != "<world>" {
				game.Players[killer] = true
				game.Kills[killer]++
			}
			if victim != "<world>" {
				game.Players[victim] = true
			}

			// Increment death cause
			game.KillsByMeans[mod]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// After processing all the lines, the function prints a report for each game, including the total kills, the number of kills for each player, and the number of kills by each means.
	for gameNumber, data := range games {
		fmt.Printf("Game %d:\n", gameNumber)
		fmt.Printf("Total Kills: %d\n", data.TotalKills)
		fmt.Println("Players:")
		for player := range data.Players {
			fmt.Printf("- %s: %d kills\n", player, data.Kills[player])
		}
		fmt.Println("Deaths by means:")
		for mod, count := range data.KillsByMeans {
			fmt.Printf("  %s: %d\n", mod, count)
		}
		fmt.Println()
	}
}

// font: https://pt.stackoverflow.com/questions/183937/abrir-arquivo-texto-e-tratar-por-linha
// regex by ChatGPT
