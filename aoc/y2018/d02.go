package y2018

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day02() aoc.Day {
	return &day02{}
}

type day02 struct{}

func (d *day02) Solve(lines []string, o command.Output) {
	var twos, threes int
	for _, line := range lines {
		m := map[rune]int{}
		for _, c := range line {
			m[c]++
		}
		var gotTwo, gotThree bool
		for _, v := range m {
			if v == 2 {
				gotTwo = true
			}
			if v == 3 {
				gotThree = true
			}
		}
		if gotTwo {
			twos++
		}
		if gotThree {
			threes++
		}
	}

	for i, a := range lines {
		for _, b := range lines[i+1:] {
			if len(a) != len(b) {
				continue
			}
			var same []string
			for j := 0; j < len(a); j++ {
				if a[j] == b[j] {
					same = append(same, a[j:j+1])
				}
			}
			if len(same) == len(a)-1 {
				o.Stdoutln(twos*threes, strings.Join(same, ""))
			}
		}
	}
}

func (d *day02) Cases() []*aoc.Case {
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
