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
	desc := "DESCRIPTION"
	return command.SerialNodes(
		command.NewFlagNode(
			command.BoolFlag(fi, 'i', "If set, new file will accept a file input; otherwise it accepts an integer, N"),
		),
		command.IntNode(pn, "Problem number", command.IntPositive()),
		command.StringNode(fs, "suffix for file name"),
		command.StringListNode(desc, "Description of the problem", 1, command.UnboundedList),
		command.ExecutableNode(func(o command.Output, d *command.Data) ([]string, error) {
			fileInput := d.Bool(fi)
			num := d.Int(pn)

			arg := "    command.IntNode(N, \"\", command.IntPositive()),"
			loader := "      n := d.Int(N)"
			printer := "      o.Stdoutln(n)"
			if fileInput {
				arg = "    command.StringNode(\"FILE\", \"\"),"
				loader = "      lines := parse.ReadFileLines(d.String(\"FILE\"))"
				printer = "      o.Stdoutln(lines)"
			}

			template := []string{
				"package eulerchallenge",
				"",
				"import (",
				"  \"github.com/leep-frog/command\"",
			}

			if fileInput {
				template = append(template, "  \"github.com/leep-frog/euler_challenge/parse\"")
			}

			template = append(template,
				")",
				"",
				fmt.Sprintf("func P%d() *command.Node {", num),
				"  return command.SerialNodes(",
				fmt.Sprintf("    command.Description(\"%s\"),", strings.Join(d.StringList(desc), " ")),
				arg,
				"    command.ExecutorNode(func(o command.Output, d *command.Data) {",
				loader,
				printer,
				"    }),",
				"  )",
				"}",
			)

			// Create go file
			if err := ioutil.WriteFile(fmt.Sprintf("p%d_%s.go", num, d.String(fs)), []byte(strings.Join(template, "\n")), 0644); err != nil {
				return nil, o.Stderrf("failed to write new file: %v", err)
			}

			// Write example files if file input
			if fileInput {
				parse.Touch(fmt.Sprintf("p%d.txt", num))
				parse.Touch(fmt.Sprintf("p%d_example.txt", num))
			}

			testFmt := "\t\t{\n\t\t\tname: \"p%d%s\",\n\t\t\targs: []string{\"%d\", \"%s\"},\n\t\t\twant: []string{\"0\"},\n\t\t},"
			testArg := "1"
			exTestArg := "1"
			if fileInput {
				testArg = fmt.Sprintf("p%d.txt", num)
				exTestArg = fmt.Sprintf("p%d_example.txt", num)
			}

			exTest := fmt.Sprintf(testFmt, num, " example", num, exTestArg)
			test := fmt.Sprintf(testFmt, num, "", num, testArg)
			return []string{
				// Add line to node.go
				fmt.Sprintf("r \"(^.*END_LIST.*$)\" '\t\t\"%d\": P%d(),\n$1' node.go", num, num),
				// Add tests to ec_test.go
				fmt.Sprintf("r \"(^.*TEST_START.*)$\" '$1\n%s\n%s' ec_test.go", test, exTest),
			}, nil
		}),
	)
}
