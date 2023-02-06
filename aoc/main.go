package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/aoc/aoc"
	"github.com/leep-frog/euler_challenge/parse"

	// YEAR_IMPOTS
	"github.com/leep-frog/euler_challenge/aoc/y2015"
	"github.com/leep-frog/euler_challenge/aoc/y2020"
	"github.com/leep-frog/euler_challenge/aoc/y2022"
	// END_YEAR_IMPORTS
)

var (
	years = map[int]*aoc.Year{
		2022: y2022.Year(),
		2020: y2020.Year(),
		2015: y2015.Year(),
		// END_AOC_YEARS
	}
	yearArg = command.MapArg("YEAR", "Problem year", years, true)
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

func node() command.Node {
	return command.SerialNodes(
		command.FlagProcessor(
			exampleFlag,
			suffixFlag,
		),
		yearArg,
		dayArg,
		command.SimpleProcessor(func(i *command.Input, o command.Output, d *command.Data, ed *command.ExecuteData) error {
			day := dayArg.Get(d)
			year := yearArg.Get(d)

			if year == nil {
				return generateYear(aoc.YearDir(yearArg.GetKey()), aoc.YearInputDir(yearArg.GetKey()), yearArg.GetKey(), ed, o)
			}

			// TODO: Mutually exclusive flags
			suffix := suffixFlag.Get(d)
			if exampleFlag.Get(d) {
				suffix = "example"
			}

			// Otherwise, run the problem
			run(year, day, suffix, o)
			return nil
		}, nil /* No logic for complete */),
	)
}

func run(year *aoc.Year, day int, suffix string, o command.Output) {
	problem := year.Days[day-1]
	lines := parse.ReadFileLines(filepath.Join(aoc.YearInputDir(year.Number), aoc.InputFile(day, suffix)))

	// Remove newlines at end
	for len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	problem.Solve(lines, o)
}

var (
	generated = 0
)

func generateYear(yearDir, yearInputDir string, year int, ed *command.ExecuteData, o command.Output) error {
	if !parse.Exists(yearDir) {
		if err := os.Mkdir(yearDir, 0644); err != nil {
			return o.Stderrf("failed to create year dir: %v", err)
		}
	}

	if !parse.Exists(yearInputDir) {
		if err := os.Mkdir(yearInputDir, 0644); err != nil {
			return o.Stderrf("failed to create year input dir: %v", err)
		}
	}

	parse.Write(filepath.Join(yearDir, "year.go"), fmt.Sprintf(strings.Join([]string{
		"package y%d",
		"",
		`import (`,
		`	"github.com/leep-frog/euler_challenge/aoc/aoc"`,
		`)`,
		"",
		"func Year() *aoc.Year {",
		"\treturn &aoc.Year{",
		"\t\tNumber: %d,",
		"\t\tDays: []aoc.Day{",
		"\t\t\t// END_OF_DAYS",
		"\t\t},",
		"\t}",
		"}",
		"",
	}, "\n"), year, year))

	for day := 1; day <= 25; day++ {
		generateDay(yearDir, yearInputDir, year, day, ed)
	}

	mainFile := parse.FullPath(".", "main.go")
	ed.Executable = append(ed.Executable,
		// Add import
		// Plus sign to avoid replacing this line as well
		fmt.Sprintf("r \"(^.*// END_YEAR"+"_IMPORTS.*$)\" '\t\"github.com/leep-frog/euler_challenge/aoc/y%d\"\n$1' %q", year, mainFile),
		// Add map value
		fmt.Sprintf("r \"(^.*// END_AOC"+"_YEARS.*$)\" '\t\t%d: y%d.Year(),\n$1' %q", year, year, mainFile),
	)

	return nil
}

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
		"type day%02d struct{}",
		"",
		"func (d *day%02d) Solve(lines []string, o command.Output) {",
		"}",
		"",
		"func (d *day%02d) Cases() []*aoc.Case {",
		"\treturn []*aoc.Case{",
		"\t\t{",
		"\t\t\tFileSuffix: \"example\",",
		"\t\t\tExpectedOutput: []string{",
		"\t\t\t\t\"\",",
		"\t\t\t},",
		"\t\t},",
		"\t\t{",
		"\t\t\tExpectedOutput: []string{",
		"\t\t\t\t\"\",",
		"\t\t\t},",
		"\t\t},",
		"\t}",
		"}",
		"",
	}, "\n"), year, day, day, day, day, day))

	yearFile := parse.FullPath(yearDir, "year.go")
	ed.Executable = append(ed.Executable,
		fmt.Sprintf("r \"(^.*END_OF_DAYS.*$)\" '\t\t\tDay%02d(),\n$1' %q", day, yearFile),
	)
}
