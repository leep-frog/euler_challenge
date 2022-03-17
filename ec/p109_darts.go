package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/bfs"
)

var (
	allDarts = []*dart{}
)

func P109() *problem {
	return intInputNode(109, func(o command.Output, n int) {
		for i := 1; i <= 20; i++ {
			allDarts = append(allDarts,
				&dart{i, 1},
				&dart{i, 2},
				&dart{i, 3},
			)
		}
		allDarts = append(allDarts, &dart{25, 1}, &dart{25, 2})

		initStates := []*dartRound{{nil, n}}
		count := 0
		bfs.DFS(initStates, &count)
		o.Stdoutln(count)
	})
}

type dart struct {
	score      int
	multiplier int
}

func (d *dart) String() string {
	m := map[int]string{
		1: "S",
		2: "D",
		3: "T",
	}
	return fmt.Sprintf("%s%d", m[d.multiplier], d.score)
}

func (d *dart) getScore() int {
	return d.score*d.multiplier
}

type dartRound struct {
	darts []*dart
	n int
}

func (dr *dartRound) Code(*bfs.Context[*int, *dartRound]) string {
	var r []string
	if len(dr.darts) > 0 {
		sort.SliceStable(dr.darts[:len(dr.darts)-1], func(i, j int) bool { return dr.darts[i].String() < dr.darts[j].String()})
	}
	for _, d := range dr.darts {
		r = append(r, d.String())
	}
	return strings.Join(r, " ")
}

func (dr *dartRound) Done(ctx *bfs.Context[*int, *dartRound]) bool {
	if len(dr.darts) == 0 || len(dr.darts) > 3 {
		return false
	}

	if dr.darts[len(dr.darts)-1].multiplier != 2 {
		return false
	}

	var score int
	for _, d := range dr.darts {
		score += d.score*d.multiplier
	}

	if score >= dr.n {
		return false
	}

	*ctx.GlobalContext += 1
	// return false because we want to explore all of them
	return false
}

func (dr *dartRound) AdjacentStates(ctx *bfs.Context[*int, *dartRound]) []*dartRound {
	if len(dr.darts) >= 3 {
		return nil
	}

	var ns []*dartRound
	for _, d := range allDarts {
		ns = append(ns, &dartRound{
			append(dr.darts, d),
			dr.n,
		})
	}
	return ns
}