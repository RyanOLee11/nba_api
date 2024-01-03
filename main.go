package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)	

func getResponse(requestURL string) []byte {
	req, _ := http.NewRequest("GET", requestURL, nil)

	req.Header.Add("X-RapidAPI-Key", "0fa6f2c001msha931157d6a00c9fp116347jsne09f513808e0")
	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
	}

	return respBody
}

func printTeam(team Team) {
	fmt.Println("City:", team.City)
	fmt.Println("Full Name:", team.FullName)
	fmt.Println("Team ID:", team.TeamID)
	fmt.Println("Nickname:", team.Nickname)
	fmt.Println("Logo:", team.Logo)
	fmt.Println("Short Name:", team.ShortName)
	fmt.Println("All Star:", team.AllStar)
	fmt.Println("NBA Franchise:", team.NbaFranchise)
}
func getTeams() {
	requestURL := "https://api-nba-v1.p.rapidapi.com/teams/league/standard"
	respBody := getResponse(requestURL)
	var teams TeamsJSON

	json.Unmarshal(respBody, &teams)

	for team := range teams.API.Teams {
		printTeam(teams.API.Teams[team])
		fmt.Println()
	}
}

func getPlayers() {
	requestURL := "https://api-nba-v1.p.rapidapi.com/players/teamId/1"
	respBody := getResponse(requestURL)
	var players PlayersJSON

	json.Unmarshal(respBody, &players)

	for player := range players.API.Players {
		printPlayer(players.API.Players[player])
		fmt.Println()
	}
}
func main() {
	requestURL := "https://api-nba-v1.p.rapidapi.com/playes/team/1"

	req, _ := http.NewRequest("GET", requestURL, nil)

	req.Header.Add("X-RapidAPI-Key", "0fa6f2c001msha931157d6a00c9fp116347jsne09f513808e0")
	req.Header.Add("X-RapidAPI-Host", "api-nba-v1.p.rapidapi.com")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
	}

	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error on response.\n[ERRO] -", err)
	}

	var teams TeamsJSON

	json.Unmarshal(respBody, &teams)

	for team := range teams.API.Teams {
		printTeam(teams.API.Teams[team])
		fmt.Println()
	}
		
}

type TeamsJSON struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Teams   []Team `json:"teams"`
	} `json:"api"`
}

type Team struct {
	City         string `json:"city"`
	FullName     string `json:"fullName"`
	TeamID       string `json:"teamId"`
	Nickname     string `json:"nickname"`
	Logo         string `json:"logo"`
	ShortName    string `json:"shortName"`
	AllStar      string `json:"allStar"`
	NbaFranchise string `json:"nbaFranchise"`
}


type PlayersJSON struct {
	API struct {
		Status  int      `json:"status"`
		Message string   `json:"message"`
		Results int      `json:"results"`
		Filters []string `json:"filters"`
		Players []Player `json:"players"`
	} `json:"api"`
}
type Player struct {
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	TeamID            string `json:"teamId"`
	YearsPro          string `json:"yearsPro"`
	CollegeName       string `json:"collegeName"`
	Country           string `json:"country"`
	PlayerID          string `json:"playerId"`
	DateOfBirth       string `json:"dateOfBirth"`
	Affiliation       string `json:"affiliation"`
	StartNba          string `json:"startNba"`
	HeightInMeters    string `json:"heightInMeters"`
	WeightInKilograms string `json:"weightInKilograms"`
	Leagues           struct {
		Standard struct {
			Jersey string `json:"jersey"`
			Active string `json:"active"`
			Pos    string `json:"pos"`
		} `json:"standard"`
	} `json:"leagues,omitempty"`
}