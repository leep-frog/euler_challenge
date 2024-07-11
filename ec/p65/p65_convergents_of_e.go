package p65

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/command/commander"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/fraction"
)

func P65() *ecmodels.Problem {
	return &ecmodels.Problem{
		Num: 65,
		N: commander.SerialNodes(
			commander.Description("https://projecteuler.net/problem=65"),
			commander.FlagProcessor(
				commander.BoolFlag("two", 't', "find the convergence for the square root of 2"),
			),
			commander.Arg[int](ecmodels.N, "", commander.Positive[int]()),
			&commander.ExecutorProcessor{F: func(o command.Output, d *command.Data) error {
				n := d.Int(ecmodels.N)

				var f *fraction.Rational
				for idx := n - 1; idx >= 0; idx-- {
					k := 1
					if idx%3 == 2 {
						k = ((idx / 3) + 1) * 2
					}
					if idx == 0 {
						k = 2
					}
					if d.Bool("two") {
						k = 2
						if idx == 0 {
							k = 1
						}
					}

					if f == nil {
						f = fraction.NewRational(k, 1)
					} else {
						f = f.Reciprocal().Plus(fraction.NewRational(k, 1))
					}
				}
				o.Stdoutln(f.Numer().DigitSum())
				return nil
			}},
		),
		Executions: []*ecmodels.Execution{
			{
				Args: []string{"100"},
				Want: "272",
			},
			{
				Args: []string{"10"},
				Want: "17",
			},
		},
	}
}
