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

			/*resp, err := http.Get("https://projecteuler.net/minimal=38")
			if err != nil {
				return nil, o.Stderrf("failed to get problem web page: %v", err)
			}

			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, o.Stderrf("failed to read body: %v", err)
			}
			body := string(bodyBytes)

			if resp.StatusCode != 200 {
				return nil, o.Stderrf("response has status code (%d):\n%v", resp.StatusCode, body)
			}

			o.Stdout(body)*/

			arg := "    command.IntNode(N, \"\", command.IntPositive()),"
			loader := "      n := n"
			printer := "      o.Stdoutln(n)"
			if fileInput {
				arg = "    command.StringNode(\"FILE\", \"\"),"
				loader = "      lines := parse.ReadFileLines(d.String(\"FILE\"))"
				printer = "      o.Stdoutln(lines)"
			} else if noInput {
				printer = "      o.Stdoutln(0)"
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
				fmt.Sprintf("func P%d() *problem {", num),
				"  return command.SerialNodes(",
				fmt.Sprintf("    command.Description(\"https://projecteuler.net/problem=%d\"),", num),
			)

			if !noInput {
				template = append(template,
					arg,
					"    command.ExecutorNode(func(o command.Output, d *command.Data) {",
					loader,
					printer,
				)
			} else {
				template = append(template, "    command.ExecutorNode(func(o command.Output, d *command.Data) {")
			}

			template = append(template,
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
				// Add line to node.go
				fmt.Sprintf("r \"(^.*END_LIST.*$)\" '\t\t\"%d\": P%d(),\n$1' node.go", num, num),
				// Add tests to ec_test.go
				testStr,
			}, nil
		}),
	)
}
