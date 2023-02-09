package y2022

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) parseMove(move string) (int, int, int) {
	r := regexp.MustCompile("^move ([0-9]+) from ([0-9]+) to ([0-9]+)$")
	parts := r.FindStringSubmatch(move)
	return parse.Atoi(parts[1]), parse.Atoi(parts[2]), parse.Atoi(parts[3])
}

func (d *day05) Solve(lines []string, o command.Output) {
	d.solve(slices.Clone(lines), o, true)
	d.solve(slices.Clone(lines), o, false)
}

func (d *day05) solve(lines []string, o command.Output, part1 bool) {
	var stacks [][]string
	i := 1
	for line := lines[0]; line[1] != '1'; i, line = i+1, lines[i] {
		for j := 1; j < len(line); j += 4 {
			stack := (j - 1) / 4
			for stack >= len(stacks) {
				stacks = append(stacks, []string{})
			}

			if s := line[j : j+1]; s != " " {
				stacks[stack] = append(stacks[stack], s)
			}
		}
	}

	for i, stack := range stacks {
		stacks[i] = bread.Reverse(stack)
	}

	for _, move := range lines[i+1:] {
		count, from, to := d.parseMove(move)
		from--
		to--
		if part1 {
			stacks[to] = append(stacks[to], bread.Reverse(stacks[from][len(stacks[from])-count:])...)
		} else {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-count:]...)
		}
		stacks[from] = stacks[from][:len(stacks[from])-count]
	}

	o.Stdoutln(strings.Join(functional.Map(stacks, func(stack []string) string { return stack[len(stack)-1] }), ""))
}

func (d *day05) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"CMZ",
				"MCD",
			},
		},
		{
			ExpectedOutput: []string{
				"ZWHVFWQWW",
				"HZFZCCWWV",
			},
		},
	}
}
