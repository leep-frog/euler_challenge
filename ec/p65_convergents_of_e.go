package eulerchallenge

import (
	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/fraction"
	"github.com/leep-frog/euler_challenge/maths"
)

func P65() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=65"),
		command.NewFlagNode(
			command.BoolFlag("two", 't', "find the convergence for the square root of 2"),
		),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			n := d.Int(N)

			var f *fraction.Fraction[*maths.Int]
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
					f = fraction.NewBig(maths.NewInt(int64(k)), maths.One())
				} else {
					f = f.Invert().Plus(maths.NewInt(int64(k)))
				}
			}
			o.Stdoutln(f.N.DigitSum())
		}),
	)
}
