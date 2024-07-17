package eulerchallenge

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/euler_challenge/parse"
)

func FileGenerator() command.Node {
	problemNumberArg := commander.Arg[int]("PROBLEM_NUMBER", "Problem number", commander.Positive[int]())
	fileSuffixArg := commander.ListArg[string]("FILE_SUFFIX", "suffix for file name", 1, command.UnboundedList)

	fileInputFlag := commander.BoolFlag("file-input", 'f', "If set, new file will accept a file input; otherwise it accepts an integer, N")
	exampleFlag := commander.BoolFlag("example", 'x', "If set, include example stuff in tests")
	noInputFlag := commander.BoolFlag("no-input", 'n', "If set, no input")

	return commander.SerialNodes(
		commander.FlagProcessor(
			fileInputFlag,
			exampleFlag,
			noInputFlag,
		),
		problemNumberArg,
		fileSuffixArg,
		commander.ExecutableProcessor(func(o command.Output, d *command.Data) ([]string, error) {
			fileInput := fileInputFlag.Get(d)
			num := problemNumberArg.Get(d)
			pNum := fmt.Sprintf("p%d", num)

			template := []string{
				fmt.Sprintf("package %s", pNum),
				"",
				"import (",
				"  \"github.com/leep-frog/command/command\"",
				"  \"github.com/leep-frog/euler_challenge/ec/ecmodels\"",
				")",
				"",
				fmt.Sprintf("func P%d() *ecmodels.Problem {", num),
			}

			if fileInput {
				template = append(template,
					fmt.Sprintf("  return ecmodels.FileInputNode(%d, func(lines []string, o command.Output) {", num),
					"    o.Stdoutln(lines)",
					"  }, []*ecmodels.Execution{",
				)

				if exampleFlag.Get(d) {
					template = append(template,
						"    {",
						`      Args: []string{"-x"},`,
						`      Want: "",`,
						"    },",
					)
				}

				template = append(template,
					"    {",
					`      Want: "",`,
					"    },",
					"  })",
					"}",
				)
			} else if noInputFlag.Get(d) {
				if exampleFlag.Get(d) {
					template = append(template,
						fmt.Sprintf("  return ecmodels.NoInputWithExampleNode(%d, func(o command.Output) {", num),
						"  }, []*ecmodels.Execution{",
						`    {`,
						`      Args: "-x",`,
						`      Want: "",`,
						`    },`,
						`    {`,
						`      Want: "",`,
						`    },`,
						"  })",
						"}",
					)
				} else {
					template = append(template,
						fmt.Sprintf("  return ecmodels.NoInputNode(%d, func(o command.Output) {", num),
						"  }, &ecmodels.Execution{",
						`    Want: "",`,
						"  })",
						"}",
					)
				}
			} else {
				template = append(template,
					fmt.Sprintf("  return ecmodels.IntInputNode(%d, func(o command.Output, n int) {", num),
					"    o.Stdoutln(n)",
					"  }, []*ecmodels.Execution{",
				)

				if exampleFlag.Get(d) {
					template = append(template,
						"    {",
						`      Args: []string{"1"},`,
						`      Want: "",`,
						"    },",
					)
				}

				template = append(template,
					"    {",
					`      Args: []string{"2"},`,
					`      Want: "",`,
					"    },",
					"  })",
					"}",
				)
			}

			// Full file paths
			_, thisFile, _, _ := runtime.Caller(0)
			ecDir := filepath.Dir(thisFile)
			nodeGo := filepath.Join(ecDir, "node.go")
			ecTestGo := filepath.Join(ecDir, "ec_test.go")

			suffix := strings.ToLower(strings.Join(fileSuffixArg.Get(d), "_"))
			newGoFile := filepath.Join(ecDir, pNum, fmt.Sprintf("p%d_%s.go", num, suffix))

			if err := os.Mkdir(filepath.Join(ecDir, pNum), 0644); err != nil {
				return nil, o.Annotate(err, "failed to create pNum directory")
			}

			// Create go file
			if err := os.WriteFile(newGoFile, []byte(strings.Join(template, "\n")), 0644); err != nil {
				return nil, o.Annotate(err, "failed to write new file")
			}

			// Write example files if file input
			if fileInput {
				touch(fmt.Sprintf("p%d.txt", num))
				if exampleFlag.Get(d) {
					touch(fmt.Sprintf("p%d_example.txt", num))
				}
			}

			return []string{
				// Add line to node.go
				fmt.Sprintf("r \"(^.*END_LIST.*$)\" '\t\t%s.P%d(),\n$1' %q", pNum, num, nodeGo),
				// Add import to node.go
				fmt.Sprintf("r \"(^.*END_IMPORT_LIST.*$)\" '\t\\\"github.com/leep-frog/euler_challenge/ec/%s\\\"\n$1' %q", pNum, nodeGo),
				// Update CURRENT_PROBLEM
				fmt.Sprintf("r '(^[^0-9]*)[0-9]+,(.*CURRENT_PROBLEM)$' '${1}%d,${2}' %q", num, ecTestGo),
			}, nil
		}),
	)
}

func touch(f string) {
	parse.Touch(filepath.Join("input", f))
}

func readFileInput(f string) string {
	return parse.ReadFileInput(filepath.Join("input", f))
}

func readFileLines(f string) []string {
	return parse.ReadFileLines(filepath.Join("input", f))
}
