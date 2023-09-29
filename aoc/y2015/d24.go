package y2015

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func Day24() aoc.Day {
	return &day24{}
}

type day24 struct{}

func (d *day24) getGroupings(idx int, rem int, parts []int, cur []int, groupings *[][]int) {
	if rem < 0 {
		return
	}
	if rem == 0 {
		*groupings = append(*groupings, bread.Copy(cur))
		return
	}

	for i := idx; i < len(parts); i++ {
		d.getGroupings(i+1, rem-parts[i], parts, append(cur, parts[i]), groupings)
	}
}

func (d *day24) partOneCheck(opts map[int]bool, groupings [][]int, length int) bool {
	for _, group := range groupings {
		if functional.All(group, func(t int) bool { return opts[t] }) && len(group) >= length && len(opts)-len(group) >= length {
			return true
		}
	}
	return false
}

func (d *day24) partTwoCheck(opts map[int]bool, groupings [][]int, length int) bool {
	for _, group := range groupings {
		var subOpts map[int]bool
		if len(group) < length {
			continue
		}
		for _, g := range group {
			if !opts[g] {
				goto NOPE
			}
		}

		subOpts = maps.Clone(opts)
		for _, g := range group {
			delete(subOpts, g)
		}

		if d.partOneCheck(subOpts, groupings, length) {
			return true
		}

	NOPE:
		if functional.All(group, func(t int) bool { return opts[t] }) && len(group) >= length && len(opts)-len(group) >= length {
			return true
		}
	}
	return false
}

func (d *day24) Solve(lines []string, o command.Output) {
	fmt.Println("START")
	var parts []int
	var total int
	for _, line := range lines {
		v := parse.Atoi(line)
		total += v
		parts = append(parts, v)
	}

	o.Stdoutln(d.solve(parts, total, 3, d.partOneCheck), d.solve(parts, total, 4, d.partTwoCheck))
}

func (d *day24) solve(partsArr []int, total, segmentCount int, checker func(map[int]bool, [][]int, int) bool) int {
	size := total / segmentCount
	var groups [][]int
	d.getGroupings(0, size, partsArr, nil, &groups)

	slices.SortFunc(groups, func(this, that []int) bool {
		return bread.Product(this) < bread.Product(that)
	})

	for i, g := range groups {
		slices.Sort(g)
		groups[i] = g
	}

	for _, g := range groups {
		if len(g) > len(partsArr)/segmentCount {
			continue
		}
		var opts []int
		optsM := map[int]bool{}
		used := maths.NewSimpleSet(g...)
		for _, k := range partsArr {
			if !used[k] {
				opts = append(opts, k)
				optsM[k] = true
			}
		}
		if checker(optsM, groups, len(g)) {
			return bread.Product(g)
		}
	}
	return 0
}

func (d *day24) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"88 44",
			},
		},
		{
			ExpectedOutput: []string{
				"11846773891 80393059",
			},
		},
	}
}
