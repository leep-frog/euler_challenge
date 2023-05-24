package y2017

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
	"github.com/leep-frog/euler_challenge/topology"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

type tower struct {
	name      string
	weight    int
	subtowers []string
}

type towerContext struct {
	root  string
	part1 bool
}

func (t *tower) Code() string {
	return t.name
}

func (t *tower) Dependencies() []string {
	return t.subtowers
}

func (t *tower) solve(dg *topology.DependencyGraph[string, *tower]) (int, int) {
	deps := dg.Dependencies[t.name]
	if len(deps) == 0 {
		return t.weight, -1
	}

	type depWht struct {
		depWht   int
		totalWht int
	}

	depWhts := map[string]*depWht{}
	weights := map[int]int{}
	total := t.weight
	for _, dep := range dg.Dependencies[t.name] {
		depT := dg.Nodes[dep]
		k, s := depT.solve(dg)
		if s != -1 {
			return 0, s
		}
		weights[k]++
		depWhts[dep] = &depWht{depT.weight, k}
		total += k
	}

	if len(weights) > 1 {
		// Pick the odd one out
		var need int
		for w, cnt := range weights {
			if cnt > 1 {
				need = w
			}
		}

		for _, dw := range depWhts {
			if dw.totalWht != need {
				return 0, need - dw.totalWht + dw.depWht
			}
		}
	}
	return total, -1
}

func (d *day07) Solve(lines []string, o command.Output) {
	r := rgx.New(`^([a-z]+) \(([0-9]+)\) ?(.*)$`)
	towers := functional.Map(lines, func(line string) *tower {
		m := r.MustMatch(line)
		var deps []string
		if m[2] != "" {
			deps = strings.Split(strings.TrimPrefix(m[2], "-> "), ", ")
		}
		return &tower{m[0], parse.Atoi(m[1]), deps}
	})

	dg := topology.GetDependencyGraph[string](towers)
	root := dg.InvertedRoots[0]
	_, part2 := root.solve(dg)

	o.Stdoutln(root.name, part2)
}

func (d *day07) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"tknk 60",
			},
		},
		{
			ExpectedOutput: []string{
				"ykpsek 1060",
			},
		},
	}
}
