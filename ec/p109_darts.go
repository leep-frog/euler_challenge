package eulerchallenge

import (
	"fmt"
	"sort"
	"strings"

	"github.com/leep-frog/command/command"
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

		ctx := &dartContext{n, 0, 0}
		bfs.PoppableContextualDFS(allDarts, ctx)
		o.Stdoutln(ctx.count)
	}, []*execution{
		{
			args:     []string{"100"},
			want:     "38182",
			estimate: 0.5,
		},
		{
			args:     []string{"6"},
			want:     "11",
			estimate: 0.5,
		},
	})
}

type dartContext struct {
	n     int
	score int
	count int
}

type dart struct {
	score      int
	multiplier int
}

func (d *dart) OnPush(ctx *dartContext, dp bfs.DFSPath[*dart]) {
	ctx.score += d.score * d.multiplier
}

func (d *dart) OnPop(ctx *dartContext, dp bfs.DFSPath[*dart]) {
	ctx.score -= d.score * d.multiplier
}

func (d *dart) String() string {
	m := map[int]string{
		1: "S",
		2: "D",
		3: "T",
	}
	return fmt.Sprintf("%s%d", m[d.multiplier], d.score)
}

func (d *dart) Code(_ *dartContext, dp bfs.DFSPath[*dart]) string {
	darts := dp.Path()
	var r []string
	for _, dart := range darts {
		r = append(r, dart.String())
	}
	r = append(r, d.String())
	sort.Strings(r[:len(r)-1])
	return strings.Join(r, " ")
}

func (d *dart) Done(ctx *dartContext, dp bfs.DFSPath[*dart]) bool {
	if d.multiplier != 2 {
		return false
	}

	if ctx.score < ctx.n {
		ctx.count++
	}
	// return false because we want to explore all of them
	return false
}

func (d *dart) AdjacentStates(ctx *dartContext, dp bfs.DFSPath[*dart]) []*dart {
	darts := dp.Path()
	if len(darts) >= 3 {
		return nil
	}

	var ns []*dart
	for _, dt := range allDarts {
		ns = append(ns, &dart{dt.score, dt.multiplier})
	}
	return ns
}
