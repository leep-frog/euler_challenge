package y2025

import (
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

func Day06() aoc.Day {
	return &day06{}
}

type day06 struct{}

func simplifyLine(line string) string {
	// Replace all multiple spaces with a single space
	for strings.Contains(line, "  ") {
		line = strings.ReplaceAll(line, "  ", " ")
	}

	return strings.TrimSpace(line)
}

func partOne(lines []string, o command.Output) {
	for i, line := range lines {
		lines[i] = simplifyLine(line)
	}
	lines[len(lines)-1] = strings.ReplaceAll(lines[len(lines)-1], " ", "")

	problems := maths.SimpleTranspose(parse.ToGrid(lines[:len(lines)-1], " "))
	operations := parse.MapToGrid(lines[len(lines)-1:], map[rune]bool{'*': true, '+': false})[0]

	var sum int
	for i, problem := range problems {
		operation := operations[i]
		if operation {
			cur := 1
			for _, v := range problem {
				cur *= v
			}
			sum += cur
		} else {
			cur := 0
			for _, v := range problem {
				cur += v
			}
			sum += cur

		}
	}

	o.Stdoutln(sum)
}

func partTwo(lines []string, o command.Output) {

	operationLine := simplifyLine(lines[len(lines)-1])
	operationLine = strings.ReplaceAll(operationLine, " ", "")
	lines = lines[:len(lines)-1]

	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	var charGrid [][]string
	for _, line := range lines {
		charGrid = append(charGrid, strings.Split(line, ""))
		for len(charGrid[len(charGrid)-1]) < maxLen {
			charGrid[len(charGrid)-1] = append(charGrid[len(charGrid)-1], " ")
		}
	}

	charGrid = maths.SimpleTranspose(charGrid)

	var rejoinedLines []string
	for _, charLine := range charGrid {
		rejoinedLines = append(rejoinedLines, strings.Join(charLine, ""))
	}

	var sum int
	idx := 0
	for _, operation := range operationLine {
		addition := operation == '+'

		var cur int
		if !addition {
			cur = 1
		}
		for ; idx < len(rejoinedLines) && strings.TrimSpace(rejoinedLines[idx]) != ""; idx++ {
			num := parse.Atoi(strings.TrimSpace(rejoinedLines[idx]))
			if addition {
				cur += num
			} else {
				cur *= num
			}
		}
		idx++
		sum += cur
	}
	o.Stdoutln(sum)
}

func (d *day06) Solve(lines []string, o command.Output) {
	// partOne(lines, o)
	partTwo(lines, o)
}

func (d *day06) Cases() []*aoc.Case {
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
