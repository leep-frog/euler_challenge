package y2020

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/bread"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
	"golang.org/x/exp/slices"
)

func Day19() aoc.Day {
	return &day19{}
}

type day19 struct{}

type rule struct {
	character string
	options   [][]int
}

func (d *day19) applyRules(depth int, line string, rules []*rule, lineIdx int, remainingRules *maths.Stack[int]) bool {
	if remainingRules.Len() == 0 {
		return lineIdx == len(line)
	}

	popped := remainingRules.Pop()
	rule := rules[popped]

	if rule.character != "" {
		if lineIdx < len(line) && line[lineIdx:lineIdx+1] == rule.character {
			r := d.applyRules(depth+1, line, rules, lineIdx+1, remainingRules)
			remainingRules.Push(popped)
			return r
		}
		remainingRules.Push(popped)
		return false
	}

	for _, group := range rule.options {
		for _, ri := range bread.Reverse(group) {
			remainingRules.Push(ri)
		}
		if d.applyRules(depth+1, line, rules, lineIdx, remainingRules) {
			return true
		}
		for i := range group {
			_ = i
			remainingRules.Pop()
		}
	}

	remainingRules.Push(popped)
	return false
}

func (d *day19) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve(lines, nil), d.solve(lines, map[string]string{
		"8: 42":     "8: 42 | 42 8",
		"11: 42 31": "11: 42 31 | 42 11 31",
	}))
}

func (d *day19) solve(lines []string, subs map[string]string) int {

	numRules := slices.Index(lines, "")
	rules := make([]*rule, numRules, numRules)

	r := regexp.MustCompile("^([0-9]+): (.*)$")

	rulesSection := true
	var cnt int
	for _, line := range lines {
		if v, ok := subs[line]; ok {
			line = v
		}

		if !rulesSection {
			if d.applyRules(0, line, rules, 0, maths.NewStack(0)) {
				cnt++
			}
			continue
		}

		if line == "" {
			rulesSection = false
			continue
		}

		m := r.FindStringSubmatch(line)
		if strings.HasPrefix(m[2], "\"") {
			rules[parse.Atoi(m[1])] = &rule{
				character: m[2][1:2],
			}
		} else {
			rules[parse.Atoi(m[1])] = &rule{
				options: functional.Map(strings.Split(m[2], " | "), func(s string) []int {
					return parse.AtoiArray(strings.Split(s, " "))
				}),
			}
		}
	}
	return cnt
}

func (d *day19) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"2 2",
			},
		},
		{
			ExpectedOutput: []string{
				"115 237",
			},
		},
	}
}
