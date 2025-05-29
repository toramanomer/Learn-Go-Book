package main

import (
	"io"
	"slices"
	"strings"
)

type Team struct {
	Name        string
	PlayerNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if l.Wins == nil {
		l.Wins = map[string]int{}
	}

	if score1 > score2 {
		l.Wins[team1] += 1
	}
	if score1 < score2 {
		l.Wins[team2] += 1
	}
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(ranker Ranker, w io.Writer) {
	io.WriteString(w, strings.Join(ranker.Ranking(), "\n"))
}

func (l *League) Ranking() []string {
	type teamwin struct {
		name string
		wins int
	}
	t := make([]teamwin, 0, len(l.Wins))
	for teamName, wins := range l.Wins {
		t = append(t, teamwin{name: teamName, wins: wins})

	}
	slices.SortStableFunc(t, func(a, b teamwin) int {
		return b.wins - a.wins
	})

	ranks := make([]string, len(t))
	for i, teamwin := range t {
		ranks[i] = teamwin.name
	}

	return ranks
}

func main() {
}
