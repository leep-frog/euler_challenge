package commandths

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/leep-frog/command"
	"github.com/leep-frog/command/sourcerer"
	"github.com/leep-frog/euler_challenge/generator"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func CLI() sourcerer.CLI {
	return &Maths{}
}

func Aliasers() sourcerer.Option {
	return sourcerer.Aliasers(map[string][]string{
		"pf": {"prime", "factor"},
	})
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

func (m *Maths) primeFactor() *command.Node {
	arg := command.ListArg[int]("N", "The numbers to prime factor", 1, command.UnboundedList)
	return command.SerialNodes(
		command.Description("Prints out the prime factors of the provided numbers"),
		arg,
		&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
			p := generator.Primes()
			for _, a := range arg.Get(d) {
				fs := p.PrimeFactors(a)
				ordered := maps.Keys(fs)
				slices.Sort(ordered)

				var r []string
				for _, f := range ordered {
					r = append(r, fmt.Sprintf("%d^%d", f, fs[f]))
				}
				o.Stdoutf("%d: %s\n", a, strings.Join(r, " * "))
			}
			return nil
		}},
	)
}

func (m *Maths) nthPrime() *command.Node {
	arg := command.Arg[int]("N", "The prime index to get", command.NonNegative[int]())
	return command.SerialNodes(
		command.Description("Prints out the Nth prime number (1-indexed)"),
		arg,
		&command.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
			o.Stdoutln(generator.Primes().Nth(arg.Get(d) - 1))
			return nil
		}},
	)
}

func (m *Maths) Node() *command.Node {
	return command.AsNode(&command.BranchNode{
		Branches: map[string]*command.Node{
			"prime": command.AsNode(&command.BranchNode{
				Branches: map[string]*command.Node{
					"factor": m.primeFactor(),
					"nth":    m.nthPrime(),
				},
			}),
		},
		Default: command.SerialNodes(
			// TODO: Flag(s) to change mode (int, float, fraction)
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
