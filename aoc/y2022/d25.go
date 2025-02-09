package y2022

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
)

func Day25() aoc.Day {
	return &day25{}
}

type day25 struct{}

var (
	snafuMap = map[string]int{
		"-": 2,
		"=": 2,
	}
)

func (d *day25) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.toSnafu(bread.Sum(functional.Map(lines, d.fromSnafu))))
}

func (d *day25) fromSnafu(k string) int {
	v := 1
	var sum int
	for _, c := range bread.Reverse(strings.Split(k, "")) {
		switch c {
		case "-":
			sum -= v
		case "=":
			sum -= 2 * v
		default:
			sum += v * parse.Atoi(c)
		}
		v *= 5
	}
	return sum
}

func (d *day25) toSnafu(k int) string {
	var r []string
	for ; k > 0; k = (k + 2) / 5 {
		m := (k % 5)
		switch m {
		case 3:
			r = append(r, "=")
		case 4:
			r = append(r, "-")
		default:
			r = append(r, parse.Itos(m))
		}
	}
	return strings.Join(bread.Reverse(r), "")
}

func (d *day25) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"2=-1=0",
			},
		},
		{
			ExpectedOutput: []string{
				"2=020-===0-1===2=020",
			},
		},
	}
}
