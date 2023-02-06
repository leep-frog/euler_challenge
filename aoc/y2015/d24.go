package y2015

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

func Day24() aoc.Day {
	return &day24{}
}

type day24 struct{}

func (d *day24) runTwo(idx, rem int, parts, opts, groupA, groupB, groupC []int, best *maths.Bester[[][]int, int]) {
	if idx == len(opts) {
		if rem == 0 {
			arrangement := [][]int{groupA, groupB, groupC}
			// fmt.Println(arrangement)
			slices.SortFunc(arrangement, func(this, that []int) bool {
				return len(this) < len(that)
			})
			for z := 0; z < len(arrangement) && len(arrangement[z]) == len(arrangement[0]); z++ {
				best.IndexCheck(bread.Copy(arrangement), bread.Product(arrangement[z]))
			}
		}
		return
	}

	d.runTwo(idx+1, rem-parts[opts[idx]], parts, opts, groupA, append(groupB, parts[opts[idx]]), groupC, best)
	d.runTwo(idx+1, rem, parts, opts, groupA, groupB, append(groupC, parts[opts[idx]]), best)
}

func (d *day24) runOne(curIdx, total, sum int, parts, unusedIdx, group []int, best *maths.Bester[[][]int, int]) {
	if sum > total {
		return
	}
	if sum == total {
		ui := bread.Copy(unusedIdx)
		for j := curIdx; j < len(parts); j++ {
			ui = append(ui, j)
		}
		d.runTwo(0, total, parts, ui, group, nil, nil, best)

		return
	}

	if curIdx >= len(parts) {
		return
	}

	// Include current one in group
	d.runOne(curIdx+1, total, sum+parts[curIdx], parts, unusedIdx, append(group, parts[curIdx]), best)
	d.runOne(curIdx+1, total, sum, parts, append(unusedIdx, curIdx), group, best)
}

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

func (d *day24) evaluate(idx, length, rem int, parts, cur []int) bool {
	if rem == 0 {
		return len(cur) >= length && len(parts)-len(cur) >= length
	}

	if idx == len(parts) {
		return false
	}

	if rem < -0 {
		return false
	}

	return (d.evaluate(idx+1, length, rem-parts[idx], parts, append(cur, parts[idx])) ||
		d.evaluate(idx+1, length, rem, parts, cur))

}

func (d *day24) Solve(lines []string, o command.Output) {
	fmt.Println("START")
	parts := map[int]bool{}
	var partsArr []int
	var total int
	for _, line := range lines {
		v := parse.Atoi(line)
		total += v
		parts[v] = true
		partsArr = append(partsArr, v)
	}

	d.solve1(partsArr, total)
}

func (d *day24) solve1(partsArr []int, total int) {
	size := total / 3
	var groups [][]int
	fmt.Println("GETTING GROUPINGS", total, size)
	d.getGroupings(0, size, partsArr, nil, &groups)
	fmt.Println("Got'em")

	slices.SortFunc(groups, func(this, that []int) bool {
		return bread.Product(this) < bread.Product(that)
	})

	for i, g := range groups {
		slices.Sort(g)
		groups[i] = g
	}

	for _, g := range groups {
		_ = g
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
			fmt.Println("YUP", g, bread.Product(g))
			break
		}
	}
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
