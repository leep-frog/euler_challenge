package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/pair"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

func Day04() aoc.Day {
	return &day04{}
}

type day04 struct{}

func (d *day04) rotate(s string, times int) string {
	var r []string
	for _, c := range s {
		if c == '-' {
			r = append(r, " ")
		} else {
			letter := rune(((int(c-'a') + times) % 26) + 'a')
			r = append(r, string(letter))
		}
	}
	return strings.Join(r, "")
}

func (d *day04) Solve(lines []string, o command.Output) {

	var sum, part2 int
	for _, line := range lines {
		parts := strings.Split(line, "-")
		parts, end := parts[:len(parts)-1], parts[len(parts)-1]
		counts := map[rune]int{}
		for _, c := range strings.Join(parts, "") {
			counts[c]++
		}

		endParts := strings.Split(strings.TrimRight(end, "]"), "[")
		sectorID := parse.Atoi(endParts[0])

		elements := pair.Zip(counts)
		slices.SortFunc(elements, func(this, that *pair.Pair[rune, int]) bool {
			if this.B != that.B {
				return this.B > that.B
			}
			return this.A < that.A
		})

		code := strings.Join(functional.Map(elements[:5], func(p *pair.Pair[rune, int]) string { return string(p.A) }), "")
		if code == endParts[1] {
			sum += sectorID
		}

		if d.rotate(strings.Join(parts, "-"), sectorID) == "northpole object storage" {
			part2 = sectorID
		}
	}
	o.Stdoutln(sum, part2)
}

func (d *day04) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"1514 0",
			},
		},
		{
			ExpectedOutput: []string{
				"245102 324",
			},
		},
	}
}
