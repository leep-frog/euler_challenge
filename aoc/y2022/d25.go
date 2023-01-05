package y2022

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
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
	o.Stdoutln(d.toSnafu(maths.SumSys(parse.Map(lines, d.fromSnafu)...)))
	// fmt.Println(d.toSnafu(sum))
}

func (d *day25) fromSnafu(k string) int {
	v := 1
	var sum int
	for _, c := range maths.Reverse(strings.Split(k, "")) {
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
	// fmt.Println(k, sum, d.toSnafu(sum))
	return sum
}

func (d *day25) toSnafu(k int) string {
	// v := 5
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
	return strings.Join(maths.Reverse(r), "")
}

func (d *day25) Cases() []*aoc.Case {
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
