package y2020

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths/commandths"
)

func Day18() aoc.Day {
	return &day18{}
}

type day18 struct{}

func (d *day18) Solve(lines []string, o command.Output) {
	var sum1, sum2 int
	ops1 := []commandths.Operation[int]{&Plus{1}, &Times{}}
	ops2 := []commandths.Operation[int]{&Plus{0}, &Times{}}
	for _, line := range lines {
		line = strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(line, ")", " ) "), "(", " ( "))
		v1, _ := commandths.Parse(line, ops1...)
		v2, _ := commandths.Parse(line, ops2...)
		sum1 += v1
		sum2 += v2
	}
	o.Stdoutln(sum1, sum2)
}

type Plus struct {
	priority commandths.PemdasPriority
}

func (*Plus) Symbols() []string                           { return []string{"+"} }
func (*Plus) Evaluate(a, b int) int                       { return a + b }
func (p *Plus) PemdasPriority() commandths.PemdasPriority { return p.priority }

type Times struct{}

func (*Times) Symbols() []string                         { return []string{"*"} }
func (*Times) Evaluate(a, b int) int                     { return a * b }
func (*Times) PemdasPriority() commandths.PemdasPriority { return 1 }

func (d *day18) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"26457 694173",
			},
		},
		{
			ExpectedOutput: []string{
				"8929569623593 231235959382961",
			},
		},
	}
}
