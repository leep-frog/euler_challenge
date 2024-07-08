package y2016

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day16() aoc.Day {
	return &day16{}
}

type day16 struct{}

func (d *day16) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines[0], parse.Atoi(lines[1])), d.solve(lines[0], parse.Atoi(lines[2])))
}

func (d *day16) solve(dataStr string, length int) string {
	var data []bool
	for _, c := range dataStr {
		data = append(data, c == '1')
	}

	for len(data) < length {
		data = d.expand(data)
	}

	data = data[:length]
	return d.toString(d.checksum(data))
}

func (d *day16) toString(data []bool) string {
	var r []string
	for _, b := range data {
		if b {
			r = append(r, "1")
		} else {
			r = append(r, "0")
		}
	}
	return strings.Join(r, "")
}

func (d *day16) expand(data []bool) []bool {
	initLength := len(data)
	data = append(data, false)
	for i := initLength - 1; i >= 0; i-- {
		data = append(data, !data[i])
	}
	return data
}

func (d *day16) checksum(data []bool) []bool {
	if len(data)%2 == 1 {
		return data
	}

	var nd []bool
	for i := 0; i < len(data); i += 2 {
		nd = append(nd, data[i] == data[i+1])
	}
	return d.checksum(nd)
}

func (d *day16) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"01100 01100",
			},
		},
		{
			ExpectedOutput: []string{
				"10010100110011100 01100100101101100",
			},
		},
	}
}
