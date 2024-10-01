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
	c.OnHTML(".Scoreboard_Row", func(e *colly.HTMLElement) {
		var awayTeam, homeTeam TeamScore

		// Extract the away team name and score
		e.ForEach(".ScoreCell__Item--away", func(_ int, el *colly.HTMLElement) {
			awayTeam.TeamName = el.ChildText(".ScoreCell__Team")
			awayTeam.Score = el.ChildText(".ScoreCell__Score") // Placeholder for the actual score class
		})

		// Extract the home team name and score
		e.ForEach(".ScoreCell__Item--home", func(_ int, el *colly.HTMLElement) {
			homeTeam.TeamName = el.ChildText(".ScoreCell__Team")
			homeTeam.Score = el.ChildText(".ScoreCell__Score") // Placeholder for the actual score class
		})

		// Append to results
		results = append(results, Match{
			Away: awayTeam,
			Home: homeTeam,
		})
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