package y2015

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
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

func (d *day24) evaluate2(opts map[int]bool, groupings [][]int, length int) bool {
	for _, group := range groupings {
		if functional.All(group, func(t int) bool { return opts[t] }) && len(group) >= length && len(opts)-len(group) >= length {
			return true
		}
	}
	return false
}

func (d *day24) evaluate3(opts map[int]bool, groupings [][]int, length int) bool {
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

		if d.evaluate2(subOpts, groupings, length) {
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

	o.Stdoutln(d.solve1(parts, total), d.solve2(parts, total))
}

func (d *day24) solve2(partsArr []int, total int) int {
	size := total / 4
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
		if len(g) > len(partsArr)/4 {
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
		if d.evaluate3(optsM, groups, len(g)) {
			return bread.Product(g)
		}
	}
	return 0
}

func (d *day24) solve1(partsArr []int, total int) int {
	size := total / 3
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
		if len(g) > len(partsArr)/3 {
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
		if d.evaluate2(optsM, groups, len(g)) {
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
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
