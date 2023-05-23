package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/combinatorics"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day21() aoc.Day {
	return &day21{}
}

type day21 struct{}

func (d *day21) Solve(lines []string, o command.Output) {
	pwd := lines[0]
	part1 := d.solve(pwd, lines[2:])

	want := lines[1]
	for _, perm := range combinatorics.StringPermutations(strings.Split(pwd, "")) {
		if d.solve(perm, lines[2:]) == want {
			o.Stdoutln(part1, perm)
			return
		}
	}
}

func (d *day21) solve(pwd string, lines []string) string {

	password := linkedlist.NewCircularList(strings.Split(pwd, "")...)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "swap" && parts[1] == "position" {

			x, y := parse.Atoi(parts[2]), parse.Atoi(parts[5])
			a, b := password.Nth(x), password.Nth(y)
			a.Value, b.Value = b.Value, a.Value

		} else if parts[0] == "swap" && parts[1] == "letter" {
			a, _ := password.Index(parts[2])
			b, _ := password.Index(parts[5])
			a.Value, b.Value = b.Value, a.Value
		} else if parts[0] == "rotate" && parts[1] == "based" {
			_, cnt := password.Index(parts[6])
			for i := 0; i <= cnt; i++ {
				password = password.Prev
			}
			if cnt >= 4 {
				password = password.Prev
			}
		} else if parts[0] == "rotate" {
			times := parse.Atoi(parts[2])
			if parts[1] == "right" {
				times *= -1
			}
			password = password.Nth(times)
		} else if parts[0] == "reverse" {
			x, y := parse.Atoi(parts[2]), parse.Atoi(parts[4])
			if x > y {
				x, y = y, x
			}
			a, b := password.Nth(x), password.Nth(y)

			for i := x; i <= (x+y-1)/2; i++ {
				a.Value, b.Value = b.Value, a.Value
				a, b = a.Next, b.Prev
			}
		} else if parts[0] == "move" {
			from, to := parse.Atoi(parts[2]), parse.Atoi(parts[5])
			if from == to {
				continue
			} else if from > to {
				k := password.PopAt(from)
				pwd := password
				if to == 0 {
					pwd = k
				}
				password.PushAt(to, k)
				password = pwd
			} else { // from < to
				k := password
				pwd := password
				if from == 0 {
					pwd = password.Next
				} else {

				}
				k.PushAt(to, k.PopAt(from))
				password = pwd
			}
		} else {
			panic("UNKNOWN")
		}
	}
	return strings.Join(password.ToSlice(), "")
}

func (d *day21) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"decab ecabd",
			},
		},
		{
			ExpectedOutput: []string{
				"gcedfahb hegbdcfa",
			},
		},
	}
}
