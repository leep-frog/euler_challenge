package commandths

import (
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/command/sourcerer"
)

func CLI() sourcerer.CLI {
	return &Maths{}
}

var (
	expArg     = command.ListArg[string]("EXPRESSION", "Expression to evaluate", 1, command.UnboundedList)
	operations = regexp.MustCompile(`([\*\^\+/\(\)])`)
	minusRegex = regexp.MustCompile(`(\-)([^0-9])`)
	whitespace = regexp.MustCompile(`\s+`)
)

type Maths struct{}

func (m *Maths) Changed() bool { return false }
func (*Maths) Setup() []string { return nil }
func (m *Maths) Name() string  { return "m" }

func (m *Maths) Node() *command.Node {
	return command.AsNode(&command.BranchNode{
		Branches: map[string]*command.Node{
			"prime": nil,
		},
		Default: command.SerialNodes(
			expArg,
			&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				expressionStr := strings.Join(expArg.Get(d), " ")
				expressionStr = operations.ReplaceAllString(expressionStr, " $1 ")
				expressionStr = minusRegex.ReplaceAllString(expressionStr, " $1 $2 ")

				expression := whitespace.Split(strings.TrimSpace(expressionStr), -1)
				v, err := parse(newSequence(expression), false)
				if err != nil {
					return o.Err(err)
				}

				o.Stdoutln(v)
				return nil
			}},
		),
		DefaultCompletion: true,
	})
}
