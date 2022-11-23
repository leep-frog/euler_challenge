package eulerchallenge

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/parse"
)

func FileGenerator() *command.Node {
	pn := "PROBLEM_NUMBER"
	fi := "file-input"
	fs := "FILE_SUFFIX"
	x := "example"
	ni := "no-input"
	return command.SerialNodes(
		command.FlagNode(
			command.BoolFlag(fi, 'f', "If set, new file will accept a file input; otherwise it accepts an integer, N"),
			command.BoolFlag(x, 'x', "If set, include example stuff in tests"),
			command.BoolFlag(ni, 'n', "If set, no input"),
		),
		command.Arg[int](pn, "Problem number", command.Positive[int]()),
		command.ListArg[string](fs, "suffix for file name", 1, command.UnboundedList),
		command.ExecutableNode(func(o command.Output, d *command.Data) ([]string, error) {
			includeExample := d.Bool(x)
			fileInput := d.Bool(fi)
			num := d.Int(pn)
			noInput := d.Bool(ni)

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
			} else if noInput {
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
			suffix := strings.ToLower(strings.Join(d.StringList(fs), "_"))
			if err := ioutil.WriteFile(fmt.Sprintf("p%d_%s.go", num, suffix), []byte(strings.Join(template, "\n")), 0644); err != nil {
				return nil, o.Stderrf("failed to write new file: %v", err)
			}

			// Write example files if file input
			if fileInput {
				parse.Touch(fmt.Sprintf("p%d.txt", num))
				if includeExample {
					parse.Touch(fmt.Sprintf("p%d_example.txt", num))
				}
			}

			return []string{
				// Add line to node.go
				fmt.Sprintf("r \"(^.*END_LIST.*$)\" '\t\tP%d(),\n$1' node.go", num),
			}, nil
		}),
	)
}
