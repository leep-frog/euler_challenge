package eulerchallenge

import (
	"fmt"
	"math"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/maths"
)

func P9() *command.Node {
	return command.SerialNodes(
		command.Description("Find the Pythagorean triplet that equals N"),
		command.IntNode(N, "", command.IntPositive()),
		command.ExecuteErrNode(func(o command.Output, d *command.Data) error {
			n := d.Int(N)

			for a := 1; a < n; a++ {
				for b := a + 1; b+a < n; b++ {
					c2 := (a*a + b*b)
					c := int(math.Sqrt(float64(c2)))
					if a+b+c == 1000 && maths.IsSquare(c2) {
						o.Stdoutf("%d", a*b*c)
						return nil
					}
				}
			}
			return fmt.Errorf("failed to find triplet")
		}),
	)
}
