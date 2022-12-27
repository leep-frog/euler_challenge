package y2022

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) toBounds(elf string) (int, int) {
	sides := strings.Split(elf, "-")
	return parse.Atoi(sides[0]), parse.Atoi(sides[1])
}

func (d *day04) Solve(lines []string, o command.Output) {
	var count1, count2 int
	for _, line := range lines {
		sides := strings.Split(line, ",")

		left1, right1 := d.toBounds(sides[0])
		left2, right2 := d.toBounds(sides[1])

		if left1 >= left2 && right1 <= right2 || left2 >= left1 && right2 <= right1 {
			count1++
		}

		// If either edge of 1 is between bounds of 2, or vice versa
		if (left1 >= left2 && left1 <= right2) || (right1 >= left2 && right1 <= right2) || (left2 >= left1 && left2 <= right1) || (right2 >= left1 && right2 <= right1) {
			count2++
		}
	}
	o.Stdoutln(count1, count2)
}

func (d *day04) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"2 4",
			},
		},
		{
			ExpectedOutput: []string{
				"503 827",
			},
		},
	}
}
