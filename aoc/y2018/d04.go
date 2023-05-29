package y2018

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
	"golang.org/x/exp/slices"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) Solve(lines []string, o command.Output) {
	slices.Sort(lines)
	var guard int
	r := rgx.New(`^\[(.*) ([0-9]+):([0-9]+)\] (.*)$`)
	shift := rgx.New(`^Guard #([0-9]+) begins shift$`)
	// map from guard to day to minutes asleep
	shifts := map[int]map[string][]int{}
	for _, line := range lines {
		m := r.MustMatch(line)
		action := m[3]
		switch action {
		case "falls asleep", "wakes up":
			if shifts[guard] == nil {
				shifts[guard] = map[string][]int{}
			}
			shifts[guard][m[0]] = append(shifts[guard][m[0]], parse.Atoi(m[2]))
		default:
			guard = shift.MatchInts(action)[0]
		}
	}

	// Map from guard to minute asleep to number of times
	minutesAsleep := map[int]map[int]int{}
	best := maths.Largest[int, int]()
	for g, m := range shifts {
		var sum int
		minutesAsleep[g] = map[int]int{}
		for _, changes := range m {
			for i := 0; i < len(changes); i += 2 {
				sum += changes[i+1] - changes[i]
				for j := changes[i]; j < changes[i+1]; j++ {
					minutesAsleep[g][j]++
				}
			}
		}
		best.IndexCheck(g, sum)
	}

	sleepy := shifts[best.BestIndex()]
	minuteCounts := map[int]int{}
	for _, changes := range sleepy {
		for i := 0; i < len(changes); i += 2 {
			for j := changes[i]; j < changes[i+1]; j++ {
				minuteCounts[j]++
			}
		}
	}
	minuteBest := maths.Largest[int, int]()
	for min, v := range minuteCounts {
		minuteBest.IndexCheck(min, v)
	}

	// Part 2
	part2 := maths.Largest[int, int]()
	for g, m := range minutesAsleep {
		for min, v := range m {
			part2.IndexCheck(g*min, v)
		}
	}
	o.Stdoutln(best.BestIndex()*minuteBest.BestIndex(), part2.BestIndex())
}

func (d *day04) Cases() []*aoc.Case {
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
