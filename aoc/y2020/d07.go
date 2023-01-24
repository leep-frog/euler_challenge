package y2020

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day07() aoc.Day {
	return &day07{}
}

type day07 struct{}

func (d *day07) explore(bag string, paths map[string][]string, checked map[string]bool) {
	if checked[bag] {
		return
	}
	checked[bag] = true
	for _, parentBag := range paths[bag] {
		d.explore(parentBag, paths, checked)
	}
}

func (d *day07) explore2(bag string, paths map[string][]*bagCount, checked map[string]int) int {
	if v, ok := checked[bag]; ok {
		return v
	}

	var sum int
	for _, bc := range paths[bag] {
		sum += bc.count * (1 + d.explore2(bc.bag, paths, checked))
	}
	checked[bag] = sum
	return sum
}

type bagCount struct {
	bag   string
	count int
}

func (bc *bagCount) String() string {
	return fmt.Sprintf("[%d %v]", bc.count, bc.bag)
}

func (d *day07) Solve(lines []string, o command.Output) {
	paths := map[string][]string{}
	revPaths := map[string][]*bagCount{}
	r := regexp.MustCompile(`^([a-z ]*) bags contain(.*)\.$`)
	sr := regexp.MustCompile(`^ ([0-9]+) ([a-z ]*) bags?$`)
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		bag := m[1]
		if m[2] == " no other bags" {
			continue
		}
		for _, part := range strings.Split(m[2], ",") {
			sm := sr.FindStringSubmatch(part)
			cnt := parse.Atoi(sm[1])
			subbag := sm[2]

			paths[subbag] = append(paths[subbag], bag)
			revPaths[bag] = append(revPaths[bag], &bagCount{subbag, cnt})
		}
	}

	containable := map[string]bool{}
	d.explore("shiny gold", paths, containable)
	o.Stdoutln(len(containable)-1, d.explore2("shiny gold", revPaths, map[string]int{}))
}

func (d *day07) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"4 32",
			},
		},
		{
			ExpectedOutput: []string{
				"101 108636",
			},
		},
	}
}
