package y2020

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day02() aoc.Day {
	return &day02{}
}

type day02 struct{}

func (d *day02) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile("^([0-9]+)-([0-9]+) ([a-z]): ([a-z]*)$")
	var part1, part2 int
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		min, max, letter, pw := parse.Atoi(m[1]), parse.Atoi(m[2]), m[3], m[4]
		if count := strings.Count(pw, letter); count >= min && count <= max {
			part1++
		}
		left := pw[min-1:min] == letter
		right := pw[max-1:max] == letter
		if left != right {
			part2++
		}
	}
	o.Stdoutln(part1, part2)
}

func (d *day02) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"2 1",
			},
		},
		{
			ExpectedOutput: []string{
				"500 313",
			},
		},
	}
}
