package y2015

import (
	"fmt"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"golang.org/x/exp/slices"
)

func Day11() aoc.Day {
	return &day11{}
}

type day11 struct{}

func (d *day11) valid(pwd []rune) bool {

	for i := 2; i < len(pwd); i++ {
		if pwd[i-2]+1 == pwd[i-1] && pwd[i-1]+1 == pwd[i] {
			goto VALID_1
		}
	}
	return false
VALID_1:

	var idxCount int
	for i := 1; i < len(pwd); i++ {
		if pwd[i-1] == pwd[i] {
			idxCount++
			i++
		}
	}
	if idxCount < 2 {
		return false
	}

	return !(slices.Contains(pwd, 'i'-'a') || slices.Contains(pwd, 'o'-'a') || slices.Contains(pwd, 'l'-'a'))
}

func (d *day11) convert(pwd []rune) string {
	var r []string
	for _, c := range pwd {
		r = append(r, fmt.Sprintf("%c", 'a'+c))
	}
	return strings.Join(r, "")
}

func (d *day11) increment(pwd []rune) {
	round := true
	for i := len(pwd) - 1; i >= 0; i-- {
		if round {
			pwd[i]++
			round = false
		}
		if pwd[i] >= 26 {
			pwd[i] = pwd[i] % 26
			round = true
		}
	}
}

func (d *day11) Solve(lines []string, o command.Output) {
	var pwd []rune
	for _, c := range lines[0] {
		pwd = append(pwd, c-'a')
	}

	for !d.valid(pwd) {
		d.increment(pwd)
	}
	part1 := d.convert(pwd)
	d.increment(pwd)
	for !d.valid(pwd) {
		d.increment(pwd)
	}

	o.Stdoutln(part1, d.convert(pwd))
}

func (d *day11) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"abbcffgh abbcfghh",
			},
		},
		{
			ExpectedOutput: []string{
				"hepxxyzz heqaabcc",
			},
		},
	}
}
