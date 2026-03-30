package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"
)

const fplBase = "https://fantasy.premierleague.com/api"

type Player struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	SecondName    string `json:"second_name"`
	TeamID        int    `json:"team"`
	TotalPoints   int    `json:"total_points"`
	NowCost       int    `json:"now_cost"`
	Form          string `json:"form"`
	SelectedByPct string `json:"selected_by_percent"`
	ElementType   int    `json:"element_type"` // 1=GK, 2=DEF, 3=MID, 4=FWD
}

type Team struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}

type Event struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	DeadlineTime string `json:"deadline_time"`
	IsNext       bool   `json:"is_next"`
	IsCurrent    bool   `json:"is_current"`
	IsFinished   bool   `json:"finished"`
}

type Bootstrap struct {
	Elements []Player `json:"elements"`
	Teams    []Team   `json:"teams"`
	Events   []Event  `json:"events"`
}

func main() {
	fmt.Println("Fetching FPL data...")

	bootstrap, err := fetchBootstrap()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	teams := make(map[int]Team)
	for _, t := range bootstrap.Teams {
		teams[t.ID] = t
	}

	for _, e := range bootstrap.Events {
		if e.IsCurrent {
			fmt.Printf("\nCurrent gameweek: %s\n", e.Name)
		}
		if e.IsNext {
			fmt.Printf("Next gameweek:    %s (deadline: %s)\n", e.Name, e.DeadlineTime)
		}
	}

	players := bootstrap.Elements
	sort.Slice(players, func(i, j int) bool {
		return players[i].TotalPoints > players[j].TotalPoints
	})

	fmt.Printf("\n%-25s %-20s %-8s %-8s %-10s\n",
		"Player", "Team", "Points", "Cost", "Selected%")
	fmt.Println("--------------------------------------------------------------------------------")

	for i, p := range players {
		if i >= 20 {
			break
		}
		team := teams[p.TeamID]
		fmt.Printf("%-25s %-20s %-8d £%-7.1fm %-10s%%\n",
			p.FirstName+" "+p.SecondName,
			team.Name,
			p.TotalPoints,
			float64(p.NowCost)/10,
			p.SelectedByPct,
		)
	}

	fmt.Printf("\nTotal players in dataset: %d\n", len(bootstrap.Elements))
	fmt.Printf("Total teams: %d\n", len(bootstrap.Teams))
}

func fetchBootstrap() (*Bootstrap, error) {
	client := &http.Client{Timeout: 30 * time.Second}

	resp, err := client.Get(fplBase + "/bootstrap-static/")
	if err != nil {
		return nil, fmt.Errorf("fetch bootstrap: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	var bootstrap Bootstrap
	if err := json.NewDecoder(resp.Body).Decode(&bootstrap); err != nil {
		return nil, fmt.Errorf("decode bootstrap: %w", err)
	}

	return &bootstrap, nil
}