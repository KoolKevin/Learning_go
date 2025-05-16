package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	TeamName    string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

/*
	Metodi:
		- attenzione al receiver!
		- qua ho deciso di mettere tutti value receiver dato che sto modificando la mappa che è un tipo riferimento (puntatore ai bucket)
		- se però avessi avuto bisogno di modificare anche l.Teams dentro a MatchResult, avrei dovuto usare un pointer receiver
			- altrimenti, ad esempio, eventuali append non sarebbero state visibili nel chiamante
*/

func (l League) MatchResult(team1Name string, team1Score int, team2Name string, team2Score int) {
	if team1Score > team2Score {
		l.Wins[team1Name]++
	} else if team2Score > team1Score {
		l.Wins[team2Name]++
	}
}

func (l League) RankingCopia() []string {
	copiaTeams := make([]Team, len(l.Teams))
	copy(copiaTeams, l.Teams)

	sort.Slice(copiaTeams, func(i, j int) bool {
		return l.Wins[copiaTeams[i].TeamName] > l.Wins[copiaTeams[j].TeamName]
	})

	ranking := make([]string, 0, len(l.Teams))
	for _, team := range copiaTeams {
		ranking = append(ranking, team.TeamName)
	}
	return ranking
}

func (l League) Ranking() []string {
	sort.Slice(l.Teams, func(i, j int) bool {
		return l.Wins[l.Teams[i].TeamName] > l.Wins[l.Teams[j].TeamName]
	})

	ranking := make([]string, 0, len(l.Teams))
	for _, team := range l.Teams {
		ranking = append(ranking, team.TeamName)
	}
	return ranking
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, rank := range r.Ranking() {
		_, err := io.WriteString(w, rank+"\n")
		if err != nil {
			fmt.Println("errore nello scrivere", rank, err)
		}
	}
}

func main() {
	teams := []Team{
		{
			TeamName:    "Cats",
			PlayerNames: []string{"kevin", "frank", "mario"},
		},
		{
			TeamName:    "Dogs",
			PlayerNames: []string{"alle", "adam", "simo"},
		},
		{
			TeamName:    "Eagles",
			PlayerNames: []string{"umberto"},
		},
	}

	wins := map[string]int{
		"Cats":   0,
		"Dogs":   0,
		"Eagles": 0,
	}

	league := League{
		Teams: teams,
		Wins:  wins,
	}

	fmt.Println("league prima dei match --------")
	fmt.Println(league)
	fmt.Println()

	league.MatchResult("Cats", 10, "Dogs", 11)
	league.MatchResult("Cats", 10, "Dogs", 9)
	league.MatchResult("Cats", 12, "Dogs", 11)
	league.MatchResult("Eagles", 12, "Dogs", 11)
	league.MatchResult("Cats", 12, "Eagles", 20)
	league.MatchResult("Eagles", 100, "Cats", 20)

	fmt.Println("league dopo i match --------")
	fmt.Println(league)
	fmt.Println()

	fmt.Println("ranking:", league.RankingCopia())
	fmt.Println("league dopo il ranking --------")
	fmt.Println(league)

	f, err := os.Create("rankfile.txt")
	if err != nil {
		fmt.Println("errore nell'aprire il file rankfile.txt")
	}
	defer f.Close()

	var w io.Writer = os.Stdout
	// var w io.Writer = f
	RankPrinter(league, w)
}
