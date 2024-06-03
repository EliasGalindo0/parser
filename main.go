package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

// GameData Struct: This struct represents the data for a single game. It contains the following fields: TotalKills, Players, Kills and KillsByMeans.
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

			// the core of the kill counter
			if killer != "<world>" {
				game.Players[killer] = true
				game.Kills[killer]++
			}

			if victim != "<world>" {
				game.Players[victim] = true
				// If <world> kills a player, decrement the victim's kills
				if killer == "<world>" {
					game.Kills[victim]--
				}

				/* Alternative version - count and decrement player's kills when they kill themselves
				if killer == victim {
					game.Kills[killer]--
				}*/
			}

			// Increment death cause
			game.KillsByMeans[mod]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Get a sorted list of game numbers
	gameNumbers := make([]int, 0, len(games))
	for gameNumber := range games {
		gameNumbers = append(gameNumbers, gameNumber)
	}
	sort.Ints(gameNumbers)

	// After processing all the lines, the function prints a report for each game in ascending order of game numbers.
	for _, gameNumber := range gameNumbers {
		data := games[gameNumber]
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
