package y2015

import (
	"regexp"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
)

func Day05() aoc.Day {
	return &day05{}
}

type day05 struct{}

func (d *day05) containsDouble(line string) bool {
	for i, c := range line[1:] {
		if string(c) == line[i:i+1] {
			return true
		}
	}
	return false
}

func (d *day05) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.solve1(lines), d.solve2(lines))
}

func (d *day05) solve2(lines []string) int {
	var count int
	for _, line := range lines {
		pairCounts := map[string]int{}
		var rule1, rule2 bool

		for i, c := range line {
			if i > 1 && !rule1 {
				if string(c) == line[i-2:i-1] {
					rule1 = true
				}
			}

			if i > 0 && !rule2 {
				idx, ok := pairCounts[line[i-1:i+1]]
				if !ok {
					pairCounts[line[i-1:i+1]] = i
				} else if i > idx+1 {
					rule2 = true
				}
			}
		}
		if rule1 && rule2 {
			count++
		}
	}
	return count
}

func (d *day05) solve1(lines []string) int {
	regexes := []*regexp.Regexp{
		regexp.MustCompile(`[aeiou].*[aeiou].*[aeiou]`),
	}
	anitRegexes := []*regexp.Regexp{
		regexp.MustCompile(`ab`),
		regexp.MustCompile(`cd`),
		regexp.MustCompile(`pq`),
		regexp.MustCompile(`xy`),
	}
	var count int
	for _, line := range lines {
		match := functional.All(regexes, func(r *regexp.Regexp) bool { return r.MatchString(line) })
		antiMatch := functional.None(anitRegexes, func(r *regexp.Regexp) bool { return r.MatchString(line) })
		if match && antiMatch && d.containsDouble(line) {
			count++
		}
	}
	return count
}

func (d *day05) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"2 2",
			},
		},
		{
			ExpectedOutput: []string{
				"238 69",
			},
		},
	}
}
