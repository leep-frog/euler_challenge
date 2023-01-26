package y2015

import (
	"fmt"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day17() aoc.Day {
	return &day17{}
}

type day17 struct{}

func (d *day17) rec(rem, curIdx int, containers []int, cache map[string]map[int]int, totalContainers int) map[int]int {
	code := fmt.Sprintf("%d %d", rem, curIdx)
	if curIdx >= len(containers) {
		if rem == 0 {
			r := map[int]int{totalContainers: 1}
			cache[code] = r
			return r
		}
		return nil
	}

	k := map[int]int{}
	for numContainers := 0; numContainers <= 1 && numContainers*containers[curIdx] <= rem; numContainers++ {
		for a, b := range d.rec(rem-numContainers*containers[curIdx], curIdx+1, containers, cache, totalContainers+numContainers) {
			k[a] += b
		}
	}
	cache[code] = k
	return k
}

func (d *day17) Solve(lines []string, o command.Output) {
	size := parse.Atoi(lines[0])
	var containers []int
	for _, line := range lines[1:] {
		containers = append(containers, parse.Atoi(line))
	}

	k := d.rec(size, 0, containers, map[string]map[int]int{}, 0)
	o.Stdoutln(bread.Sum(maps.Values(k)), k[maths.Min(maps.Keys(k)...)])
}

func (d *day17) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 3",
			},
		},
		{
			ExpectedOutput: []string{
				"4372",
			},
		},
	}
}
