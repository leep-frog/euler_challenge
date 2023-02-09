package y2015

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day16() aoc.Day {
	return &day16{}
}

type day16 struct{}

func (d *day16) Solve(lines []string, o command.Output) {
	r := regexp.MustCompile("^Sue (-?[0-9]+): (.*)$")
	var sues []map[string]int
	for _, line := range lines {
		m := r.FindStringSubmatch(line)
		sue := map[string]int{}
		for _, pair := range strings.Split(m[2], ", ") {
			parts := strings.Split(pair, ": ")
			sue[parts[0]] = parse.Atoi(parts[1])
		}
		sues = append(sues, sue)
	}

	requirements := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	var part1, part2 int
	for idx, sue := range sues {
		valid1, valid2 := true, true
		for k, v := range sue {
			// Part 1
			if requirements[k] != v {
				valid1 = false
			}

			// Part 2
			if k == "cats" || k == "trees" {
				if v <= requirements[k] {
					valid2 = false
				}
			} else if k == "pomeranians" || k == "goldfish" {
				if v >= requirements[k] {
					valid2 = false
				}
			} else if requirements[k] != v {
				valid2 = false
			}

		}
		if valid1 {
			part1 = idx + 1
		}
		if valid2 {
			part2 = idx + 1
		}
	}
	o.Stdoutln(part1, part2)
}

func (d *day16) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			ExpectedOutput: []string{
				"40 241",
			},
		},
	}
}
