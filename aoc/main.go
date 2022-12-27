package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/aoc/y2022"
	"github.com/leep-frog/euler_challenge/parse"
)

var (
	years = map[int]*aoc.Year{
		2022: y2022.Year(),
	}
	yearArg = command.MapArg("YEAR", "Problem year", years, false)
	dayArg  = command.Arg[int]("DAY", "Problem day", command.Between(1, 25, true))

	exampleFlag = command.BoolFlag("example", 'x', "Whether or not to run on the example input")
	suffixFlag  = command.Flag[string]("suffix", 's', "File suffix to use for problem")
)

func main() {
	command.RunNodes(node())
}

func goFile(day int) string {
	return fmt.Sprintf("d%02d.go", day)
}

func node() *command.Node {
	return command.SerialNodes(
		command.FlagNode(
			exampleFlag,
		),
		yearArg,
		dayArg,
		command.SimpleProcessor(func(i *command.Input, o command.Output, d *command.Data, ed *command.ExecuteData) error {
			day := dayArg.Get(d)
			year := yearArg.Get(d)

			dayList := year.Days

			// Create all files
			var done bool
			for d := len(dayList) + 1; d <= day; d++ {
				done = true
				generateDay(aoc.YearDir(year.Number), aoc.YearInputDir(year.Number), year.Number, d, ed)
			}
			if done {
				return nil
			}

			// TODO: Mutually exclusive flags
			suffix := suffixFlag.Get(d)
			if exampleFlag.Get(d) {
				suffix = "example"
			}

			// Otherwise, run the problem
			aoc.Run(year, day, suffix, o)
			return nil
		}, nil /* No logic for complete */),
	)
}

var (
	generated = 0
)

func generateDay(yearDir, yearInputDir string, year, day int, ed *command.ExecuteData) {
	generated++
	if generated > 25 {
		panic("Too many files generated!")
	}
	parse.Touch(filepath.Join(yearInputDir, aoc.InputFile(day, "")))
	parse.Touch(filepath.Join(yearInputDir, aoc.ExampleFile(day)))
	parse.Write(filepath.Join(yearDir, goFile(day)), fmt.Sprintf(strings.Join([]string{
		"package y%d",
		"",
		`import (`,
		`	"github.com/leep-frog/command"`,
		`	"github.com/leep-frog/euler_challenge/aoc/aoc"`,
		`)`,
		"",
		"func Day%02d() aoc.Day {",
		"\treturn &day%02d{}",
		"}",
		"",
		"type day%02d struct {}",
		"",
		"func (d *day%02d) Solve1(lines []string, o command.Output) {",
		"}",
		"",
		"func (d *day%02d) Solve2(lines []string, o command.Output) {",
		"}",
		"",
		"func (d *day%02d) Cases() []*aoc.Case {",
		"\treturn []*aoc.Case{",
		"\t\t{",
		"\t\t\tExpectedOutput: \"\",",
		"\t\t},",
		"\t\t{",
		"\t\t\tFileSuffix: \"example\",",
		"\t\t\tExpectedOutput: \"\",",
		"\t\t},",
		"\t}",
		"}",
		"",
	}, "\n"), year, day, day, day, day, day, day))

	yearFile := parse.FullPath(yearDir, "year.go")
	ed.Executable = append(ed.Executable,
		fmt.Sprintf("r \"(^.*END_OF_DAYS.*$)\" '\t\t\tDay%02d(),\n$1' %q", day, yearFile),
	)
}
