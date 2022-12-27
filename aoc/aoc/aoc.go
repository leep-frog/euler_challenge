package aoc

import (
	"fmt"
	"path/filepath"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
	"golang.org/x/exp/slices"
)

type Year struct {
	Number int
	Days   []Day
}

type Case struct {
	FileSuffix     string
	ExpectedOutput string
}

type Day interface {
	Solve1([]string, command.Output)
	Solve2([]string, command.Output)
	Cases() []*Case
}

const (
	InputDir = "inputs"
)

func InputFile(day int, suffix string) string {
	if suffix == "" {
		return fmt.Sprintf("d%02d.txt", day)
	}
	return fmt.Sprintf("d%02d_%s.txt", day, suffix)
}

func ExampleFile(day int) string {
	return InputFile(day, "example")
}

func YearDir(year int) string {
	return fmt.Sprintf("y%d", year)
}

func YearInputDir(year int) string {
	return filepath.Join(YearDir(year), InputDir)
}

func Run(year *Year, day int, suffix string, o command.Output) {
	lines := parse.ReadFileLines(filepath.Join(YearInputDir(year.Number), InputFile(day, suffix)))

	problem := year.Days[day-1]
	problem.Solve1(slices.Clone(lines), o)
	problem.Solve2(lines, o)
}
