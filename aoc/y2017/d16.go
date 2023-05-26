package y2017

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/functional"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day16() aoc.Day {
	return &day16{}
}

type day16 struct{}

func (d *day16) Solve(lines []string, o command.Output) {
	letterMap := map[string]int{}
	revLetterMap := map[int]string{}
	numLetters := 16
	for i := 0; i < numLetters; i++ {
		letterMap[string('a'+i)] = i
		revLetterMap[i] = string('a' + i)
	}

	var dances []func(*linkedlist.Node[int]) *linkedlist.Node[int]
	for _, code := range strings.Split(lines[0], ",") {
		switch code[0] {
		case 's':
			num := parse.Atoi(code[1:])
			dances = append(dances, func(l *linkedlist.Node[int]) *linkedlist.Node[int] {
				return l.Nth(16 - num)
			})
		case 'x':
			parts := strings.Split(code[1:], "/")
			dances = append(dances, func(l *linkedlist.Node[int]) *linkedlist.Node[int] {
				a, b := l.Nth(parse.Atoi(parts[0])), l.Nth(parse.Atoi(parts[1]))
				a.Value, b.Value = b.Value, a.Value
				return l
			})
		case 'p':
			parts := strings.Split(code[1:], "/")
			dances = append(dances, func(l *linkedlist.Node[int]) *linkedlist.Node[int] {
				a, _ := l.Index(letterMap[parts[0]])
				b, _ := l.Index(letterMap[parts[1]])
				a.Value, b.Value = b.Value, a.Value
				return l
			})
		}
	}

	root := linkedlist.CircularNumbered(numLetters)
	seen := map[string]int{}
	upTo := 1_000_000_000
	var part1 string

	for i := 0; i < upTo; i++ {
		code := linkedlist.CircularRepresentation(root)
		if _, ok := seen[code]; ok {
			upTo = i + (upTo % i)
		}
		seen[code] = i

		for _, dance := range dances {
			root = dance(root)
		}
		if i == 0 {
			part1 = strings.Join(functional.Map(root.ToSlice(), func(k int) string { return revLetterMap[k] }), "")
		}
	}
	o.Stdoutln(part1, strings.Join(functional.Map(root.ToSlice(), func(k int) string { return revLetterMap[k] }), ""))
}

func (d *day16) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"jcobhadfnmpkglie pclhmengojfdkaib",
			},
		},
	}
}
