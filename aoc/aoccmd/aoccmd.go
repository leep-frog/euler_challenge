package aoccmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/command/sourcerer"
	"github.com/leep-frog/euler_challenge/aoc"
	"github.com/leep-frog/euler_challenge/parse"

	// YEAR_IMPOTS
	"github.com/leep-frog/euler_challenge/aoc/y2015"
	"github.com/leep-frog/euler_challenge/aoc/y2016"
	"github.com/leep-frog/euler_challenge/aoc/y2020"
	"github.com/leep-frog/euler_challenge/aoc/y2022"
	"github.com/leep-frog/euler_challenge/aoc/y2017"
	// END_YEAR_IMPORTS
)

var (
	years = map[int]*aoc.Year{
		2022: y2022.Year(),
		22:   y2022.Year(),
		2020: y2020.Year(),
		20:   y2020.Year(),
		2015: y2015.Year(),
		15:   y2015.Year(),
		2016: y2016.Year(),
		16:   y2016.Year(),
		2017: y2017.Year(),
		17: y2017.Year(),
		// END_AOC_YEARS
	}
)

func CLI() sourcerer.CLI {
	return &AdventOfCode{years, 0, false}
}

func Aliasers() sourcerer.Option {
	return sourcerer.Aliasers(map[string][]string{
		"ac": nil,
	})
}

type AdventOfCode struct {
	years       map[int]*aoc.Year
	DefaultYear int
	changed     bool
}

func (*AdventOfCode) Name() string    { return "aoc" }
func (*AdventOfCode) Setup() []string { return nil }

func (a *AdventOfCode) Changed() bool { return a.changed }

func goFile(day int) string {
	return fmt.Sprintf("d%02d.go", day)
}

func (a *AdventOfCode) Node() command.Node {
	yearArg := command.MapArg("YEAR", "Problem year", a.years, true)
	dayArg := command.Arg[int]("DAY", "Problem day", command.Between(1, 25, true))

	exampleFlag := command.BoolFlag("example", 'x', "Whether or not to run on the example input")
	suffixFlag := command.Flag[string]("suffix", 's', "File suffix to use for problem")

	var usedDefault bool

	mainNode := command.SerialNodes(
		command.FlagProcessor(
			exampleFlag,
			suffixFlag,
		),
		// Adds default year if first argument is "d"
		command.SuperSimpleProcessor(func(i *command.Input, d *command.Data) error {
			if s, ok := i.Peek(); !ok || s != "d" {
				return nil
			}
			i.Pop(d)
			i.PushFront(fmt.Sprintf("%d", a.DefaultYear))
			usedDefault = true
			return nil
		}),
		yearArg,
		dayArg,
		command.SimpleProcessor(func(i *command.Input, o command.Output, d *command.Data, ed *command.ExecuteData) error {
			day := dayArg.Get(d)
			year := yearArg.Get(d)

			if year == nil {
				if usedDefault {
					return o.Stderrf("Default year does not exist!")
				}
				yearNumber := yearArg.GetKey()
				if yearNumber < 2000 {
					yearNumber += 2000
				}
				return generateYear(aoc.YearDir(yearNumber), aoc.YearInputDir(yearNumber), yearNumber, ed, o)
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

	return &command.BranchNode{
		Branches: map[string]command.Node{
			"setDefault": command.SerialNodes(
				command.Arg[int]("DEFAULT_YEAR", "Default year to use"),
				&command.ExecutorProcessor{func(o command.Output, d *command.Data) error {
					a.DefaultYear = d.Int("DEFAULT_YEAR")
					a.changed = true
					return nil
				}},
			),
		},
		Default: mainNode,
	}
}

func run(year *aoc.Year, day int, suffix string, o command.Output) {
	problem := year.Days[day-1]
	lines := parse.ReadFileLines(filepath.Join("..", aoc.YearInputDir(year.Number), aoc.InputFile(day, suffix)))

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
	yearDir = filepath.Join("..", yearDir)
	yearInputDir = filepath.Join("..", yearInputDir)
	parse.Mkdir(yearDir)
	parse.Mkdir(yearInputDir)

	dayFmt := "\t\t\tDay%02d(),"
	parse.Write(filepath.Join(yearDir, "year.go"), fmt.Sprintf(strings.Join([]string{
		"package y%d",
		"",
		`import (`,
		`	"github.com/leep-frog/euler_challenge/aoc"`,
		`)`,
		"",
		"func Year() *aoc.Year {",
		"\treturn &aoc.Year{",
		"\t\tNumber: %d,",
		"\t\tDays: []aoc.Day{",
		fmt.Sprintf(dayFmt, 1),
		fmt.Sprintf(dayFmt, 2),
		fmt.Sprintf(dayFmt, 3),
		fmt.Sprintf(dayFmt, 4),
		fmt.Sprintf(dayFmt, 5),
		fmt.Sprintf(dayFmt, 6),
		fmt.Sprintf(dayFmt, 7),
		fmt.Sprintf(dayFmt, 8),
		fmt.Sprintf(dayFmt, 9),
		fmt.Sprintf(dayFmt, 10),
		fmt.Sprintf(dayFmt, 11),
		fmt.Sprintf(dayFmt, 12),
		fmt.Sprintf(dayFmt, 13),
		fmt.Sprintf(dayFmt, 14),
		fmt.Sprintf(dayFmt, 15),
		fmt.Sprintf(dayFmt, 16),
		fmt.Sprintf(dayFmt, 17),
		fmt.Sprintf(dayFmt, 18),
		fmt.Sprintf(dayFmt, 19),
		fmt.Sprintf(dayFmt, 20),
		fmt.Sprintf(dayFmt, 21),
		fmt.Sprintf(dayFmt, 22),
		fmt.Sprintf(dayFmt, 23),
		fmt.Sprintf(dayFmt, 24),
		fmt.Sprintf(dayFmt, 25),
		"\t\t},",
		"\t}",
		"}",
		"",
	}, "\n"), year, year))

	for day := 1; day <= 25; day++ {
		generateDay(yearDir, yearInputDir, year, day, ed)
	}

	mainFile := parse.FullPath(".", "aoccmd.go")
	ed.Executable = append(ed.Executable,
		// Add import
		// Plus sign to avoid replacing this line as well
		fmt.Sprintf("r \"(^.*// END_YEAR"+"_IMPORTS.*$)\" '\t\\\"github.com/leep-frog/euler_challenge/aoc/y%d\\\"\n$1' %q", year, mainFile),
		// Add map value
		fmt.Sprintf("r \"(^.*// END_AOC"+"_YEARS.*$)\" '\t\t%d: y%d.Year(),\n\t\t%d: y%d.Year(),\n$1' %q", year, year, year%100, year, mainFile),
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
		`	"github.com/leep-frog/euler_challenge/aoc"`,
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
}
