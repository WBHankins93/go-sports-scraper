# NFL Score Scraper

This project is a simple web scraper built with Go (`golang`) that scrapes NFL game data from ESPN's scoreboard page. It extracts information such as team names, scores, and game dates and stores the data in a JSON file.

## Features

- **Scrape Team Names and Scores**: Get the names of NFL teams and their respective scores from the ESPN scoreboard.
- **Save Data to JSON**: The scraped data is saved in a `nfl_scores.json` file for easy access and further processing.
- **Date-Specific Scraping**: The scraper currently targets a specific week, year, and season type.

## Project Structure

- **`scraper.go`**: The main Go file that contains all the scraping logic using the `colly` library.

## Getting Started

### Prerequisites

- **Go (Golang)**: Make sure you have Go installed on your system. You can download it [here](https://golang.org/dl/).
- **Colly Library**: The scraper uses the `colly` package. 
Install it using:
    ```
    go get github.com/gocolly/colly/v2
    ```

### Installation

1. **Clone the Repository**
   ```
   git clone https://github.com/WBHankins93/nfl-score-scraper.git
   cd nfl-score-scraper 
   ```

2. **Build the Project**

If you want to build the project and create an executable, run:
```
go build scraper.go
```

3. **Run the Scraper**

You can run the scraper with:
```
go run scraper.go
```
This will scrape the NFL scoreboard for the specified week, year, and season type and save the data in nfl_scores.json.

Future Plans
1. **Add Database Support**
Currently, the scraper stores data in a JSON file. The plan is to add database support (e.g., SQLite, PostgreSQL) to allow for better data storage, retrieval, and querying.
2. **Specify Date Range for Scraping**
Allow the user to specify a custom date or week range to scrape scores for multiple weeks or entire seasons, not just a hardcoded week/year.
3. **Include Team Record Information**
Scrape and store additional data for each team, such as their win-loss record, to provide more context about team performance.
4. **Scrape Team Schedule**
Add functionality to scrape each team's schedule (upcoming and past games), providing a more comprehensive view of each team's season.