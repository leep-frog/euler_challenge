package y2015

import (
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/maps"
)

func Day13() aoc.Day {
	return &day13{}
}

type day13 struct{}

func (d *day13) rec(startingPerson, currentPerson string, utils map[string]map[string]int, people map[string]bool, total int, best *maths.Bester[int, int]) {
	if len(people) == 1 {
		best.Check(total + utils[startingPerson][currentPerson] + utils[currentPerson][startingPerson])
		return
	}

	delete(people, currentPerson)
	for _, nextPerson := range maps.Keys(people) {
		d.rec(startingPerson, nextPerson, utils, people, total+utils[currentPerson][nextPerson]+utils[nextPerson][currentPerson], best)
	}
	people[currentPerson] = true
}

func (d *day13) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, false), d.solve(lines, true))
}

func (d *day13) solve(lines []string, part2 bool) int {
	utils := map[string]map[string]int{}
	people := map[string]bool{}
	r := regexp.MustCompile(`^(.*) would (.*) ([0-9]+) happiness units by sitting next to (.*)\.$`)
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		a, u, b := m[1], parse.Atoi(m[3]), m[4]
		if m[2] == "lose" {
			u = -u
		}
		maths.Insert(utils, a, b, u)
		people[a] = true
	}

	if part2 {
		people["you"] = true
		utils["you"] = map[string]int{}
	}

	best := maths.Largest[int, int]()
	delete(people, "Alice")
	d.rec("Alice", "Alice", utils, people, 0, best)
	return best.Best()
}

func (d *day13) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"330 286",
			},
		},
		{
			ExpectedOutput: []string{
				"709 668",
			},
		},
	}
}
