package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

const API_URL = "https://api-web.nhle.com/v1"

type Config struct {
	MatchDate       string
	TeamAbreviation string
}

type GameWeek struct {
	Date          string `json:"date"`
	DayAbbrev     string `json:"dayAbbrev"`
	NumberOfGames int    `json:"numberOfGames"`
	Games         []Game `json:"games"`
}

type Game struct {
	ID                int    `json:"id"`
	GameState         string `json:"gameState"`
	GameScheduleState string `json:"gameScheduleState"`
	GameDate          string `json:"gameDate"`
	AwayTeam          struct {
		PlaceName struct {
			Default string `json:"default"`
		} `json:"placeName"`
		Abbrev   string `json:"abbrev"`
		Logo     string `json:"logo"`
		DarkLogo string `json:"darkLogo"`
		Score    int    `json:"score"`
	} `json:"awayTeam"`
	HomeTeam struct {
		PlaceName struct {
			Default string `json:"default"`
		} `json:"placeName"`
		Abbrev   string `json:"abbrev"`
		Logo     string `json:"logo"`
		DarkLogo string `json:"darkLogo"`
		Score    int    `json:"score"`
	} `json:"homeTeam"`
	GameOutcome struct {
		PeriodType string `json:"lastPeriodType"`
	} `json:"gameOutcome"`
}

// Fetch all available games on specified date for specific team.
func fetchTeamWeekSchedule(team, matchDate string) {
	apiURL := fmt.Sprintf("%s/club-schedule/%s/week/%s", API_URL, team, matchDate)

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var data struct {
		Games []Game `json:"games"`
	}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(data.Games) == 0 {
		fmt.Println("No games found this week for ", team)
		return
	}

	games := data.Games
	fmt.Println("Schedule for ", team, " on ", matchDate)
	outputGamesInformation(games, true)
}

// Fetch all available games on specified date.
func fetchSchedule(matchDate string) {
	apiURL := fmt.Sprintf("%s/schedule/%s", API_URL, matchDate)

	response, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	var data struct {
		GameWeek []GameWeek `json:"gameWeek"`
	}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(data.GameWeek) == 0 || len(data.GameWeek[0].Games) == 0 {
		fmt.Println("No games found for ", matchDate)
		return
	}

	games := data.GameWeek[0].Games
	fmt.Println("Schedule for NHL games on ", matchDate)
	outputGamesInformation(games, false)
}

// Displays information for selected games on standard out.
func outputGamesInformation(games []Game, showDate bool) {
	for _, game := range games {
		if showDate {
			fmt.Printf("%s - ", game.GameDate)
		}
		fmt.Printf(
			"%s: %d vs %s: %d",
			game.AwayTeam.Abbrev,
			game.AwayTeam.Score,
			game.HomeTeam.Abbrev,
			game.HomeTeam.Score,
		)
		if game.GameOutcome.PeriodType == "OT" {
			fmt.Println(" (OT)")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.MatchDate, "date", time.Now().Format("2006-01-02"), "Gets schedule of games for specified date.")
	flag.StringVar(&cfg.TeamAbreviation, "team", "", "The abreviation name for a NHL team (e.g., NJD)")
	flag.Parse()

	if cfg.TeamAbreviation != "" {
		fetchTeamWeekSchedule(cfg.TeamAbreviation, cfg.MatchDate)
	} else {
		fetchSchedule(cfg.MatchDate)
	}
}
