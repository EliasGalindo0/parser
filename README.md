# Quake Game Log Parser

This project contains a Go script that parses Quake game logs to generate detailed reports for each match. The reports include total kills, player rankings, and a breakdown of deaths by cause.

## Aspects
The key aspects of the code are:

File Processing: The code uses the 
os.Open() and bufio.NewScanner() functions to read the game data from the "qgames.txt" file line by line.

Regular Expression: The code uses a regular expression to parse each line and extract the relevant information about the kills.

Data Structures: The code uses a GameData struct to store the data for each game, and maps to keep track of the players, kills, and kill means.

Game Reporting: The code generates a report for each game, printing the total kills, the number of kills for each player, and the number of kills by each means.

Overall, this code is designed to process game data from a file and generate detailed reports for each game. The use of data structures and regular expressions makes the code efficient and flexible in handling the game data.

## Features

- **Match Summaries**: Provides a summary for each match, including total kills and individual player statistics.
- **Player Rankings**: Ranks players based on their kill counts.
- **Death Cause Analysis**: Breaks down the number of deaths by each cause specified in the game's means of death enumeration.

## Prerequisites

To run this script, you will need:
- [Go (Golang)](https://golang.org/dl/) installed on your system.
- A log file named `qgames.txt` located in the same directory as the script, or modify the script to point to the correct file location.

## Installation

Clone this repository or download the files directly:

```bash
git clone https://your-repository-url.git
cd quake-log-parser
```

Ensure that the qgames.txt file is in the same directory as your script or update the file path in the script to match the location of your log file.

## Usage

To run the script, execute the following command from the terminal:

```bash
Copiar código
go run main.go
```

This command will parse the log file and print the game reports to the standard output.

## Output Example
The script will produce output in the following format for each game parsed:

```yaml
Copiar código
Game 1:
Total Kills: 45
Players:
- Dono da bola: 5 kills
- Isgalamido: 18 kills
- Zeh: 20 kills
Deaths by means:
  MOD_SHOTGUN: 10
  MOD_RAILGUN: 2
  MOD_GAUNTLET: 1
```

Each game's output will list the total kills, detailed player statistics, and a count of deaths by each cause.

## Contributing
Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes or improvements.