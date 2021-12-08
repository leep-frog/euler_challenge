package main

import (
	"fmt"
	"math"

	"github.com/leep-frog/command"
)

func p9() *command.Node {
	return command.SerialNodes(
		command.Description("Find the Pythagorean triplet that equals N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecutorNode(func(o command.Output, d *command.Data) error {
			n := d.Int(N)

			for a := 1; a < n; a++ {
				for b := a + 1; b+a < n; b++ {
					c2 := (a*a + b*b)
					c := int(math.Sqrt(float64(c2)))
					if a+b+c == 1000 && IsSquare(c2) {
						o.Stdoutf("%d", a*b*c)
						return nil
					}
				}
			}
			return fmt.Errorf("failed to find triplet")
		}),
	)
}
