package aoc

import (
	"fmt"
	"path/filepath"

	"github.com/leep-frog/command/command"
)

type Day interface {
	Solve([]string, command.Output)
	Cases() []*Case
}

type Case struct {
	FileSuffix     string
	ExpectedOutput []string
}

type Year struct {
	Number int
	Days   []Day
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
