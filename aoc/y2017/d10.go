package y2017

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/linkedlist"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
	"github.com/leep-frog/functional"
)

func Day10() aoc.Day {
	return &day10{}
}

type day10 struct{}

func (d *day10) Solve(lines []string, o command.Output) {
	root := linkedlist.CircularNumbered(parse.Atoi(lines[0]) + 1)
	fixed := root
	for skipSize, length := range parse.AtoiArray(strings.Split(lines[1], ",")) {
		root.ReverseUpTo(length - 1)
		root = root.Nth(skipSize + length)
		skipSize++
	}

	// part2
	input := lines[1]
	if len(lines) > 2 {
		input = lines[2]
	}

	part1 := fixed.Value * fixed.Next.Value
	o.Stdoutln(part1, knotHash(input))
}

func knotHash(input string) string {
	lengths := functional.Map(strings.Split(input, ""), func(s string) int {
		return int(s[0])
	})
	lengths = append(lengths, 17, 31, 73, 47, 23)

	skipSize := 0
	root := linkedlist.CircularNumbered(256)
	fixed := root
	for round := 0; round < 64; round++ {
		for _, length := range lengths {
			root.ReverseUpTo(length - 1)
			root = root.Nth(skipSize + length)
			skipSize++
		}
	}

	parts := fixed.ToSlice()
	var results []string
	for i := 0; i < 16; i++ {
		res := functional.Reduce(0, parts[16*i:16*(i+1)], func(b, p int) int {
			return b ^ p
		})
		hex := maths.ToHex(res)
		for len(hex) < 2 {
			hex = "0" + hex
		}
		results = append(results, hex)
	}

	return strings.ToLower(strings.Join(results, ""))
}

func (d *day10) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"",
			},
		},
		{
			ExpectedOutput: []string{
				"",
			},
		},
	}
}
