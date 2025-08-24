package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	N string
	W int
	D int
	L int
}

func (t Team) MP() int {
	return t.D + t.L + t.W
}

func (t Team) P() int {
	return t.W*3 + t.D
}

func (t Team) IsZero() bool {
	return t.D == 0 && t.L == 0 && t.W == 0
}

type Tournament map[string]Team

func (t Tournament) Merge(a Team) {
	if v, ok := t[a.N]; !ok {
		t[a.N] = a
	} else {
		t[a.N] = Team{
			N: a.N,
			W: v.W + a.W,
			D: v.D + a.D,
			L: v.L + a.L,
		}
	}
}

func (t Tournament) Game(game string) {
	g := strings.Split(game, ";")
	if len(g) != 3 {
		return
	}
	one := g[0]
	two := g[1]
	typ := g[2]

	switch typ {
	case "win":
		t.Merge(Team{one, 1, 0, 0})
		t.Merge(Team{two, 0, 0, 1})
	case "draw":
		t.Merge(Team{one, 0, 1, 0})
		t.Merge(Team{two, 0, 1, 0})
	case "loss":
		t.Merge(Team{one, 0, 0, 1})
		t.Merge(Team{two, 1, 0, 0})
	}
}

type TournamentTable []Team

func (table TournamentTable) Len() int {
	return len(table)
}

func (table TournamentTable) Less(i, j int) bool {

	return (table[i].P() > table[j].P()) || (table[i].P() == table[j].P() && table[i].N < table[j].N)
}

func (table TournamentTable) Swap(i, j int) {
	table[i], table[j] = table[j], table[i]
}

func (t Tournament) Table() (table TournamentTable) {
	for _, v := range t {
		if !v.IsZero() {
			table = append(table, v)
		}
	}
	sort.Sort(table)
	return table
}

func (table TournamentTable) Draw(w io.Writer) {
	io.WriteString(w, "Team                           | MP |  W |  D |  L |  P\n")

	for k, v := range table {
		io.WriteString(w, fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d", v.N, v.MP(), v.W, v.D, v.L, v.P()))
		if k < table.Len() {
			io.WriteString(w, "\n")
		}
	}
}

func Tally(reader io.Reader, writer io.Writer) error {
	table, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("not found reader")
	}
	var tournament Tournament = make(Tournament)
	for _, v := range strings.Split(string(table), "\n") {
		tournament.Game(v)
	}
	if len(tournament) == 0 {
		return fmt.Errorf("tournament is empty")
	}
	t := tournament.Table()
	t.Draw(writer)

	return nil
}
