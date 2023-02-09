package y2016

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) Solve(lines []string, o command.Output) {
	// var r []string
	r := make([]string, 8, 8)
	var cnt int
	doorID := lines[0]
	for j := 0; cnt < 8; j++ {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", doorID, j))
		if hashed := fmt.Sprintf("%x", h.Sum(nil)); strings.HasPrefix(hashed, "00000") {
			if idx, ok := parse.AtoiOK(hashed[5:6]); ok && idx >= 0 && idx < 8 && r[idx] == "" {
				// r = append(r, hashed[6:7])

				r[idx] = hashed[6:7]
				o.Stdoutln(j, hashed[6:7])
				cnt++
			}
			// break
		}
	}
	o.Stdoutln(strings.Join(r, ""))
}

func (d *day05) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"18f47a30 05ace8e3",
			},
		},
		{
			ExpectedOutput: []string{
				"801b56a7 	",
			},
		},
	}
}
