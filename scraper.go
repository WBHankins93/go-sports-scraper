package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

// Struct to store team name and score information
type TeamScore struct {
	TeamName string `json:"team_name"`
	Score    string `json:"score"`
}

// Struct to represent a match between two teams
type Match struct {
	Date string    `json:"date"`
	Away TeamScore `json:"away"`
	Home TeamScore `json:"home"`
}

func main() {
	// Initialize the collector
	c := colly.NewCollector()

	// Slice to store scraped match data
	var results []Match

	// ESPN NFL scoreboard URL
	url := "https://www.espn.com/nfl/scoreboard"

	// Callback for when a "Scoreboard_Row" is found
	c.OnHTML(".ScoreboardScoreCell", func(e *colly.HTMLElement) {
		var awayTeam, homeTeam TeamScore

		// Extract the game date
		date := e.DOM.Find(".Card__Header__Title").Text()

		// Extract team names and scores
		teams := e.ChildTexts(".ScoreCell__TeamName")
		scores := e.ChildTexts(".ScoreCell__Score")

		// Handle missing scores
		if len(scores) < 2 {
			scores = []string{"-", "-"}
		}

		// Populate away and home teams
		if len(teams) == 2 {
			awayTeam.TeamName = teams[0]
			homeTeam.TeamName = teams[1]
		}

		// Populate scores if available
		if len(scores) == 2 {
			awayTeam.Score = scores[0]
			homeTeam.Score = scores[1]
		}

		// Append to results if valid date is found
		if awayTeam.TeamName != "" && homeTeam.TeamName != "" {
			results = append(results, Match{
				Date: date,
				Away: awayTeam,
				Home: homeTeam,
			})
		}
	})

	// OnRequest callback to log URL being visited
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Visit the target URL
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	// Marshal results into JSON
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Save JSON data to file
	file, err := os.Create("nfl_scores.json")
	if err != nil {
		log.Fatal("Could not create file:", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Could not write to file:", err)
	}

	fmt.Println("Scraping completed and saved to nfl_scores.json")
}