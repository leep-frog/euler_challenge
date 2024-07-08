package y2015

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines[0], 5), d.solve(lines[0], 6))
}

func (d *day04) solve(code string, numZeroes int) int {
	prefix := strings.Repeat("0", numZeroes)
	for i := 1; ; i++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", code, i))
		if hashed := fmt.Sprintf("%x", h.Sum(nil)); strings.HasPrefix(hashed, prefix) {
			return i
		}
	}
}

func (d *day04) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"609043 6742839",
			},
		},
		{
			ExpectedOutput: []string{
				"282749 9962624",
			},
		},
	}
}
