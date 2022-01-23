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
		command.NewFlagNode(
			command.BoolFlag(fi, 'i', "If set, new file will accept a file input; otherwise it accepts an integer, N"),
			command.BoolFlag(x, 'x', "If set, include example stuff in tests"),
			command.BoolFlag(ni, 'n', "If set, no input"),
		),
		command.Arg[int](pn, "Problem number", command.Positive[int]()),
		command.Arg[string](fs, "suffix for file name"),
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
					fmt.Sprintf("return fileInputNode(%d, func(lines []string, o command.Output) {", num),
					"o.Stdoutln(lines)",
				)
			} else if noInput {
				template = append(template,
					fmt.Sprintf("return fileInputNode(%d, func(o command.Output) {", num),
				)
			} else {
				template = append(template,
					fmt.Sprintf("  return intInputNode(%d, func(o command.Output, n int) {", num),
					"    o.Stdoutln(n)",
				)
			}

			template = append(template,
				"  })",
				"}",
			)

			// Create go file
			if err := ioutil.WriteFile(fmt.Sprintf("p%d_%s.go", num, d.String(fs)), []byte(strings.Join(template, "\n")), 0644); err != nil {
				return nil, o.Stderrf("failed to write new file: %v", err)
			}

			// Write example files if file input
			if fileInput {
				parse.Touch(fmt.Sprintf("p%d.txt", num))
				if includeExample {
					parse.Touch(fmt.Sprintf("p%d_example.txt", num))
				}
			}

			testFmt := "\t\t{\n\t\t\tname: \"p%d%s\",\n\t\t\targs: []string{\"%d\"%s},\n\t\t\twant: []string{\"0\"},\n\t\t},"
			testArg := ", \"1\""
			exTestArg := ", \"1\""
			if fileInput {
				testArg = fmt.Sprintf(", \"p%d.txt\"", num)
				exTestArg = fmt.Sprintf(", \"p%d_example.txt\"", num)
			} else if noInput {
				testArg = ""
				exTestArg = ""
			}

			exTest := fmt.Sprintf(testFmt, num, " example", num, exTestArg)
			test := fmt.Sprintf(testFmt, num, "", num, testArg)
			testStr := fmt.Sprintf("r \"(^.*TEST_START.*)$\" '$1\n%s' ec_test.go", test)
			if includeExample {
				testStr = fmt.Sprintf("r \"(^.*TEST_START.*)$\" '$1\n%s\n%s' ec_test.go", test, exTest)
			}
			return []string{
				`r "\/\*\{" "{" ec_test.go`,
				// Add line to node.go
				fmt.Sprintf("r \"(^.*END_LIST.*$)\" '\t\tP%d(),\n$1' node.go", num),
				// Add tests to ec_test.go
				testStr,
				// TODO: remove comment start"r "/*"
			}, nil
		}),
	)
}
