package main

import (
	"bytes"
	"slices"
	"testing"
)

func TestMatchResult(t *testing.T) {
	team1 := Team{Name: "team1", PlayerNames: []string{"p1", "p2"}}
	team2 := Team{Name: "team2", PlayerNames: []string{"p3", "p4"}}

	testCases := []struct {
		name      string
		team1wins int
		team2wins int
	}{
		{
			name:      "team1 wins once",
			team1wins: 1,
			team2wins: 0,
		},
		{
			name:      "team2 wins once",
			team1wins: 0,
			team2wins: 1,
		},
		{
			name:      "team1 wins 3 times",
			team1wins: 3,
			team2wins: 0,
		},
		{
			name:      "team2 wins 3 times",
			team1wins: 0,
			team2wins: 3,
		},
		{
			name:      "both teams win multiple times",
			team1wins: 2,
			team2wins: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			league := League{Teams: []Team{team1, team2}}
			for range testCase.team1wins {
				league.MatchResult(team1.Name, 10, team2.Name, 5)
			}
			for range testCase.team2wins {
				league.MatchResult(team1.Name, 5, team2.Name, 10)
			}

			var (
				team1WantWins = testCase.team1wins
				team2WantWins = testCase.team2wins
				team1GotWins  = league.Wins[team1.Name]
				team2GotWins  = league.Wins[team2.Name]
			)

			if team1WantWins != team1GotWins {
				t.Errorf("team 1 want wins: %d, got wins: %d", team1WantWins, team1GotWins)
			}
			if team2WantWins != team2GotWins {
				t.Errorf("team 2 want wins: %d, got wins: %d", team2WantWins, team2GotWins)
			}
		})
	}
}

func TestRanking(t *testing.T) {
	team1 := Team{Name: "team1", PlayerNames: []string{"p1", "p2"}}
	team2 := Team{Name: "team2", PlayerNames: []string{"p3", "p4"}}
	league := League{Teams: []Team{team1, team2}}

	league.MatchResult(team1.Name, 10, team2.Name, 5)
	league.MatchResult(team1.Name, 10, team2.Name, 5)

	league.MatchResult(team1.Name, 10, team2.Name, 15)

	wantRank := []string{team1.Name, team2.Name}
	gotRank := league.Ranking()

	if !slices.Equal(wantRank, gotRank) {
		t.Errorf("expected: %v, got: %q", wantRank, gotRank)
	}
}

type TestRanker []string

func (t TestRanker) Ranking() []string {
	return t
}

func TestRankPrinter(t *testing.T) {
	var buf bytes.Buffer

	ranker := TestRanker([]string{"a", "b", "c"})

	RankPrinter(ranker, &buf)

	got := buf.String()
	want := "a\nb\nc"
	if want != got {
		t.Errorf("want: %q, got: %q", want, got)
	}
}
