package main

import (
	"fmt"
	"net/http"
	"offside/internal/fpl"
	"os"
	"sort"
	"time"
)

const fplBase = "https://fantasy.premierleague.com/api"


func main() {
	fmt.Println("Fetching FPL data...")

	client := &http.Client{Timeout: 30 * time.Second}

	bootstrap, err := fpl.FetchBootstrap(client, fplBase)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	teams := make(map[int]fpl.Team)
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