package y2018

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/euler_challenge/rgx"
	"golang.org/x/exp/slices"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

func (d *day07) Solve(lines []string, o command.Output) {
	numWorkers := parse.Atoi(lines[0])
	secondsBuffer := parse.Atoi(lines[1])
	lines = lines[2:]
	part1, _ := d.solve(lines, 1, 0)
	_, part2 := d.solve(lines, numWorkers, secondsBuffer)
	o.Stdoutln(part1, part2)

}

func (d *day07) solve(lines []string, numWorkers, secondsBuffer int) (string, int) {
	r := rgx.New(`^Step ([A-Za-z]+) must be finished before step ([A-Za-z]+) can begin.$`)
	forward, backward := map[string]map[string]bool{}, map[string]map[string]bool{}
	possible := map[string]int{}
	for _, line := range lines {
		match := r.MustMatch(line)
		maths.Insert(forward, match[1], match[0], true)
		maths.Insert(backward, match[0], match[1], true)

		possible[match[1]]++
		if _, ok := possible[match[0]]; !ok {
			possible[match[0]] = 0
		}

	}

	processed := map[string]bool{}
	var available []string
	for k, v := range possible {
		if v == 0 {
			available = append(available, k)
		}
	}

	var order []string

	// doneAt is a map from seconds to a list of letters that are done at that second
	doneAt := map[int][]string{}
	var seconds int
	for ; len(processed) < len(possible); seconds++ {
		// First free up any workers
		for _, completed := range doneAt[seconds] {
			processed[completed] = true
			order = append(order, completed)
			for dep := range backward[completed] {
				delete(forward[dep], completed)
				if len(forward[dep]) == 0 {
					available = append(available, dep)
				}
			}
		}
		delete(doneAt, seconds)
		slices.Sort(available)

		for len(available) > 0 {
			if len(doneAt) >= numWorkers {
				break
			}

			letter := available[0]
			available = available[1:]
			finishTime := seconds + secondsBuffer + int((1 + letter[0] - 'A'))
			doneAt[finishTime] = append(doneAt[finishTime], letter)
		}
	}

	return strings.Join(order, ""), seconds - 1
}

func (d *day07) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"CABDFE 15",
			},
		},
		{
			ExpectedOutput: []string{
				"IJLFUVDACEHGRZPNKQWSBTMXOY 1072",
			},
		},
	}
}
