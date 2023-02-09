package y2016

import (
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day09() aoc.Day {
	return &day09{}
}

type day09 struct{}

func (d *day09) Solve(lines []string, o command.Output) {
	o.Stdoutln(d.decompress(strings.Split(lines[0], ""), false), d.decompress(strings.Split(lines[0], ""), true))
}

func (d *day09) decompress(input []string, part2 bool) int {
	var inParen bool
	var parenContents []string
	var length int
	for len(input) > 0 {
		// Pop the character
		c := input[0]
		input = input[1:]

		if inParen {
			if c == ")" {
				parts := strings.Split(strings.Join(parenContents, ""), "x")
				numChars, numTimes := parse.Atoi(parts[0]), parse.Atoi(parts[1])

				inParen = false
				if part2 {
					length += d.decompress(input[:numChars], true) * numTimes
				} else {
					length += numChars * numTimes
				}
				input = input[numChars:]
			} else {
				parenContents = append(parenContents, c)
			}
			continue
		}

		// Not in paren
		if c == "(" {
			inParen = true
			parenContents = nil
		} else {
			length++
		}
	}
	return length
}

func (d *day09) Cases() []*aoc.Case {
	return []*aoc.Case{
		{
			FileSuffix: "example",
			ExpectedOutput: []string{
				"18 20",
			},
		},
		{
			ExpectedOutput: []string{
				"110346 10774309173",
			},
		},
	}
}
