# Quake Game Log Parser

This project contains a Go script that parses Quake game logs to generate detailed reports for each match. The reports include total kills, player rankings, and a breakdown of deaths by cause.

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


To run the script, execute the following command from the terminal:

bash
Copiar código
go run main.go
This command will parse the log file and print the game reports to the standard output.

Output Example
The script will produce output in the following format for each game parsed:

yaml
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
Each game's output will list the total kills, detailed player statistics, and a count of deaths by each cause.

Contributing
Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes or improvements.