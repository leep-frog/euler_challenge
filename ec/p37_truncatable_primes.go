package eulerchallenge

import (
	"strconv"

	"github.com/leep-frog/command"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/parse"
)

func P37() *command.Node {
	return command.SerialNodes(
		command.Description("https://projecteuler.net/problem=37"),
		command.ExecutorNode(func(o command.Output, d *command.Data) {
			p := generator.Primes()

			var count, sum int
			for i := 0; count < 11; i++ {
				pn := p.Nth(i)
				if pn < 10 {
					continue
				}
				pnStr := strconv.Itoa(pn)
				valid := true
				for j := 1; j < len(pnStr); j++ {
					leftTrunc := pnStr[j:]
					rightTrunc := pnStr[:j]
					if !p.Contains(parse.Atoi(leftTrunc)) || !p.Contains(parse.Atoi(rightTrunc)) {
						valid = false
						break
					}
				}
				if valid {
					count++
					sum += pn
				}
			}
			o.Stdoutln(sum)
		}),
	)
}
