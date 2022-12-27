package eulerchallenge

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func FileGenerator() *command.Node {
	problemNumberArg := command.Arg[int]("PROBLEM_NUMBER", "Problem number", command.Positive[int]())
	fileSuffixArg := command.ListArg[string]("FILE_SUFFIX", "suffix for file name", 1, command.UnboundedList)

	fileInputFlag := command.BoolFlag("file-input", 'f', "If set, new file will accept a file input; otherwise it accepts an integer, N")
	exampleFlag := command.BoolFlag("example", 'x', "If set, include example stuff in tests")
	noInputFlag := command.BoolFlag("no-input", 'n', "If set, no input")

	return command.SerialNodes(
		command.FlagNode(
			fileInputFlag,
			exampleFlag,
			noInputFlag,
		),
		problemNumberArg,
		fileSuffixArg,
		command.ExecutableNode(func(o command.Output, d *command.Data) ([]string, error) {
			fileInput := fileInputFlag.Get(d)
			num := problemNumberArg.Get(d)

			template := []string{
				"package eulerchallenge",
				"",
				"import (",
				"  \"github.com/leep-frog/command\"",
				")",
				"",
				fmt.Sprintf("func P%d() *problem {", num),
			}

			if fileInput {
				template = append(template,
					fmt.Sprintf("  return fileInputNode(%d, func(lines []string, o command.Output) {", num),
					"    o.Stdoutln(lines)",
					"  }, []*execution{",
					"    {",
					fmt.Sprintf(`      args: []string{"%d"},`, num),
					`      want: "",`,
					"    },",
					"  })",
					"}",
				)
			} else if noInputFlag.Get(d) {
				template = append(template,
					fmt.Sprintf("  return noInputNode(%d, func(o command.Output) {", num),
					"  })",
					"}",
				)
			} else {
				template = append(template,
					fmt.Sprintf("  return intInputNode(%d, func(o command.Output, n int) {", num),
					"    o.Stdoutln(n)",
					"  }, []*execution{",
					"    {",
					`      args: []string{"1"},`,
					`      want: "",`,
					"    },",
					"  })",
					"}",
				)
			}

			template = append(template)

			// Create go file
			suffix := strings.ToLower(strings.Join(fileSuffixArg.Get(d), "_"))
			if err := ioutil.WriteFile(fmt.Sprintf("p%d_%s.go", num, suffix), []byte(strings.Join(template, "\n")), 0644); err != nil {
				return nil, o.Stderrf("failed to write new file: %v", err)
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
				fmt.Sprintf("r \"(^.*END_LIST.*$)\" '\t\tP%d(),\n$1' node.go", num),
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
